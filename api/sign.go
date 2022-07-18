package api

import (
	"encoding/json"
	"sync"

	putils "github.com/NpoolPlatform/sphinx-plugin/pkg/rpc"
	plugin_types "github.com/NpoolPlatform/sphinx-plugin/pkg/types"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/sphinxproxy"
)

type mSign struct {
	signServer sphinxproxy.SphinxProxy_ProxySignServer
	exitChan   chan struct{}
	// conn error
	connCloseChan chan struct{}
	once          sync.Once
	closeChan     chan struct{}
	walletNew     chan *sphinxproxy.ProxySignRequest
}

func newSignStream(stream sphinxproxy.SphinxProxy_ProxySignServer) {
	lc := &mSign{
		signServer:    stream,
		exitChan:      make(chan struct{}),
		closeChan:     make(chan struct{}),
		connCloseChan: make(chan struct{}),
		walletNew:     make(chan *sphinxproxy.ProxySignRequest, channelBufSize),
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
	defer func() { logger.Sugar().Warn("sphinx sign stream send exit") }()

	for {
		select {
		case <-s.exitChan:
			logger.Sugar().Info("sign send chan exit")
			return
		case <-s.connCloseChan:
			logger.Sugar().Info("sign send chan close")
			return
		case info := <-s.walletNew:
			logger.Sugar().Infof("proxy->sign TransactionID: %v start", info.GetTransactionID())
			if err := s.signServer.Send(info); err != nil {
				logger.Sugar().Errorf(
					"proxy->sign TransactionID: %v, CoinType: %v error: %v",
					info.GetCoinType(),
					info.GetTransactionID(),
					err,
				)
				if ch, ok := walletDoneChannel.Load(info.GetTransactionID()); ok {
					ch.(chan walletDoneInfo) <- walletDoneInfo{
						success: false,
						message: "create wallet error",
					}
				}
				if putils.CheckCode(err) {
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
	defer func() { logger.Sugar().Warn("sphinx sign stream recv exit") }()

	for {
		select {
		case <-s.exitChan:
			logger.Sugar().Info("sign recv chan exit")
			return
		case <-s.connCloseChan:
			logger.Sugar().Info("sign recv chan close")
			return
		default:
			ssResponse, err := s.signServer.Recv()
			if err != nil {
				logger.Sugar().Errorf(
					"proxy->sign error: %v",
					err,
				)
				if putils.CheckCode(err) {
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

				naResponse := &plugin_types.NewAccountResponse{}
				err = json.Unmarshal(ssResponse.GetPayload(), naResponse)
				if err != nil {
					logger.Sugar().Infof("TransactionID: %v create wallet error: %v", ssResponse.GetTransactionID(), err)
					ch.(chan walletDoneInfo) <- walletDoneInfo{
						success: false,
						message: ssResponse.GetRPCExitMessage(),
					}
					continue
				}

				ch.(chan walletDoneInfo) <- walletDoneInfo{
					success: true,
					address: naResponse.Address,
				}
				logger.Sugar().Infof("TransactionID: %v create wallet ok", ssResponse.GetTransactionID())
			default:
				logger.Sugar().Errorf("TransactionID: %v for TransactionType: %v not support", ssResponse.GetTransactionID(), ssResponse.GetTransactionType())
			}
		}
	}
}

func (s *mSign) watch(wg *sync.WaitGroup) {
	defer wg.Done()
	defer logger.Sugar().Warn("sphinx sign stream watch exit")

	select {
	case <-s.exitChan:
		logger.Sugar().Info("sign watch chan exit")
		return
	case <-s.closeChan:
		logger.Sugar().Info("sign watch chan close")
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
