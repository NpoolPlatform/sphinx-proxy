package api

import (
	"fmt"
	"sync"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/message/npool/sphinxproxy"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/utils"
)

type mPlugin struct {
	pluginServer sphinxproxy.SphinxProxy_ProxyPluginServer
	coinType     sphinxplugin.CoinType
	exitChan     chan struct{}
	once         sync.Once
	closeChan    chan struct{}
	balance      chan *sphinxproxy.ProxyPluginRequest
	registerCoin chan *sphinxproxy.ProxyPluginResponse
}

func newPluginStream(stream sphinxproxy.SphinxProxy_ProxyPluginServer) {
	lp := &mPlugin{
		pluginServer: stream,
		exitChan:     make(chan struct{}),
		closeChan:    make(chan struct{}),
		balance:      make(chan *sphinxproxy.ProxyPluginRequest, channelBufSize),
		registerCoin: make(chan *sphinxproxy.ProxyPluginResponse),
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

	for {
		select {
		case <-p.exitChan:
			return
		case info := <-p.balance:
			if err := p.pluginServer.Send(info); err != nil {
				logger.Sugar().Errorf(
					"proxy->plugin TransactionID: %v CoinType: %v error: %v",
					info.GetTransactionID(),
					info.GetCoinType(),
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
			logger.Sugar().Infof(
				"proxy->plugin TransactionID: %v Addr: %v ok",
				info.GetTransactionID(),
				info.GetAddress(),
			)
			continue
		}
	}
}

func (p *mPlugin) pluginStreamRecv(wg *sync.WaitGroup) {
	defer wg.Done()
	defer logger.Sugar().Warn("sphinx plugin stream recv exit")
	for {
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
				success: true,
				payload: psResponse.GetPayload(),
			}
			logger.Sugar().Infof("TransactionID: %v get balance ok", psResponse.GetTransactionID())
		}
	}
}

func (p *mPlugin) watch(wg *sync.WaitGroup) {
	defer wg.Done()
	defer logger.Sugar().Warn("sphinx plugin stream watch exit")

	<-p.closeChan
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

	// current conn exit
	p.once.Do(func() {
		close(p.exitChan)
	})
}

// nolint
// func isErrGasLow(msg string) bool {
// 	if msg == "" {
// 		return false
// 	}
// 	// messagepool.ErrGasFeeCapTooLow
// 	// messagepool.go:76
// 	// messagepool.go:884
// 	return regexp.MustCompile(
// 		`minimum expected nonce is [0-9]{1,}: message nonce too low`,
// 	).MatchString(msg)
// }
