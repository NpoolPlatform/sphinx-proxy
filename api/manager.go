package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/signproxy"
	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/message/npool/trading"
	constant "github.com/NpoolPlatform/sphinx-proxy/pkg/message/const"
	msgcli "github.com/NpoolPlatform/sphinx-proxy/pkg/message/message"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/unit"
	sconst "github.com/NpoolPlatform/sphinx-service/pkg/message/const"
	msgproducer "github.com/NpoolPlatform/sphinx-service/pkg/message/message"
	"github.com/filecoin-project/specs-actors/actors/builtin"
)

var (
	ErrNoSignServiceFound   = errors.New("no sign service found")
	ErrNoPluginServiceFound = errors.New("no plugin service found")
)

type lmPluginType map[sphinxplugin.CoinType][]*mPlugin

type combineProxy struct {
	mPlugin *mPlugin
	data    *msgproducer.NotificationTransaction
}

var (
	rnd                = rand.New(rand.NewSource(time.Now().UnixNano()))
	done               = make(chan struct{})
	slk                sync.RWMutex
	plk                sync.RWMutex
	lmSign             = make([]*mSign, 0)
	lmPlugin           = make(lmPluginType)
	combineProxyChannl = make(chan combineProxy, 1024)
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
	signServer signproxy.SignProxy_ProxySignServer
	close      chan struct{}
	walletNew  chan *signproxy.ProxySignRequest // address
	sign       chan *signproxy.ProxySignRequest
	ack        chan *trading.ACKRequest
}

func newSignStream(stream signproxy.SignProxy_ProxySignServer) {
	lc := &mSign{
		signServer: stream,
		close:      make(chan struct{}),
		walletNew:  make(chan *signproxy.ProxySignRequest),
		sign:       make(chan *signproxy.ProxySignRequest),
		ack:        make(chan *trading.ACKRequest),
	}
	slk.Lock()
	lmSign = append(lmSign, lc)
	slk.Unlock()
	go lc.signStreamSend()
	go lc.signStreamRecv()
	go lc.handleACK()
}

func (s *mSign) handleACK() {
	for {
		ackReq := <-s.ack
		if err := ack(ackReq); err != nil {
			logger.Sugar().Errorf(
				"handle ack info: %v error: %v",
				ackReq,
				err,
			)
		}
	}
}

// wallet new
func (s *mSign) signStreamSend() {
	for {
		select {
		case info := <-s.walletNew:
			if err := s.signServer.Send(&signproxy.ProxySignRequest{
				TransactionType:     info.GetTransactionType(),
				CoinType:            info.GetCoinType(),
				TransactionIDInsite: info.GetTransactionIDInsite(),
			}); err != nil {
				logger.Sugar().Errorf(
					"proxy->sign TransactionIDInsite: %v TransactionType %v, CoinType: %v error: %v",
					info.GetTransactionType(),
					info.GetCoinType(),
					info.GetTransactionIDInsite(),
					err,
				)
				s.ack <- &trading.ACKRequest{
					IsOkay:              false,
					TransactionType:     info.GetTransactionType(),
					TransactionIdInsite: info.GetTransactionIDInsite(),
					ErrorMessage:        fmt.Sprintf("create new account error: %v", err),
				}
			}
			logger.Sugar().Infof("proxy->sign TransactionIDInsite: %v ok", info.GetTransactionIDInsite())
		case info := <-s.sign:
			if err := s.signServer.Send(&signproxy.ProxySignRequest{
				TransactionType:     info.GetTransactionType(),
				CoinType:            info.GetCoinType(),
				TransactionIDInsite: info.GetTransactionIDInsite(),
				Message:             info.GetMessage(),
			}); err != nil {
				done <- struct{}{}
				s.ack <- &trading.ACKRequest{
					IsOkay:              false,
					TransactionType:     info.GetTransactionType(),
					TransactionIdInsite: info.GetTransactionIDInsite(),
					ErrorMessage:        fmt.Sprintf("create new account error: %v", err),
				}
				logger.Sugar().Errorf(
					"proxy->sign TransactionIDInsite: %v TransactionType %v, CoinType: %v Message: %v error: %v",
					info.GetTransactionIDInsite(),
					info.GetTransactionType(),
					info.GetCoinType(),
					info.GetMessage(),
					err,
				)
			}
		}
	}
}

func (s *mSign) signStreamRecv() {
	for {
		ssResponse, err := s.signServer.Recv()
		if err != nil {
			logger.Sugar().Errorf(
				"proxy->sign error: %v",
				err,
			)
			continue
		}

		logger.Sugar().Infof(
			"proxy->sign recv TransactionIDInsite: %v TransactionType %v, CoinType: %v",
			ssResponse.GetTransactionType(),
			ssResponse.GetCoinType(),
			ssResponse.GetTransactionIDInsite(),
		)

		switch ssResponse.GetTransactionType() {
		case signproxy.TransactionType_WalletNew:
			s.ack <- &trading.ACKRequest{
				IsOkay:              true,
				TransactionType:     ssResponse.GetTransactionType(),
				CoinTypeId:          int32(ssResponse.GetCoinType()),
				TransactionIdInsite: ssResponse.GetTransactionIDInsite(),
				Address:             ssResponse.GetInfo().GetAddress(),
			}
		case signproxy.TransactionType_Signature:
			pluginProxy, err := getProxyPlugin(ssResponse.GetCoinType())
			if err != nil {
				logger.Sugar().Error("proxy->plugin no invalid connection")
				continue
			}
			pluginProxy.mpoolPush <- &signproxy.ProxyPluginRequest{
				CoinType:            ssResponse.GetCoinType(),
				TransactionType:     ssResponse.GetTransactionType(),
				TransactionIDInsite: ssResponse.GetTransactionIDInsite(),
				Message:             ssResponse.GetInfo().GetMessage(),
				Signature:           ssResponse.GetInfo().GetSignature(),
			}
		}
	}
}

type mPlugin struct {
	pluginServer signproxy.SignProxy_ProxyPluginServer
	close        chan struct{}
	coinType     sphinxplugin.CoinType              // coin type
	balance      chan *signproxy.ProxyPluginRequest // address
	preSign      chan *signproxy.ProxyPluginRequest
	mpoolPush    chan *signproxy.ProxyPluginRequest
	registerCoin chan *signproxy.ProxyPluginResponse
	ack          chan *trading.ACKRequest
}

func newPluginStream(stream signproxy.SignProxy_ProxyPluginServer) {
	lp := &mPlugin{
		pluginServer: stream,
		close:        make(chan struct{}),
		balance:      make(chan *signproxy.ProxyPluginRequest),
		preSign:      make(chan *signproxy.ProxyPluginRequest),
		mpoolPush:    make(chan *signproxy.ProxyPluginRequest),
		registerCoin: make(chan *signproxy.ProxyPluginResponse),
		ack:          make(chan *trading.ACKRequest),
	}
	go lp.pluginStreamSend()
	go lp.pluginStreamRecv()
	go lp.handleACK()
}

func (p *mPlugin) handleACK() {
	for {
		ackReq := <-p.ack
		if err := ack(ackReq); err != nil {
			logger.Sugar().Errorf(
				"handle ack info: %v error: %v",
				ackReq,
				err,
			)
		}
	}
}

// add new coin type
func (lp lmPluginType) append(coinType sphinxplugin.CoinType, lmp mPlugin) {
	plk.Lock()
	defer plk.Unlock()
	if _, ok := lp[coinType]; !ok {
		lmp.coinType = coinType
		lp[coinType] = append(lp[coinType], &lmp)
	} else {
		for _, info := range lp[coinType] {
			if info.pluginServer != lmp.pluginServer {
				lp[coinType] = append(lp[coinType], &lmp)
				break
			}
		}
	}
}

func (p *mPlugin) pluginStreamSend() {
	for {
		select {
		case info := <-p.balance:
			if err := p.pluginServer.Send(&signproxy.ProxyPluginRequest{
				CoinType:            info.GetCoinType(),
				TransactionType:     info.GetTransactionType(),
				TransactionIDInsite: info.GetTransactionIDInsite(),
				Address:             info.GetAddress(),
			}); err != nil {
				logger.Sugar().Errorf(
					"proxy->plugin TransactionIDInsite: %v TransactionType: %v CoinType: %v Address: %v error: %v",
					info.GetTransactionIDInsite(),
					info.GetTransactionType(),
					info.GetCoinType(),
					info.GetAddress(),
					err,
				)
				p.ack <- &trading.ACKRequest{
					IsOkay:              false,
					TransactionType:     info.GetTransactionType(),
					TransactionIdInsite: info.GetTransactionIDInsite(),
					Address:             info.GetAddress(),
					ErrorMessage:        fmt.Sprintf("get wallet balance error: %v", err),
				}
			}
			logger.Sugar().Infof("proxy->plugin TransactionIDInsite: %v ok", info.GetTransactionIDInsite())
		case info := <-p.preSign:
			if err := p.pluginServer.Send(&signproxy.ProxyPluginRequest{
				CoinType:            info.GetCoinType(),
				TransactionType:     info.GetTransactionType(),
				TransactionIDInsite: info.GetTransactionIDInsite(),
				Address:             info.GetAddress(),
			}); err != nil {
				done <- struct{}{}
				logger.Sugar().Errorf(
					"proxy->plugin TransactionIDInsite: %v TransactionType %v, CoinType: %v Address: %v error: %v",
					info.GetTransactionIDInsite(),
					info.GetTransactionType(),
					info.GetCoinType(),
					info.GetAddress(),
					err,
				)
				p.ack <- &trading.ACKRequest{
					IsOkay:              false,
					TransactionType:     info.GetTransactionType(),
					TransactionIdInsite: info.GetTransactionIDInsite(),
					Address:             info.GetAddress(),
					ErrorMessage:        fmt.Sprintf("get wallet balance error: %v", err),
				}
			}
			logger.Sugar().Infof("proxy->plugin TransactionIDInsite: %v ok", info.GetTransactionIDInsite())
		case info := <-p.mpoolPush:
			if err := p.pluginServer.Send(&signproxy.ProxyPluginRequest{
				CoinType:            info.GetCoinType(),
				TransactionType:     info.GetTransactionType(),
				TransactionIDInsite: info.GetTransactionIDInsite(),
				Message:             info.GetMessage(),
				Signature:           info.GetSignature(),
			}); err != nil {
				done <- struct{}{}
				p.ack <- &trading.ACKRequest{
					IsOkay:              false,
					TransactionType:     info.GetTransactionType(),
					TransactionIdInsite: info.GetTransactionIDInsite(),
					Address:             info.GetAddress(),
					ErrorMessage:        fmt.Sprintf("get wallet balance error: %v", err),
				}
				logger.Sugar().Errorf(
					"proxy->plugin TransactionIDInsite: %v TransactionType %v, CoinType: %v Message: %v error: %v",
					info.GetTransactionIDInsite(),
					info.GetTransactionType(),
					info.GetCoinType(),
					info.GetMessage(),
					err,
				)
			}
			logger.Sugar().Infof("proxy->plugin TransactionIDInsite: %v ok", info.GetTransactionIDInsite())
		}
	}
}

func (p *mPlugin) pluginStreamRecv() {
	for {
		psResponse, err := p.pluginServer.Recv()
		if err != nil {
			logger.Sugar().Errorf(
				"proxy->plugin error: %v",
				err,
			)
			continue
		}

		switch psResponse.GetTransactionType() {
		case signproxy.TransactionType_RegisterCoin:
			lmPlugin.append(psResponse.GetCoinType(), *p)
			p.ack <- &trading.ACKRequest{
				IsOkay:              true,
				TransactionType:     psResponse.GetTransactionType(),
				TransactionIdInsite: psResponse.GetTransactionIDInsite(),
				CoinTypeId:          int32(psResponse.GetCoinType()),
			}
			logger.Sugar().Infof("plugin register new coin: %v ok", psResponse.GetCoinType())
			continue
		case signproxy.TransactionType_Balance:
			p.ack <- &trading.ACKRequest{
				IsOkay:              true,
				TransactionType:     psResponse.GetTransactionType(),
				TransactionIdInsite: psResponse.GetTransactionIDInsite(),
				Balance:             float64(psResponse.GetBalance()),
				CoinTypeId:          int32(psResponse.GetCoinType()),
			}
			logger.Sugar().Infof("get TransactionIDInsite: %v wallet balance ok", psResponse.GetTransactionIDInsite())
			continue
		case signproxy.TransactionType_PreSign:
			signProxy, err := getProxySign()
			if err != nil {
				logger.Sugar().Error("proxy->sign no invalid connection")
				p.ack <- &trading.ACKRequest{
					IsOkay:              false,
					TransactionType:     psResponse.GetTransactionType(),
					TransactionIdInsite: psResponse.GetTransactionIDInsite(),
					CoinTypeId:          int32(psResponse.GetCoinType()),
					ErrorMessage:        "proxy->sign no invalid connection",
				}
				done <- struct{}{}
				continue
			}
			signProxy.sign <- &signproxy.ProxySignRequest{
				TransactionType:     psResponse.GetTransactionType(),
				CoinType:            psResponse.GetCoinType(),
				TransactionIDInsite: psResponse.GetTransactionIDInsite(),
				Message: &sphinxplugin.UnsignedMessage{
					To:         psResponse.GetMessage().GetTo(),
					From:       psResponse.GetMessage().GetFrom(),
					Nonce:      psResponse.GetNonce(),
					Value:      psResponse.GetMessage().GetValue(),
					GasLimit:   psResponse.GetMessage().GetGasLimit(),
					GasFeeCap:  psResponse.GetMessage().GetGasFeeCap(),
					GasPremium: psResponse.GetMessage().GetGasPremium(),
					Method:     uint64(builtin.MethodSend),
				},
			}
		case signproxy.TransactionType_Broadcast:
			done <- struct{}{}
			p.ack <- &trading.ACKRequest{
				IsOkay:              true,
				TransactionType:     psResponse.GetTransactionType(),
				TransactionIdInsite: psResponse.GetTransactionIDInsite(),
				CoinTypeId:          int32(psResponse.GetCoinType()),
				TransactionIdChain:  psResponse.GetCID(),
			}
			logger.Sugar().Infof("Broadcast TransactionIDInsite: %v message ok", psResponse.GetTransactionIDInsite())
		}
	}
}

// handle complex trans
// block one to one
func watch() {
	for cproxy := range combineProxyChannl {
		logger.Sugar().Infof("handle new trans info: %v", cproxy.data)
		<-done
		cproxy.mPlugin.preSign <- &signproxy.ProxyPluginRequest{
			CoinType:            cproxy.data.CoinType,
			TransactionType:     cproxy.data.TransactionType,
			TransactionIDInsite: cproxy.data.TransactionIDInsite,
			Address:             cproxy.data.AddressFrom,
			Message: &sphinxplugin.UnsignedMessage{
				To:         cproxy.data.AddressTo,
				From:       cproxy.data.AddressFrom,
				Value:      uint64(unit.FIL2AttoFIL(cproxy.data.AmountFloat64)),
				GasLimit:   655063,
				GasFeeCap:  2300,
				GasPremium: 2250,
				Method:     uint64(builtin.MethodSend),
			},
		}
	}
}

// ConsumerMQ dispatch tran
func ConsumerMQ() error {
	// dispatch combine
	go func() {
		// if start and has trans do it
		done <- struct{}{}
	}()
	go func() {
		watch()
	}()

	consumerInfo, err := msgcli.ConsumeTrans()
	if err != nil {
		logger.Sugar().Errorf("fail to consume %s: %v", msgproducer.GetQueueName(), err)
		return err
	}

	for info := range consumerInfo {
		tinfo := &msgproducer.NotificationTransaction{}
		if err := json.Unmarshal(info.Body, tinfo); err != nil {
			logger.Sugar().Errorf("Unmarshal consumer info error: %v ", err)
			continue
		}

		logger.Sugar().Infof(
			"recv consumer info: %v",
			tinfo,
		)

		ackReq := &trading.ACKRequest{
			TransactionType:     tinfo.TransactionType,
			CoinTypeId:          int32(tinfo.CoinType),
			TransactionIdInsite: tinfo.TransactionIDInsite,
		}

		switch tinfo.TransactionType {
		case signproxy.TransactionType_WalletNew:
			signStream, err := getProxySign()
			if err != nil {
				logger.Sugar().Error("proxy->sign no invalid connection")
				ackReq.IsOkay = false
				goto ackTag
			}
			signStream.walletNew <- &signproxy.ProxySignRequest{
				TransactionType:     tinfo.TransactionType,
				CoinType:            tinfo.CoinType,
				TransactionIDInsite: tinfo.TransactionIDInsite,
			}
		case signproxy.TransactionType_TransactionNew:
			mPlugin, err := getProxyPlugin(tinfo.CoinType)
			if err != nil {
				logger.Sugar().Error("proxy->plugin no invalid connection")
				ackReq.IsOkay = false
				goto ackTag
			}
			combineProxyChannl <- combineProxy{
				mPlugin: mPlugin,
				data:    tinfo,
			}
		case signproxy.TransactionType_Balance:
			pluginStream, err := getProxyPlugin(tinfo.CoinType)
			if err != nil {
				logger.Sugar().Error("proxy->plugin no invalid connection")
				ackReq.IsOkay = false
				goto ackTag
			}
			pluginStream.balance <- &signproxy.ProxyPluginRequest{
				TransactionType:     tinfo.TransactionType,
				CoinType:            tinfo.CoinType,
				TransactionIDInsite: tinfo.TransactionIDInsite,
				Address:             tinfo.AddressFrom,
			}
		default:
			logger.Sugar().Error("consumer info TransactionType: %v invalid", tinfo.TransactionType)
			ackReq.IsOkay = false
			goto ackTag
		}

		logger.Sugar().Infof(
			"deal consumer info TranID: %v CoinType: %v From: %v To: %v ok",
			tinfo.TransactionIDInsite,
			tinfo.CoinType,
			tinfo.AddressFrom,
			tinfo.AddressTo,
		)

		// TODO global ack
		continue

	ackTag:
		{
			if err := ack(ackReq); err != nil {
				logger.Sugar().Infof(
					"deal consumer info TranID: %v CoinType: %v error: %v",
					tinfo.TransactionIDInsite,
					tinfo.CoinType,
					err,
				)
			}
		}
	}

	return nil
}

func ack(ackRequest *trading.ACKRequest) error {
	ackConn, err := grpc2.GetGRPCConn(sconst.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return err
	}
	ackClient := trading.NewTradingClient(ackConn)
	ctx, cancel := context.WithTimeout(context.Background(), constant.GrpcTimeout)
	defer cancel()
	_, err = ackClient.ACK(ctx, ackRequest)
	return err
}
