package api

import (
	"fmt"
	"sync"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	coinpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"
	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/message/npool/sphinxproxy"
	putils "github.com/NpoolPlatform/sphinx-plugin/pkg/rpc"
)

type mPlugin struct {
	pluginServer  sphinxproxy.SphinxProxy_ProxyPluginServer
	coinType      sphinxplugin.CoinType
	pluginInfo    string
	exitChan      chan struct{}
	connCloseChan chan struct{}
	once          sync.Once
	closeChan     chan struct{}
	pluginReq     chan *sphinxproxy.ProxyPluginRequest
	registerCoin  chan *sphinxproxy.ProxyPluginResponse
}

func newPluginStream(stream sphinxproxy.SphinxProxy_ProxyPluginServer) {
	lp := &mPlugin{
		pluginServer:  stream,
		exitChan:      make(chan struct{}),
		connCloseChan: make(chan struct{}),
		closeChan:     make(chan struct{}),
		pluginReq:     make(chan *sphinxproxy.ProxyPluginRequest, channelBufSize),
		registerCoin:  make(chan *sphinxproxy.ProxyPluginResponse),
	}
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go lp.pluginStreamSend(wg)
	go lp.pluginStreamRecv(wg)
	go lp.watch(wg)
	wg.Wait()
	logger.Sugar().Infof("some plugin(%v) client down, close it", lp.pluginInfo)
}

// add new coin type
func (lp lmPluginType) append(coinType sphinxplugin.CoinType, pluginInfo string, lmp *mPlugin) (exist bool) {
	plk.Lock()
	defer plk.Unlock()
	lmp.coinType = coinType
	lmp.pluginInfo = pluginInfo
	if _, ok := lp[coinType]; ok {
		for _, info := range lp[coinType] {
			if info.pluginServer == lmp.pluginServer {
				return true
			}
		}
	}
	lp[coinType] = append(lp[coinType], lmp)
	logger.Sugar().Infof("plugin %v %v is added", pluginInfo, coinType)
	return false
}

func (p *mPlugin) pluginStreamSend(wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() { logger.Sugar().Warn("sphinx plugin stream send exit") }()

	for {
		select {
		case <-p.exitChan:
			logger.Sugar().Info("plugin send chan exit")
			return
		case <-p.connCloseChan:
			logger.Sugar().Info("plugin send chan close")
			return
		case info := <-p.pluginReq:
			if err := p.pluginServer.Send(info); err != nil {
				logger.Sugar().Errorf(
					"proxy->plugin %v %v, TransactionType: %v, TransactionID: %v , error: %v",
					p.pluginInfo,
					info.GetName(),
					info.GetTransactionType(),
					info.GetTransactionID(),
					err,
				)

				switch info.GetTransactionType() {
				case sphinxproxy.TransactionType_Balance:
					ch, ok := balanceDoneChannel.Load(info.GetTransactionID())
					if !ok {
						logger.Sugar().Warnf(
							"%v %v: TransactionType: %v, TransactionID: %v, req maybe timeout",
							p.pluginInfo,
							info.GetName(),
							info.GetTransactionType(),
							info.GetTransactionID(),
						)
					}

					ch.(chan balanceDoneInfo) <- balanceDoneInfo{
						success: false,
						message: "send request to plugin error",
					}
				case sphinxproxy.TransactionType_EstimateGas:
					ch, ok := esGasDoneChannel.Load(info.GetTransactionID())
					if !ok {
						logger.Sugar().Warnf(
							"%v %v: TransactionType: %v, TransactionID: %v, req maybe timeout",
							p.pluginInfo,
							info.GetName(),
							info.GetTransactionType(),
							info.GetTransactionID(),
						)
					}

					ch.(chan esGasDoneInfo) <- esGasDoneInfo{
						success: false,
						message: "send request to plugin error",
					}
				}

				if putils.CheckCode(err) {
					p.closeChan <- struct{}{}
					return
				}
			}
			logger.Sugar().Infof(
				"proxy->plugin %v %v: TransactionType: %v, TransactionID: %v",
				p.pluginInfo,
				info.GetName(),
				info.GetTransactionType(),
				info.GetTransactionID(),
			)
			continue
		}
	}
}

func (p *mPlugin) pluginStreamRecv(wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() { logger.Sugar().Warnf("%v: sphinx plugin stream recv exit", p.pluginInfo) }()
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
				if putils.CheckCode(err) {
					p.closeChan <- struct{}{}
					return
				}
				continue
			}

			switch psResponse.GetTransactionType() {
			case sphinxproxy.TransactionType_RegisterCoin:
				pluginInfo := fmt.Sprintf("%v-%v", psResponse.PluginPosition, psResponse.PluginWanIP)
				exist := lmPlugin.append(psResponse.GetCoinType(), pluginInfo, p)
				registered, err := haveCoin(psResponse.Name)

				if exist && registered && err == nil {
					continue
				}

				chainType := psResponse.ChainType.String()
				if err := registerCoin(&coinpb.CoinReq{
					Name:            &psResponse.Name,
					Unit:            &psResponse.Unit,
					ENV:             &psResponse.ENV,
					ChainType:       &chainType,
					ChainNativeUnit: &psResponse.ChainNativeUnit,
					ChainAtomicUnit: &psResponse.ChainAtomicUnit,
					ChainUnitExp:    &psResponse.ChainUnitExp,
					ChainID:         &psResponse.ChainID,
					ChainNickname:   &psResponse.ChainNickname,
					GasType:         &psResponse.GasType,
				}); err != nil {
					logger.Sugar().Infof(
						"plugin %v: register new coin: %v, error: %v",
						pluginInfo,
						psResponse.GetName(),
						err,
					)
					continue
				}
				logger.Sugar().Infof("plugin: %v ,register new coin: %v ok,", pluginInfo, psResponse.GetName())
			case sphinxproxy.TransactionType_Balance:
				ch, ok := balanceDoneChannel.Load(psResponse.GetTransactionID())
				if !ok {
					logger.Sugar().Warnf("%v %v :TransactionID: %v get balance maybe timeout", p.pluginInfo, psResponse.GetCoinType(), psResponse.GetTransactionID())
					continue
				}

				if psResponse.GetRPCExitMessage() != "" {
					logger.Sugar().Infof(
						"%v: TransactionID: %v, get balance error: %v",
						p.pluginInfo,
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
				logger.Sugar().Infof("%v %v : TransactionID: %v get balance ok", p.pluginInfo, psResponse.GetCoinType(), psResponse.GetTransactionID())

			case sphinxproxy.TransactionType_EstimateGas:
				ch, ok := esGasDoneChannel.Load(psResponse.GetTransactionID())
				if !ok {
					logger.Sugar().Warnf("%v %v :TransactionID: %v estimate gas maybe timeout", p.pluginInfo, psResponse.GetCoinType(), psResponse.GetTransactionID())
					continue
				}

				if psResponse.GetRPCExitMessage() != "" {
					logger.Sugar().Infof(
						"%v: TransactionID: %v, estimate gas error: %v",
						p.pluginInfo,
						psResponse.GetTransactionID(),
						psResponse.GetRPCExitMessage(),
					)
					ch.(chan esGasDoneInfo) <- esGasDoneInfo{
						success: false,
						message: psResponse.GetRPCExitMessage(),
					}
					continue
				}

				ch.(chan esGasDoneInfo) <- esGasDoneInfo{
					success: true,
					payload: psResponse.GetPayload(),
				}
				logger.Sugar().Infof("%v %v : TransactionID: %v estimate gas ok", p.pluginInfo, psResponse.GetCoinType(), psResponse.GetTransactionID())
			}
		}
	}
}

func (p *mPlugin) watch(wg *sync.WaitGroup) {
	defer wg.Done()
	defer logger.Sugar().Warnf("sphinx plugin %v stream watch exit", p.pluginInfo)

	select {
	case <-p.exitChan:
		logger.Sugar().Info("plugin %v watch chan exit", p.pluginInfo)
		return
	case <-p.closeChan:
		logger.Sugar().Info("plugin watch chan close")
		plk.Lock()
		nlmPlugin := make([]*mPlugin, 0, len(lmPlugin[p.coinType]))
		for _, plugin := range lmPlugin[p.coinType] {
			if plugin.pluginServer == p.pluginServer {
				logger.Sugar().Infof("some plugin %v client closed, proxy remove it", p.pluginInfo)
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
