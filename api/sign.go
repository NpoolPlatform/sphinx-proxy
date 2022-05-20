package api

import (
	"fmt"
	"sync"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/sphinxproxy"
)

type mSign struct {
	signServer sphinxproxy.SphinxProxy_ProxySignServer
	// pod exit
	exitChan chan struct{}
	// conn error
	connCloseChan chan struct{}
	closeChan     chan struct{}
	once          sync.Once
	walletNew     chan *sphinxproxy.ProxySignRequest
	sign          chan *sphinxproxy.ProxySignRequest
}

func newSignStream(stream sphinxproxy.SphinxProxy_ProxySignServer) {
	lc := &mSign{
		signServer:    stream,
		exitChan:      make(chan struct{}),
		connCloseChan: make(chan struct{}),
		closeChan:     make(chan struct{}),
		walletNew:     make(chan *sphinxproxy.ProxySignRequest, channelBufSize),
		sign:          make(chan *sphinxproxy.ProxySignRequest, channelBufSize),
	}
	slk.Lock()
	lmSign = append(lmSign, lc)
	slk.Unlock()

	wg := &sync.WaitGroup{}
	wg.Add(3)
	go lc.signStreamSend(wg)
	go lc.signStreamRecv(wg)
	go lc.watch(wg)
	wg.Wait()
	logger.Sugar().Info("some sign client down, close it")
}

func (s *mSign) signStreamSend(wg *sync.WaitGroup) {
	defer wg.Done()
	defer logger.Sugar().Warn("sphinx sign stream send exit")

	for {
		select {
		case <-s.exitChan:
			return
		case <-s.connCloseChan:
			return
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
					s.closeChan <- struct{}{}
					return
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
					s.closeChan <- struct{}{}
					return
				}
				continue
			}
			logger.Sugar().Infof("proxy->sign TransactionID: %v ok", info.GetTransactionID())
		}
	}
}

func (s *mSign) signStreamRecv(wg *sync.WaitGroup) {
	defer wg.Done()
	defer logger.Sugar().Warn("sphinx sign stream recv exit")

	for {
		select {
		case <-s.exitChan:
			return
		case <-s.connCloseChan:
			return
		default:
			ssResponse, err := s.signServer.Recv()
			if err != nil {
				logger.Sugar().Errorf(
					"proxy->sign error: %v",
					err,
				)
				if checkCode(err) {
					s.closeChan <- struct{}{}
					return
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
				ch, ok := walletDoneChannel.Load(ssResponse.GetTransactionID())
				if !ok {
					logger.Sugar().Warnf("TransactionID: %v create wallet maybe timeout", ssResponse.GetTransactionID())
					continue
				}

				if ssResponse.GetRPCExitMessage() != "" {
					logger.Sugar().Infof("TransactionID: %v create wallet error: %v", ssResponse.GetTransactionID(), ssResponse.GetRPCExitMessage())
					ch.(chan walletDoneInfo) <- walletDoneInfo{
						success: false,
						message: ssResponse.GetRPCExitMessage(),
					}
					continue
				}

				ch.(chan walletDoneInfo) <- walletDoneInfo{
					success: true,
					address: ssResponse.GetInfo().GetAddress(),
				}
				logger.Sugar().Infof("TransactionID: %v create wallet ok", ssResponse.GetTransactionID())
			case sphinxproxy.TransactionType_Signature:
				pluginProxy, err := getProxyPlugin(ssResponse.GetCoinType())
				if err != nil {
					logger.Sugar().Errorf("proxy->plugin no valid connection for coin: %v transaction: %v",
						ssResponse.GetCoinType(),
						ssResponse.GetTransactionID(),
					)
					continue
				}
				if ssResponse.GetRPCExitMessage() != "" {
					logger.Sugar().Errorf("proxy->sign signature for coin: %v transaction: %v error: %v",
						ssResponse.GetCoinType(),
						ssResponse.GetTransactionID(),
						ssResponse.GetRPCExitMessage(),
					)
					continue
				}
				pluginProxy.mpoolPush <- &sphinxproxy.ProxyPluginRequest{
					CoinType:        ssResponse.GetCoinType(),
					TransactionType: sphinxproxy.TransactionType_Broadcast,
					TransactionID:   ssResponse.GetTransactionID(),
					// fil/sol
					Message:   ssResponse.GetInfo().GetMessage(),
					Signature: ssResponse.GetInfo().GetSignature(),
					// btc
					MsgTx: ssResponse.GetMsgTx(),
					// eth/er20
					SignedRawTxHex: ssResponse.GetSignedRawTxHex(),
				}
			}
		}
	}
}

func (s *mSign) watch(wg *sync.WaitGroup) {
	defer wg.Done()
	defer logger.Sugar().Warn("sphinx sign stream watch exit")

	select {
	case <-s.exitChan:
		return
	case <-s.closeChan:
		slk.Lock()
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
		s.once.Do(func() {
			close(s.connCloseChan)
		})
	}
}
