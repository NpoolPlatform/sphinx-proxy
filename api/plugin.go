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
	pluginSN     string
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
	logger.Sugar().Infof("some plugin(%v) client down, close it", lp.pluginSN)
}

// add new coin type
func (lp lmPluginType) append(coinType sphinxplugin.CoinType, pluginSN string, lmp *mPlugin) {
	plk.Lock()
	defer plk.Unlock()
	lmp.coinType = coinType
	lmp.pluginSN = pluginSN
	if _, ok := lp[coinType]; !ok {
		goto appendCoinType
	} else {
		exist := false
		for _, info := range lp[coinType] {
			if info.pluginServer == lmp.pluginServer {
				exist = true
				break
			}
		}
		if !exist {
			goto appendCoinType
		}
	}
	return
appendCoinType:
	{
		lp[coinType] = append(lp[coinType], lmp)
		logger.Sugar().Infof("plugin %v, CoinType: %v is registered", pluginSN, coinType)
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
					"proxy->plugin %v: TransactionID: %v ,CoinType:  %v, error: %v",
					p.pluginSN,
					info.GetTransactionID(),
					info.GetCoinType(),
					err,
				)
				ch, ok := balanceDoneChannel.Load(info.GetTransactionID())
				if !ok {
					logger.Sugar().Warnf(
						"%v: TransactionID: %v ,Addr: %v get balance maybe timeout",
						p.pluginSN,
						info.GetTransactionID(),
						info.GetAddress(),
					)
				}

				ch.(chan balanceDoneInfo) <- balanceDoneInfo{
					success: false,
					message: fmt.Sprintf("proxy->plugin %v: ,send get wallet balance error: %v", p.pluginSN, err),
				}

				if checkCode(err) {
					p.closeChan <- struct{}{}
					return
				}
			}
			logger.Sugar().Infof(
				"proxy->plugin %v: TransactionID: %v ,Addr: %v ok",
				p.pluginSN,
				info.GetTransactionID(),
				info.GetAddress(),
			)
			continue
		}
	}
}

func (p *mPlugin) pluginStreamRecv(wg *sync.WaitGroup) {
	defer wg.Done()
	defer logger.Sugar().Warnf("%v: sphinx plugin stream recv exit", p.pluginSN)
	for {
		psResponse, err := p.pluginServer.Recv()
		if err != nil {
			logger.Sugar().Errorf(
				"proxy->plugin %v: ,error: %v",
				p.pluginSN,
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
			lmPlugin.append(psResponse.GetCoinType(), psResponse.PluginSerialNumber, p)

			if ok, err := haveCoin(utils.TruncateCoinTypePrefix(psResponse.GetCoinType())); ok && err == nil {
				continue
			}

			if err := registerCoin(&coininfo.CreateCoinInfoRequest{
				Name: utils.TruncateCoinTypePrefix(psResponse.GetCoinType()),
				ENV:  psResponse.GetENV(),
				Unit: psResponse.GetUnit(),
			}); err != nil {
				logger.Sugar().Infof(
					"plugin %v: register new coin: %v, error: %v",
					psResponse.PluginSerialNumber,
					psResponse.GetCoinType(),
					err,
				)
				continue
			}
			logger.Sugar().Infof("plugin: %v ,register new coin: %v ok,", psResponse.PluginSerialNumber, psResponse.GetCoinType())
		case sphinxproxy.TransactionType_Balance:
			ch, ok := balanceDoneChannel.Load(psResponse.GetTransactionID())
			if !ok {
				logger.Sugar().Warnf("TransactionID: %v get balance maybe timeout at %v", psResponse.GetTransactionID(), p.pluginSN)
				continue
			}

			if psResponse.GetRPCExitMessage() != "" {
				logger.Sugar().Infof(
					"%v: TransactionID: %v, get balance error: %v",
					p.pluginSN,
					psResponse.GetTransactionID(),
					psResponse.GetRPCExitMessage(),
				)
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
			logger.Sugar().Infof("%v: TransactionID: %v get balance ok", p.pluginSN, psResponse.GetTransactionID())
		}
	}
}

func (p *mPlugin) watch(wg *sync.WaitGroup) {
	defer wg.Done()
	defer logger.Sugar().Warnf("sphinx plugin %v stream watch exit", p.pluginSN)

	<-p.closeChan
	plk.Lock()
	nlmPlugin := make([]*mPlugin, 0, len(lmPlugin[p.coinType]))
	for _, plugin := range lmPlugin[p.coinType] {
		if plugin.pluginServer == p.pluginServer {
			logger.Sugar().Infof("some plugin %v client closed, proxy remove it", p.pluginSN)
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
