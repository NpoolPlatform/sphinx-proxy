package api

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/message/npool/sphinxproxy"
	"github.com/NpoolPlatform/sphinx-plugin/pkg/plugin/eth"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/crud"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/utils"
)

type mPlugin struct {
	pluginServer sphinxproxy.SphinxProxy_ProxyPluginServer
	coinType     sphinxplugin.CoinType
	// pod exit
	exitChan chan struct{}
	// conn error
	connCloseChan chan struct{}
	closeChan     chan struct{}
	once          sync.Once
	balance       chan *sphinxproxy.ProxyPluginRequest
	preSign       chan *sphinxproxy.ProxyPluginRequest
	mpoolPush     chan *sphinxproxy.ProxyPluginRequest
	syncMsg       chan *sphinxproxy.ProxyPluginRequest
	registerCoin  chan *sphinxproxy.ProxyPluginResponse
}

func newPluginStream(stream sphinxproxy.SphinxProxy_ProxyPluginServer) {
	lp := &mPlugin{
		pluginServer:  stream,
		exitChan:      make(chan struct{}),
		connCloseChan: make(chan struct{}),
		closeChan:     make(chan struct{}),
		balance:       make(chan *sphinxproxy.ProxyPluginRequest, channelBufSize),
		preSign:       make(chan *sphinxproxy.ProxyPluginRequest, channelBufSize),
		mpoolPush:     make(chan *sphinxproxy.ProxyPluginRequest, channelBufSize),
		syncMsg:       make(chan *sphinxproxy.ProxyPluginRequest, channelBufSize),
		registerCoin:  make(chan *sphinxproxy.ProxyPluginResponse),
	}
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go lp.pluginStreamSend(wg)
	go lp.pluginStreamRecv(wg)
	go lp.watch(wg)
	wg.Wait()
	logger.Sugar().Info("some plugin client down, close it")
}

// add new coin type
func (lp lmPluginType) append(coinType sphinxplugin.CoinType, lmp *mPlugin) {
	plk.Lock()
	defer plk.Unlock()
	lmp.coinType = coinType

	logger.Sugar().Infof("some plugin %v", coinType)

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
	defer logger.Sugar().Warn("sphinx plugin stream send exit")

	var request *sphinxproxy.ProxyPluginRequest
	for {
		select {
		case <-p.exitChan:
			logger.Sugar().Info("plugin send chan exit")
			return
		case <-p.connCloseChan:
			logger.Sugar().Info("plugin send chan close")
			return
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
					p.closeChan <- struct{}{}
					return
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
				MsgTx:           info.GetMsgTx(),
				SignedRawTxHex:  info.GetSignedRawTxHex(),
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
				p.closeChan <- struct{}{}
				return
			}
		}
		logger.Sugar().Infof("proxy->plugin TransactionType: %v TransactionID: %v Addr: %v ok", request.GetTransactionType(), request.GetTransactionID(), request.GetAddress())
	}
}

// nolint
func (p *mPlugin) pluginStreamRecv(wg *sync.WaitGroup) {
	defer wg.Done()
	defer logger.Sugar().Warn("sphinx plugin stream recv exit")
	for {
		select {
		case <-p.exitChan:
			logger.Sugar().Info("plugin recv chan exit")
			return
		case <-p.connCloseChan:
			logger.Sugar().Info("plugin recv chan close")
			return
		default:
			psResponse, err := p.pluginServer.Recv()
			if err != nil {
				logger.Sugar().Errorf(
					"proxy->plugin error: %v",
					err,
				)
				if checkCode(err) {
					p.closeChan <- struct{}{}
					return
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
				ch, ok := balanceDoneChannel.Load(psResponse.GetTransactionID())
				if !ok {
					logger.Sugar().Warnf("TransactionID: %v get balance maybe timeout", psResponse.GetTransactionID())
					continue
				}

				if psResponse.GetRPCExitMessage() != "" {
					logger.Sugar().Infof("TransactionID: %v get balance error: %v", psResponse.GetTransactionID(), psResponse.GetRPCExitMessage())
					ch.(chan balanceDoneInfo) <- balanceDoneInfo{
						success: false,
						message: psResponse.GetRPCExitMessage(),
					}
					continue
				}

				ch.(chan balanceDoneInfo) <- balanceDoneInfo{
					success:    true,
					balance:    psResponse.GetBalance(),
					balanceStr: psResponse.GetBalanceStr(),
				}
				logger.Sugar().Infof("TransactionID: %v get balance ok", psResponse.GetTransactionID())
			case sphinxproxy.TransactionType_PreSign:
				if psResponse.GetRPCExitMessage() != "" {
					logger.Sugar().Errorf("PreSign TransactionID: %v error: %v", psResponse.GetTransactionID(), psResponse.GetRPCExitMessage())
					if isErrTRXBalanceLow(psResponse.GetRPCExitMessage()) {
						if err := crud.UpdateTransaction(context.Background(), &crud.UpdateTransactionParams{
							TransactionID: psResponse.GetTransactionID(),
							State:         sphinxproxy.TransactionState_TransactionStateFail,
						}); err != nil {
							logger.Sugar().Infof("PreSign TransactionID: %v error: %v", psResponse.GetTransactionID(), err)
						}
						continue
					}

				}

				if err := crud.UpdateTransaction(context.Background(), &crud.UpdateTransactionParams{
					TransactionID: psResponse.GetTransactionID(),
					State:         sphinxproxy.TransactionState_TransactionStateSign,
					Nonce:         psResponse.GetMessage().GetNonce(),
					RecentBhash:   psResponse.GetMessage().GetRecentBhash(),
					UTXO:          psResponse.GetMessage().GetUnspent(),
					Pre: &eth.PreSignInfo{
						ChainID:    psResponse.GetMessage().GetChainID(),
						ContractID: psResponse.GetMessage().GetContractID(),
						Nonce:      psResponse.GetMessage().GetNonce(),
						GasPrice:   psResponse.GetMessage().GetGasPrice(),
						GasLimit:   psResponse.GetMessage().GetGasLimit(),
					},
					TxData: psResponse.GetMessage().GetTxData(),
				}); err != nil {
					logger.Sugar().Infof("TransactionID: %v get nonce: %v error: %v", psResponse.GetTransactionID(), psResponse.GetMessage().GetNonce(), err)
					continue
				}
				logger.Sugar().Infof("TransactionID: %v get nonce: %v ok", psResponse.GetTransactionID(), psResponse.GetMessage().GetNonce())
			case sphinxproxy.TransactionType_Broadcast:
				state := sphinxproxy.TransactionState_TransactionStateSync
				if psResponse.GetRPCExitMessage() != "" {
					logger.Sugar().Errorf("Broadcast TransactionID: %v error: %v", psResponse.GetTransactionID(), psResponse.GetRPCExitMessage())
					if !isErrFILGasLow(psResponse.GetRPCExitMessage()) &&
						!isErrTRC20Expired(psResponse.GetRPCExitMessage()) &&
						!isErrETHFundsLow(psResponse.GetRPCExitMessage()) &&
						!isErrTRXBalanceLow(psResponse.GetRPCExitMessage()) &&
						!isErrERC20GasLow(psResponse.GetRPCExitMessage()) {
						continue
					}
					state = sphinxproxy.TransactionState_TransactionStateFail
				}

				if err := crud.UpdateTransaction(context.Background(), &crud.UpdateTransactionParams{
					TransactionID: psResponse.GetTransactionID(),
					State:         state,
					Cid:           psResponse.GetCID(),
				}); err != nil {
					logger.Sugar().Infof("Broadcast TransactionID: %v error: %v", psResponse.GetTransactionID(), err)
					continue
				}
				logger.Sugar().Infof("Broadcast TransactionID: %v message ok", psResponse.GetTransactionID())
			case sphinxproxy.TransactionType_SyncMsgState:
				if psResponse.GetRPCExitMessage() != "" {
					logger.Sugar().Infof("SyncMsgState TransactionID: %v error: %v", psResponse.GetTransactionID(), psResponse.GetRPCExitMessage())
					continue
				}

				_state := sphinxproxy.TransactionState_TransactionStateDone
				// 0: sync ok, other fail
				if psResponse.GetExitCode() != int64(0) {
					_state = sphinxproxy.TransactionState_TransactionStateFail
				}
				if err := crud.UpdateTransaction(context.Background(), &crud.UpdateTransactionParams{
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
}

func (p *mPlugin) watch(wg *sync.WaitGroup) {
	defer wg.Done()
	defer logger.Sugar().Warn("sphinx plugin stream watch exit")

	select {
	case <-p.exitChan:
		logger.Sugar().Info("plugin watch chan exit")
		return
	case <-p.closeChan:
		logger.Sugar().Info("plugin watch chan close")
		plk.Lock()
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
		p.once.Do(func() {
			close(p.connCloseChan)
		})
	}
}

func isErrFILGasLow(msg string) bool {
	if msg == "" {
		return false
	}
	// messagepool.ErrGasFeeCapTooLow
	// messagepool.go:76
	// messagepool.go:884
	return strings.Contains(
		msg,
		`gas fee cap too low`,
	)
}

func isErrERC20GasLow(msg string) bool {
	if msg == "" {
		return false
	}
	return strings.Contains(
		msg,
		`intrinsic gas too low`,
	)
}

func isErrETHFundsLow(msg string) bool {
	if msg == "" {
		return false
	}
	return strings.Contains(
		msg,
		`insufficient funds for gas * price + value`,
	)
}

func isErrTRC20Expired(msg string) bool {
	if msg == "" {
		return false
	}
	return strings.Contains(
		msg,
		`Transaction expired`,
	)
}

func isErrTRXBalanceLow(msg string) bool {
	if msg == "" {
		return false
	}
	return strings.Contains(
		msg,
		`balance is not sufficient`,
	)
}
