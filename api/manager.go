package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"sync"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/message/npool/signproxy"
	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/message/npool/trading"
	scconst "github.com/NpoolPlatform/sphinx-coininfo/pkg/message/const"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/check"
	spconst "github.com/NpoolPlatform/sphinx-proxy/pkg/message/const"
	msgcli "github.com/NpoolPlatform/sphinx-proxy/pkg/message/message"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/unit"
	sconst "github.com/NpoolPlatform/sphinx-service/pkg/message/const"
	msgproducer "github.com/NpoolPlatform/sphinx-service/pkg/message/message"
	"github.com/filecoin-project/specs-actors/actors/builtin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type lmPluginType map[sphinxplugin.CoinType][]*mPlugin

var (
	ErrNoSignServiceFound   = errors.New("no sign service found")
	ErrNoPluginServiceFound = errors.New("no plugin service found")
)

var (
	// rand stream client
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

	ackChannel = make(chan *trading.ACKRequest)

	slk      sync.RWMutex
	plk      sync.RWMutex
	lmSign   = make([]*mSign, 0)
	lmPlugin = make(lmPluginType)

	// control only one transaction can deal at the same time
	done              = make(chan struct{})
	cacheProxyChannel = make(chan *msgproducer.NotificationTransaction, 1024)
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
	signServer   signproxy.SignProxy_ProxySignServer
	closeChannel chan struct{}
	closeFlag    bool
	walletNew    chan *signproxy.ProxySignRequest // address
	sign         chan *signproxy.ProxySignRequest
}

func newSignStream(stream signproxy.SignProxy_ProxySignServer) {
	lc := &mSign{
		signServer:   stream,
		closeChannel: make(chan struct{}),
		closeFlag:    false,
		walletNew:    make(chan *signproxy.ProxySignRequest),
		sign:         make(chan *signproxy.ProxySignRequest),
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

// wallet new
func (s *mSign) signStreamSend(wg *sync.WaitGroup) {
	defer wg.Done()
	for !s.closeFlag {
		select {
		case info := <-s.walletNew:
			logger.Sugar().Infof("proxy->sign TransactionIDInsite: %v start", info.GetTransactionIDInsite())
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
				ackChannel <- &trading.ACKRequest{
					IsOkay:              false,
					TransactionType:     info.GetTransactionType(),
					TransactionIdInsite: info.GetTransactionIDInsite(),
					ErrorMessage:        fmt.Sprintf("create new account error: %v", err),
				}
			}
			logger.Sugar().Infof("proxy->sign TransactionIDInsite: %v ok", info.GetTransactionIDInsite())
		case info := <-s.sign:
			logger.Sugar().Infof("proxy->sign TransactionIDInsite: %v start", info.GetTransactionIDInsite())
			if err := s.signServer.Send(&signproxy.ProxySignRequest{
				TransactionType:     info.GetTransactionType(),
				CoinType:            info.GetCoinType(),
				TransactionIDInsite: info.GetTransactionIDInsite(),
				Message:             info.GetMessage(),
			}); err != nil {
				ackChannel <- &trading.ACKRequest{
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
			logger.Sugar().Infof("proxy->sign TransactionIDInsite: %v ok", info.GetTransactionIDInsite())
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
			"proxy->sign recv TransactionIDInsite: %v TransactionType %v, CoinType: %v",
			ssResponse.GetTransactionIDInsite(),
			ssResponse.GetTransactionType(),
			ssResponse.GetCoinType(),
		)

		switch ssResponse.GetTransactionType() {
		case signproxy.TransactionType_WalletNew:
			ackChannel <- &trading.ACKRequest{
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
				ackChannel <- &trading.ACKRequest{
					IsOkay:              false,
					TransactionType:     ssResponse.GetTransactionType(),
					CoinTypeId:          int32(ssResponse.GetCoinType()),
					TransactionIdInsite: ssResponse.GetTransactionIDInsite(),
					Address:             ssResponse.GetInfo().GetAddress(),
				}
				continue
			}
			pluginProxy.mpoolPush <- &signproxy.ProxyPluginRequest{
				CoinType:            ssResponse.GetCoinType(),
				TransactionType:     signproxy.TransactionType_Broadcast,
				TransactionIDInsite: ssResponse.GetTransactionIDInsite(),
				Message:             ssResponse.GetInfo().GetMessage(),
				Signature:           ssResponse.GetInfo().GetSignature(),
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
	pluginServer signproxy.SignProxy_ProxyPluginServer
	coinType     sphinxplugin.CoinType
	closeChannel chan struct{}
	closeFlag    bool
	balance      chan *signproxy.ProxyPluginRequest
	preSign      chan *signproxy.ProxyPluginRequest
	mpoolPush    chan *signproxy.ProxyPluginRequest
	registerCoin chan *signproxy.ProxyPluginResponse
}

func newPluginStream(stream signproxy.SignProxy_ProxyPluginServer) {
	lp := &mPlugin{
		pluginServer: stream,
		closeChannel: make(chan struct{}),
		closeFlag:    false,
		balance:      make(chan *signproxy.ProxyPluginRequest),
		preSign:      make(chan *signproxy.ProxyPluginRequest),
		mpoolPush:    make(chan *signproxy.ProxyPluginRequest),
		registerCoin: make(chan *signproxy.ProxyPluginResponse),
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
		select {
		case info := <-p.balance:
			if err := p.pluginServer.Send(&signproxy.ProxyPluginRequest{
				CoinType:            info.GetCoinType(),
				TransactionType:     info.GetTransactionType(),
				TransactionIDInsite: info.GetTransactionIDInsite(),
				Address:             info.GetAddress(),
			}); err != nil {
				ackChannel <- &trading.ACKRequest{
					IsOkay:              false,
					TransactionType:     info.GetTransactionType(),
					TransactionIdInsite: info.GetTransactionIDInsite(),
					Address:             info.GetAddress(),
					ErrorMessage:        fmt.Sprintf("get wallet balance error: %v", err),
				}
				logger.Sugar().Errorf(
					"proxy->plugin TransactionIDInsite: %v TransactionType: %v CoinType: %v Address: %v error: %v",
					info.GetTransactionIDInsite(),
					info.GetTransactionType(),
					info.GetCoinType(),
					info.GetAddress(),
					err,
				)
				continue
			}
			logger.Sugar().Infof("proxy->plugin TransactionIDInsite: %v ok", info.GetTransactionIDInsite())
		case info := <-p.preSign:
			if err := p.pluginServer.Send(&signproxy.ProxyPluginRequest{
				CoinType:            info.GetCoinType(),
				TransactionType:     info.GetTransactionType(),
				TransactionIDInsite: info.GetTransactionIDInsite(),
				Address:             info.GetAddress(),
				Message:             info.GetMessage(),
			}); err != nil {
				ackChannel <- &trading.ACKRequest{
					IsOkay:              false,
					TransactionType:     info.GetTransactionType(),
					TransactionIdInsite: info.GetTransactionIDInsite(),
					Address:             info.GetAddress(),
					ErrorMessage:        fmt.Sprintf("get wallet nonce error: %v", err),
				}
				logger.Sugar().Errorf(
					"proxy->plugin TransactionIDInsite: %v TransactionType %v, CoinType: %v Address: %v error: %v",
					info.GetTransactionIDInsite(),
					info.GetTransactionType(),
					info.GetCoinType(),
					info.GetAddress(),
					err,
				)
				continue
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
				ackChannel <- &trading.ACKRequest{
					IsOkay:              false,
					TransactionType:     info.GetTransactionType(),
					TransactionIdInsite: info.GetTransactionIDInsite(),
					Address:             info.GetAddress(),
					ErrorMessage:        fmt.Sprintf("mpool push error: %v", err),
				}
				logger.Sugar().Errorf(
					"proxy->plugin TransactionIDInsite: %v TransactionType %v, CoinType: %v Message: %v error: %v",
					info.GetTransactionIDInsite(),
					info.GetTransactionType(),
					info.GetCoinType(),
					info.GetMessage(),
					err,
				)
				continue
			}
			logger.Sugar().Infof("proxy->plugin TransactionIDInsite: %v ok", info.GetTransactionIDInsite())
		}
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
		case signproxy.TransactionType_RegisterCoin:
			lmPlugin.append(psResponse.GetCoinType(), p)
			ackChannel <- &trading.ACKRequest{
				IsOkay:              true,
				TransactionType:     psResponse.GetTransactionType(),
				TransactionIdInsite: psResponse.GetTransactionIDInsite(),
				CoinTypeId:          int32(psResponse.GetCoinType()),
			}
			logger.Sugar().Infof("plugin register new coin: %v ok", psResponse.GetCoinType())
			continue
		case signproxy.TransactionType_Balance:
			ackChannel <- &trading.ACKRequest{
				IsOkay:              true,
				TransactionType:     psResponse.GetTransactionType(),
				TransactionIdInsite: psResponse.GetTransactionIDInsite(),
				Balance:             psResponse.GetBalance(),
				CoinTypeId:          int32(psResponse.GetCoinType()),
			}
			logger.Sugar().Infof("get TransactionIDInsite: %v wallet balance ok", psResponse.GetTransactionIDInsite())
			continue
		case signproxy.TransactionType_PreSign:
			signProxy, err := getProxySign()
			if err != nil {
				logger.Sugar().Error("proxy->sign no invalid connection")
				ackChannel <- &trading.ACKRequest{
					IsOkay:              false,
					TransactionType:     psResponse.GetTransactionType(),
					TransactionIdInsite: psResponse.GetTransactionIDInsite(),
					CoinTypeId:          int32(psResponse.GetCoinType()),
					ErrorMessage:        "proxy->sign no invalid connection",
				}
				continue
			}
			signProxy.sign <- &signproxy.ProxySignRequest{
				TransactionType:     signproxy.TransactionType_Signature,
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
			ackChannel <- &trading.ACKRequest{
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

// ConsumerMQ dispatch tran
func ConsumerMQ() error {
	go func() {
		// if start and has trans do it
		done <- struct{}{}
	}()
	go func() {
		watch()
	}()
	go func() {
		ack()
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

		if err := checkParam(tinfo, ackReq); err != nil {
			logger.Sugar().Errorf("call checkParam error: %v", err)
			ackChannel <- ackReq
			continue
		}

		switch tinfo.TransactionType {
		case signproxy.TransactionType_WalletNew:
			signStream, err := getProxySign()
			if err != nil {
				logger.Sugar().Error("proxy->sign no invalid connection")
				ackReq.IsOkay = false
				ackChannel <- ackReq
				continue
			}
			signStream.walletNew <- &signproxy.ProxySignRequest{
				TransactionType:     tinfo.TransactionType,
				CoinType:            tinfo.CoinType,
				TransactionIDInsite: tinfo.TransactionIDInsite,
			}
		case signproxy.TransactionType_TransactionNew:
			// cache and limit speed
			cacheProxyChannel <- tinfo
		case signproxy.TransactionType_Balance:
			if err := check.CoinType(tinfo.CoinType); err != nil {
				logger.Sugar().Errorf("consumer info CoinType: %v invalid", tinfo.CoinType)
				ackReq.IsOkay = false
				ackReq.ErrorMessage = fmt.Sprintf("CoinType: %v invalid", tinfo.CoinType)
				ackChannel <- ackReq
				continue
			}

			if tinfo.AddressFrom == "" {
				logger.Sugar().Errorf("consumer info AddressFrom empty")
				ackReq.IsOkay = false
				ackReq.ErrorMessage = "AddressFrom is empty"
				ackChannel <- ackReq
				continue
			}

			pluginStream, err := getProxyPlugin(tinfo.CoinType)
			if err != nil {
				logger.Sugar().Error("proxy->plugin no invalid connection")
				ackReq.IsOkay = false
				ackChannel <- ackReq
				continue
			}
			pluginStream.balance <- &signproxy.ProxyPluginRequest{
				TransactionType:     tinfo.TransactionType,
				CoinType:            tinfo.CoinType,
				TransactionIDInsite: tinfo.TransactionIDInsite,
				Address:             tinfo.AddressFrom,
			}
		default:
			logger.Sugar().Errorf("consumer info TransactionType: %v invalid", tinfo.TransactionType)
			ackReq.IsOkay = false
			ackChannel <- ackReq
			continue
		}

		logger.Sugar().Infof(
			"deal consumer info TranID: %v CoinType: %v From: %v To: %v ok",
			tinfo.TransactionIDInsite,
			tinfo.CoinType,
			tinfo.AddressFrom,
			tinfo.AddressTo,
		)
	}
	return nil
}

func checkParam(tinfo *msgproducer.NotificationTransaction, ackReq *trading.ACKRequest) error {
	if err := check.CoinType(tinfo.CoinType); err != nil {
		logger.Sugar().Errorf("consumer info CoinType: %v invalid", tinfo.CoinType)
		ackReq.IsOkay = false
		ackReq.ErrorMessage = "CoinType invalid"
		return err
	}

	if err := check.TransactionType(tinfo.TransactionType); err != nil {
		logger.Sugar().Errorf("consumer info TransactionType: %v invalid", tinfo.TransactionType)
		ackReq.IsOkay = false
		ackReq.ErrorMessage = "TransactionType invalid"
		return err
	}

	return nil
}

// handle complex trans
// block one to one
func watch() {
	for cproxy := range cacheProxyChannel {
		logger.Sugar().Infof("handle new trans info: %v", cproxy)
		<-done
		pluginProxy, err := getProxyPlugin(cproxy.CoinType)
		if err != nil {
			ackChannel <- &trading.ACKRequest{
				IsOkay:              false,
				TransactionType:     cproxy.TransactionType,
				CoinTypeId:          int32(cproxy.CoinType),
				TransactionIdInsite: cproxy.TransactionIDInsite,
			}
			continue
		}
		pluginProxy.preSign <- &signproxy.ProxyPluginRequest{
			CoinType:            cproxy.CoinType,
			TransactionType:     signproxy.TransactionType_PreSign,
			TransactionIDInsite: cproxy.TransactionIDInsite,
			Address:             cproxy.AddressFrom,
			Message: &sphinxplugin.UnsignedMessage{
				To:         cproxy.AddressTo,
				From:       cproxy.AddressFrom,
				Value:      uint64(unit.FIL2AttoFIL(cproxy.AmountFloat64)),
				GasLimit:   655063,
				GasFeeCap:  2300,
				GasPremium: 2250,
				Method:     uint64(builtin.MethodSend),
			},
		}
	}
}

func ack() {
	for {
		ackInfo := <-ackChannel
		switch ackInfo.TransactionType {
		case signproxy.TransactionType_RegisterCoin:
			registerCoin(ackInfo)
		default:
			transaction(ackInfo)
		}
	}
}

func transaction(ackInfo *trading.ACKRequest) {
	logger.Sugar().Infof("ack info: %v", ackInfo)

	// check transaction type
	switch ackInfo.GetTransactionType() {
	case
		signproxy.TransactionType_TransactionNew,
		signproxy.TransactionType_Signature,
		signproxy.TransactionType_PreSign,
		signproxy.TransactionType_Broadcast:
		// TransactionNew PreSign Signature when fail
		// Broadcast succ or fail
		go func() { done <- struct{}{} }()
	}

	ackConn, err := grpc2.GetGRPCConn(sconst.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		logger.Sugar().Errorf("ack call GetGRPCConn error: %v", err)
		return
	}

	ackClient := trading.NewTradingClient(ackConn)
	ctx, cancel := context.WithTimeout(context.Background(), spconst.GrpcTimeout)
	defer cancel()
	_, err = ackClient.ACK(ctx, ackInfo)
	if err != nil {
		logger.Sugar().Errorf("ack error: %v", err)
	}
}

func registerCoin(ackInfo *trading.ACKRequest) {
	logger.Sugar().Infof("register info: %v", ackInfo)
	ackConn, err := grpc2.GetGRPCConn(scconst.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		logger.Sugar().Errorf("ack call GetGRPCConn error: %v", err)
		return
	}
	client := coininfo.NewSphinxCoinInfoClient(ackConn)
	ctx, cancel := context.WithTimeout(context.Background(), spconst.GrpcTimeout)
	defer cancel()

	// define in plugin
	_, err = client.CreateCoinInfo(ctx, &coininfo.CreateCoinInfoRequest{
		Info: &coininfo.CoinInfo{
			Enum: ackInfo.CoinTypeId,
			Name: check.TruncateCoinTypePrefix(sphinxplugin.CoinType(ackInfo.CoinTypeId)),
			Unit: "FIL",
		},
	})

	if err != nil {
		logger.Sugar().Errorf("ack error: %v", err)
	}
}

func checkCode(err error) bool {
	if err == io.EOF ||
		status.Code(err) == codes.Unavailable ||
		status.Code(err) == codes.Canceled {
		return true
	}
	return false
}
