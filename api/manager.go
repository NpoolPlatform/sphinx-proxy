package api

import (
	"context"
	"errors"
	"fmt"
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
	"github.com/NpoolPlatform/sphinx-proxy/pkg/unit"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/utils"
	"github.com/filecoin-project/go-state-types/exitcode"
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
	logger.Sugar().Infof("get proxy plugin length: %v", len(lmPlugin[coinType]))
	if len(lmPlugin[coinType]) == 0 {
		return nil, ErrNoPluginServiceFound
	}
	return lmPlugin[coinType][rnd.Intn(len(lmPlugin[coinType]))], nil
}

type mSign struct {
	signServer   sphinxproxy.SphinxProxy_ProxySignServer
	closeFlag    bool
	closeChannel chan struct{}
	walletNew    chan *sphinxproxy.ProxySignRequest
	sign         chan *sphinxproxy.ProxySignRequest
}

func newSignStream(stream sphinxproxy.SphinxProxy_ProxySignServer) {
	lc := &mSign{
		signServer:   stream,
		closeFlag:    false,
		closeChannel: make(chan struct{}),
		walletNew:    make(chan *sphinxproxy.ProxySignRequest, channelBufSize),
		sign:         make(chan *sphinxproxy.ProxySignRequest, channelBufSize),
	}
	slk.Lock()
	lmSign = append(lmSign, lc)
	slk.Unlock()

	wg := &sync.WaitGroup{}
	wg.Add(3)
	go lc.signStreamSend(wg)
	go lc.signStreamRecv(wg)
	go lc.close(wg)
	wg.Wait()
	logger.Sugar().Info("some sign client down, close it")
}

func (s *mSign) signStreamSend(wg *sync.WaitGroup) {
	defer wg.Done()
send:
	for !s.closeFlag {
		select {
		case info := <-s.walletNew:
			logger.Sugar().Infof("proxy->sign TransactionID: %v start", info.GetTransactionID())
			if err := s.signServer.Send(&sphinxproxy.ProxySignRequest{
				TransactionType: info.GetTransactionType(),
				CoinType:        info.GetCoinType(),
				TransactionID:   info.GetTransactionID(),
			}); err != nil {
				logger.Sugar().Errorf(
					"proxy->sign TransactionID: %v TransactionType %v, CoinType: %v error: %v",
					info.GetTransactionType(),
					info.GetCoinType(),
					info.GetTransactionID(),
					err,
				)
				if ch, ok := walletDoneChannel.Load(info.GetTransactionID()); ok {
					ch.(chan walletDoneInfo) <- walletDoneInfo{
						success: false,
						message: fmt.Sprintf("proxy->sign send create wallet error: %v", err),
					}
				}
				if checkCode(err) {
					s.closeChannel <- struct{}{}
					break send
				}
				continue
			}
			logger.Sugar().Infof("proxy->sign TransactionID: %v ok", info.GetTransactionID())
		case info := <-s.sign:
			logger.Sugar().Infof("proxy->sign TransactionID: %v start", info.GetTransactionID())
			if err := s.signServer.Send(&sphinxproxy.ProxySignRequest{
				TransactionType: info.GetTransactionType(),
				CoinType:        info.GetCoinType(),
				TransactionID:   info.GetTransactionID(),
				Message:         info.GetMessage(),
			}); err != nil {
				logger.Sugar().Errorf(
					"proxy->sign TransactionID: %v TransactionType %v, CoinType: %v Message: %v error: %v",
					info.GetTransactionID(),
					info.GetTransactionType(),
					info.GetCoinType(),
					info.GetMessage(),
					err,
				)
				if checkCode(err) {
					s.closeChannel <- struct{}{}
					break send
				}
				continue
			}
			logger.Sugar().Infof("proxy->sign TransactionID: %v ok", info.GetTransactionID())
		}
	}
}

func (s *mSign) signStreamRecv(wg *sync.WaitGroup) {
	defer wg.Done()
	for !s.closeFlag {
		ssResponse, err := s.signServer.Recv()
		if err != nil {
			logger.Sugar().Errorf(
				"proxy->sign error: %v",
				err,
			)
			if checkCode(err) {
				s.closeChannel <- struct{}{}
				break
			}
			continue
		}

		logger.Sugar().Infof(
			"proxy->sign recv TransactionID: %v TransactionType %v, CoinType: %v",
			ssResponse.GetTransactionID(),
			ssResponse.GetTransactionType(),
			ssResponse.GetCoinType(),
		)

		switch ssResponse.GetTransactionType() {
		case sphinxproxy.TransactionType_WalletNew:
			if ch, ok := walletDoneChannel.Load(ssResponse.GetTransactionID()); ok {
				ch.(chan walletDoneInfo) <- walletDoneInfo{
					success: true,
					address: ssResponse.GetInfo().GetAddress(),
				}
			}
		case sphinxproxy.TransactionType_Signature:
			pluginProxy, err := getProxyPlugin(ssResponse.GetCoinType())
			if err != nil {
				logger.Sugar().Error("proxy->plugin no invalid connection")
				continue
			}
			pluginProxy.mpoolPush <- &sphinxproxy.ProxyPluginRequest{
				CoinType:        ssResponse.GetCoinType(),
				TransactionType: sphinxproxy.TransactionType_Broadcast,
				TransactionID:   ssResponse.GetTransactionID(),
				Message:         ssResponse.GetInfo().GetMessage(),
				Signature:       ssResponse.GetInfo().GetSignature(),
			}
		}
	}
}

func (s *mSign) close(wg *sync.WaitGroup) {
	defer wg.Done()
	<-s.closeChannel
	slk.Lock()
	// cancel recv transaction
	s.closeFlag = true
	nlmSign := make([]*mSign, 0, len(lmSign))
	for _, sign := range lmSign {
		if sign.signServer == s.signServer {
			logger.Sugar().Infof("some sign client closed, proxy remove it")
			continue
		}
		nlmSign = append(nlmSign, sign)
	}
	lmSign = nlmSign
	slk.Unlock()
}

type mPlugin struct {
	pluginServer sphinxproxy.SphinxProxy_ProxyPluginServer
	coinType     sphinxplugin.CoinType
	closeFlag    bool
	closeChannel chan struct{}
	balance      chan *sphinxproxy.ProxyPluginRequest
	preSign      chan *sphinxproxy.ProxyPluginRequest
	mpoolPush    chan *sphinxproxy.ProxyPluginRequest
	syncMsg      chan *sphinxproxy.ProxyPluginRequest
	registerCoin chan *sphinxproxy.ProxyPluginResponse
}

func newPluginStream(stream sphinxproxy.SphinxProxy_ProxyPluginServer) {
	lp := &mPlugin{
		pluginServer: stream,
		closeFlag:    false,
		closeChannel: make(chan struct{}),
		balance:      make(chan *sphinxproxy.ProxyPluginRequest, channelBufSize),
		preSign:      make(chan *sphinxproxy.ProxyPluginRequest, channelBufSize),
		mpoolPush:    make(chan *sphinxproxy.ProxyPluginRequest, channelBufSize),
		syncMsg:      make(chan *sphinxproxy.ProxyPluginRequest, channelBufSize),
		registerCoin: make(chan *sphinxproxy.ProxyPluginResponse),
	}
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go lp.pluginStreamSend(wg)
	go lp.pluginStreamRecv(wg)
	go lp.close(wg)
	wg.Wait()
	logger.Sugar().Info("some plugin client down, close it")
}

// add new coin type
func (lp lmPluginType) append(coinType sphinxplugin.CoinType, lmp *mPlugin) {
	plk.Lock()
	defer plk.Unlock()
	lmp.coinType = coinType
	if _, ok := lp[coinType]; !ok {
		lp[coinType] = append(lp[coinType], lmp)
	} else {
		exist := false
		for _, info := range lp[coinType] {
			if info.pluginServer == lmp.pluginServer {
				exist = true
				break
			}
		}
		if !exist {
			lp[coinType] = append(lp[coinType], lmp)
		}
	}
}

func (p *mPlugin) pluginStreamSend(wg *sync.WaitGroup) {
	defer wg.Done()
	for !p.closeFlag {
		var request *sphinxproxy.ProxyPluginRequest
		select {
		case info := <-p.balance:
			if err := p.pluginServer.Send(&sphinxproxy.ProxyPluginRequest{
				CoinType:        info.GetCoinType(),
				TransactionType: info.GetTransactionType(),
				TransactionID:   info.GetTransactionID(),
				Address:         info.GetAddress(),
			}); err != nil {
				logger.Sugar().Errorf(
					"proxy->plugin TransactionID: %v TransactionType: %v CoinType: %v Address: %v error: %v",
					info.GetTransactionID(),
					info.GetTransactionType(),
					info.GetCoinType(),
					info.GetAddress(),
					err,
				)
				ch, ok := balanceDoneChannel.Load(info.GetTransactionID())
				if !ok {
					logger.Sugar().Warnf("TransactionID: %v Addr: %v get balance maybe timeout", info.GetTransactionID(), info.GetAddress())
				}

				ch.(chan balanceDoneInfo) <- balanceDoneInfo{
					success: false,
					message: fmt.Sprintf("proxy->plugin send get wallet balance error: %v", err),
				}

				if checkCode(err) {
					p.closeChannel <- struct{}{}
					break
				}
			}
			logger.Sugar().Infof("proxy->plugin TransactionID: %v Addr: %v ok", info.GetTransactionID(), info.GetAddress())
			continue
		case info := <-p.preSign:
			request = &sphinxproxy.ProxyPluginRequest{
				CoinType:        info.GetCoinType(),
				TransactionType: info.GetTransactionType(),
				TransactionID:   info.GetTransactionID(),
				Address:         info.GetAddress(),
				Message:         info.GetMessage(),
			}
		case info := <-p.mpoolPush:
			request = &sphinxproxy.ProxyPluginRequest{
				CoinType:        info.GetCoinType(),
				TransactionType: info.GetTransactionType(),
				TransactionID:   info.GetTransactionID(),
				Message:         info.GetMessage(),
				Signature:       info.GetSignature(),
			}
		case info := <-p.syncMsg:
			request = &sphinxproxy.ProxyPluginRequest{
				CoinType:        info.GetCoinType(),
				TransactionType: info.GetTransactionType(),
				TransactionID:   info.GetTransactionID(),
				Message:         info.GetMessage(),
				CID:             info.GetCID(),
			}
		}
		if err := p.pluginServer.Send(request); err != nil {
			logger.Sugar().Errorf(
				"proxy->plugin TransactionID: %v TransactionType: %v CoinType: %v Address: %v error: %v",
				request.GetTransactionID(),
				request.GetTransactionType(),
				request.GetCoinType(),
				request.GetAddress(),
				err,
			)
			if checkCode(err) {
				p.closeChannel <- struct{}{}
				break
			}
		}
		logger.Sugar().Infof("proxy->plugin TransactionType: %v TransactionID: %v Addr: %v ok", request.GetTransactionType(), request.GetTransactionID(), request.GetAddress())
	}
}

func (p *mPlugin) pluginStreamRecv(wg *sync.WaitGroup) {
	defer wg.Done()
	for !p.closeFlag {
		psResponse, err := p.pluginServer.Recv()
		if err != nil {
			logger.Sugar().Errorf(
				"proxy->plugin error: %v",
				err,
			)
			if checkCode(err) {
				p.closeChannel <- struct{}{}
				break
			}
			continue
		}

		switch psResponse.GetTransactionType() {
		case sphinxproxy.TransactionType_RegisterCoin:
			lmPlugin.append(psResponse.GetCoinType(), p)
			if err := registerCoin(&coininfo.CreateCoinInfoRequest{
				Name: utils.TruncateCoinTypePrefix(psResponse.GetCoinType()),
				ENV:  psResponse.GetENV(),
				Unit: psResponse.GetUnit(),
			}); err != nil {
				logger.Sugar().Infof("plugin register new coin: %v error: %v", psResponse.GetCoinType(), err)
				continue
			}
			logger.Sugar().Infof("plugin register new coin: %v ok", psResponse.GetCoinType())
		case sphinxproxy.TransactionType_Balance:
			v, exact := unit.AttoFIL2FIL(psResponse.GetBalance())
			if !exact {
				logger.Sugar().Warnf("AttoFIL2FIL balance from->to(%v->%v) not exact", psResponse.GetBalance(), v)
			}

			ch, ok := balanceDoneChannel.Load(psResponse.GetTransactionID())
			if !ok {
				logger.Sugar().Warnf("TransactionID: %v Addr: %v get balance maybe timeout", psResponse.GetTransactionID(), psResponse)
				continue
			}
			ch.(chan balanceDoneInfo) <- balanceDoneInfo{
				success:    true,
				balance:    v,
				balanceStr: psResponse.GetBalanceStr(),
			}
			logger.Sugar().Infof("TransactionID: %v Addr: %v get balance ok", psResponse.GetTransactionID(), psResponse)
		case sphinxproxy.TransactionType_PreSign:
			if err := crud.UpdateTransaction(context.Background(), crud.UpdateTransactionParams{
				TransactionID: psResponse.GetTransactionID(),
				State:         sphinxproxy.TransactionState_TransactionStateSign,
				Nonce:         psResponse.GetNonce(),
			}); err != nil {
				logger.Sugar().Infof("TransactionID: %v get nonce: %v error: %v", psResponse.GetTransactionID(), psResponse.GetNonce(), err)
				continue
			}
			logger.Sugar().Infof("TransactionID: %v get nonce: %v ok", psResponse.GetTransactionID(), psResponse.GetNonce())
		case sphinxproxy.TransactionType_Broadcast:
			if err := crud.UpdateTransaction(context.Background(), crud.UpdateTransactionParams{
				TransactionID: psResponse.GetTransactionID(),
				State:         sphinxproxy.TransactionState_TransactionStateSync,
				Cid:           psResponse.GetCID(),
			}); err != nil {
				logger.Sugar().Infof("Broadcast TransactionID: %v error: %v", psResponse.GetTransactionID(), err)
				continue
			}
			logger.Sugar().Infof("Broadcast TransactionID: %v message ok", psResponse.GetTransactionID())
		case sphinxproxy.TransactionType_SyncMsgState:
			_state := sphinxproxy.TransactionState_TransactionStateFail
			if psResponse.GetExitCode() == int64(exitcode.Ok) {
				_state = sphinxproxy.TransactionState_TransactionStateDone
			}
			if err := crud.UpdateTransaction(context.Background(), crud.UpdateTransactionParams{
				TransactionID: psResponse.GetTransactionID(),
				State:         _state,
				ExitCode:      psResponse.GetExitCode(),
			}); err != nil {
				logger.Sugar().Infof("SyncMsgState TransactionID: %v error: %v", psResponse.GetTransactionID(), err)
				continue
			}
			logger.Sugar().Infof("SyncMsgState TransactionID: %v ExitCode: %v message ok", psResponse.GetTransactionID(), psResponse.GetExitCode())
		}
	}
}

func (p *mPlugin) close(wg *sync.WaitGroup) {
	defer wg.Done()
	<-p.closeChannel
	plk.Lock()
	p.closeFlag = true
	nlmPlugin := make([]*mPlugin, 0, len(lmPlugin[p.coinType]))
	for _, plugin := range lmPlugin[p.coinType] {
		if plugin.pluginServer == p.pluginServer {
			logger.Sugar().Info("some plugin client closed, proxy remove it")
			continue
		}
		nlmPlugin = append(nlmPlugin, plugin)
	}
	lmPlugin[p.coinType] = nlmPlugin
	plk.Unlock()
}

// Transaction from db fetch transactions
// multi wallet can parall
func Transaction() {
	for range time.NewTicker(constant.TaskDuration).C {
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
						logger.Sugar().Error("proxy->plugin no invalid connection")
						continue
					}
					pluginProxy.preSign <- &sphinxproxy.ProxyPluginRequest{
						CoinType:        coinType,
						TransactionType: sphinxproxy.TransactionType_PreSign,
						TransactionID:   tran.TransactionID,
						Address:         tran.From,
					}
				case uint8(sphinxproxy.TransactionState_TransactionStateSign):
					// sign -> broadcast
					signProxy, err := getProxySign()
					if err != nil {
						logger.Sugar().Error("proxy->sign no invalid connection")
						continue
					}
					coinType := sphinxplugin.CoinType(tran.CoinType)

					signProxy.sign <- &sphinxproxy.ProxySignRequest{
						TransactionType: sphinxproxy.TransactionType_Signature,
						CoinType:        coinType,
						TransactionID:   tran.TransactionID,
						Message: &sphinxplugin.UnsignedMessage{
							To:    tran.To,
							From:  tran.From,
							Nonce: tran.Nonce,
							Value: price.DBPriceToVisualPrice(tran.Amount),
							// TODO from chain get
							GasLimit:   200000000,
							GasFeeCap:  10000000,
							GasPremium: 1000000,
							Method:     uint64(builtin.MethodSend),
						},
					}
				case uint8(sphinxproxy.TransactionState_TransactionStateSync):
					coinType := sphinxplugin.CoinType(tran.CoinType)
					pluginProxy, err := getProxyPlugin(coinType)
					if err != nil {
						logger.Sugar().Error("proxy->plugin no invalid connection")
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
