package api

import (
	"context"
	"errors"
	"io"
	"math/rand"
	"sync"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/price"
	"github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/message/npool/sphinxproxy"
	scconst "github.com/NpoolPlatform/sphinx-coininfo/pkg/message/const"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/crud"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent"
	constant "github.com/NpoolPlatform/sphinx-proxy/pkg/message/const"
	"github.com/filecoin-project/specs-actors/actors/builtin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/*
!important
now record sign and plugin conn in service memory, cause not start multi pod,
next we can record the connect in db or other service(eg: redis), then
we can start multi pod
*/

type lmPluginType map[sphinxplugin.CoinType][]*mPlugin

var (
	ErrNoSignServiceFound   = errors.New("no sign service conn found")
	ErrNoPluginServiceFound = errors.New("no plugin service conn found")
)

var (
	// rand stream client
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

	slk            sync.RWMutex
	plk            sync.RWMutex
	lmSign         = make([]*mSign, 0)
	lmPlugin       = make(lmPluginType)
	channelBufSize = 100
)

func getProxySign() (*mSign, error) {
	slk.RLock()
	defer slk.RUnlock()
	logger.Sugar().Infof("get proxy sign length: %v", len(lmSign))
	if len(lmSign) == 0 {
		return nil, ErrNoSignServiceFound
	}
	return lmSign[rnd.Intn(len(lmSign))], nil
}

func getProxyPlugin(coinType sphinxplugin.CoinType) (*mPlugin, error) {
	plk.RLock()
	defer plk.RUnlock()
	logger.Sugar().Infof("get coin %v proxy plugin length: %v", coinType, len(lmPlugin[coinType]))
	if len(lmPlugin[coinType]) == 0 {
		return nil, ErrNoPluginServiceFound
	}
	return lmPlugin[coinType][rnd.Intn(len(lmPlugin[coinType]))], nil
}

// nolint
func Transaction(exitChan chan struct{}) {
	for {
		select {
		case <-exitChan:
			plk.Lock()
			logger.Sugar().Info("release plugin conn resource")
			for _, plugin := range lmPlugin {
				for _, pc := range plugin {
					close(pc.exitChan)
				}
			}
			plk.Unlock()
			slk.Lock()
			logger.Sugar().Info("release sign conn resource")
			for _, sign := range lmSign {
				close(sign.exitChan)
			}
			slk.Unlock()
			return
		case <-time.NewTicker(constant.TaskDuration).C:
			func() {
				ctx, cancel := context.WithTimeout(context.Background(), constant.TaskTimeout)
				defer cancel()

				var (
					err   error
					trans []*ent.Transaction
				)

				// priority deal sign
				// TODO if one wallet error will block all
				// here should wait all done
				trans, err = crud.GetTransactions(ctx)
				if err != nil {
					logger.Sugar().Errorf("call GetTransactions get transaction error: %v", err)
					return
				}
				for _, tran := range trans {
					switch tran.State {
					case uint8(sphinxproxy.TransactionState_TransactionStateWait):
						// from wallet get nonce/utxo
						coinType := sphinxplugin.CoinType(tran.CoinType)
						pluginProxy, err := getProxyPlugin(coinType)
						if err != nil {
							logger.Sugar().Errorf("proxy->plugin invalid coin %v connection", coinType)
							continue
						}
						ppRequest := &sphinxproxy.ProxyPluginRequest{
							CoinType:        coinType,
							TransactionType: sphinxproxy.TransactionType_PreSign,
							TransactionID:   tran.TransactionID,
							Address:         tran.From,
						}

						switch coinType {
						case
							sphinxplugin.CoinType_CoinTypeusdttrc20,
							sphinxplugin.CoinType_CoinTypetusdttrc20:
							ppRequest.Message = &sphinxplugin.UnsignedMessage{}
							ppRequest.Message.From = tran.From
							ppRequest.Message.To = tran.To
							ppRequest.Message.Value = price.DBPriceToVisualPrice(tran.Amount)
						}
						pluginProxy.preSign <- ppRequest
					case uint8(sphinxproxy.TransactionState_TransactionStateSign):
						// sign -> broadcast
						signProxy, err := getProxySign()
						if err != nil {
							logger.Sugar().Errorf("proxy->sign invalid coin %v connection", tran.CoinType)
							continue
						}

						coinType := sphinxplugin.CoinType(tran.CoinType)

						gasLimit := int64(0)
						nonce := uint64(0)
						recentBHash := string("")
						txData := []byte{}

						switch coinType {
						case
							sphinxplugin.CoinType_CoinTypefilecoin,
							sphinxplugin.CoinType_CoinTypetfilecoin:
							gasLimit = 200000000
							nonce = tran.Nonce
						case
							sphinxplugin.CoinType_CoinTypeethereum,
							sphinxplugin.CoinType_CoinTypeusdterc20,
							sphinxplugin.CoinType_CoinTypetethereum,
							sphinxplugin.CoinType_CoinTypetusdterc20,
							sphinxplugin.CoinType_CoinTypebsc,
							sphinxplugin.CoinType_CoinTypetbsc,
							sphinxplugin.CoinType_CoinTypebusdbep20,
							sphinxplugin.CoinType_CoinTypetbusdbep20:
							gasLimit = tran.Pre.GasLimit
							nonce = tran.Pre.Nonce
						case
							sphinxplugin.CoinType_CoinTypesolana,
							sphinxplugin.CoinType_CoinTypetsolana:
							recentBHash = tran.RecentBhash
						case
							sphinxplugin.CoinType_CoinTypeusdttrc20,
							sphinxplugin.CoinType_CoinTypetusdttrc20:
							txData = tran.TxData
						}

						signProxy.sign <- &sphinxproxy.ProxySignRequest{
							TransactionType: sphinxproxy.TransactionType_Signature,
							CoinType:        coinType,
							TransactionID:   tran.TransactionID,
							Message: &sphinxplugin.UnsignedMessage{
								To:    tran.To,
								From:  tran.From,
								Value: price.DBPriceToVisualPrice(tran.Amount),
								// TODO from chain get
								GasLimit:   gasLimit,
								GasFeeCap:  10000000,
								GasPremium: 1000000,
								Method:     uint64(builtin.MethodSend),
								// fil
								Nonce: nonce,
								// TODO optimize btc
								Unspent: tran.Utxo,
								// eth/erc20/bsc/bep20
								GasPrice:   tran.Pre.GasPrice,
								ChainID:    tran.Pre.ChainID,
								ContractID: tran.Pre.ContractID,
								// sol
								RecentBhash: recentBHash,
								// trc20
								TxData: txData,
							},
						}
					case uint8(sphinxproxy.TransactionState_TransactionStateSync):
						coinType := sphinxplugin.CoinType(tran.CoinType)
						pluginProxy, err := getProxyPlugin(coinType)
						if err != nil {
							logger.Sugar().Errorf("proxy->plugin invalid coin %v connection", coinType)
							continue
						}
						pluginProxy.syncMsg <- &sphinxproxy.ProxyPluginRequest{
							CoinType:        coinType,
							TransactionType: sphinxproxy.TransactionType_SyncMsgState,
							TransactionID:   tran.TransactionID,
							CID:             tran.Cid,
							Message: &sphinxplugin.UnsignedMessage{
								From: tran.From,
								To:   tran.To,
							},
						}
					}
				}
			}()
		}
	}
}

func registerCoin(coinInfo *coininfo.CreateCoinInfoRequest) error {
	ackConn, err := grpc2.GetGRPCConn(scconst.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return err
	}
	defer ackConn.Close()

	client := coininfo.NewSphinxCoinInfoClient(ackConn)

	ctx, cancel := context.WithTimeout(context.Background(), constant.GrpcTimeout)
	defer cancel()

	// define in plugin
	_, err = client.CreateCoinInfo(ctx, coinInfo)
	return err
}

func checkCode(err error) bool {
	if err == io.EOF ||
		status.Code(err) == codes.Unavailable ||
		status.Code(err) == codes.Canceled {
		return true
	}
	return false
}
