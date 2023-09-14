package api

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	coinpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"
	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/message/npool/sphinxproxy"

	coincli "github.com/NpoolPlatform/chain-middleware/pkg/client/coin"
	"github.com/NpoolPlatform/sphinx-plugin/pkg/coins/getter"
	ct "github.com/NpoolPlatform/sphinx-plugin/pkg/types"
	constant "github.com/NpoolPlatform/sphinx-proxy/pkg/const"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/utils"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type balanceDoneInfo struct {
	success bool
	message string
	payload []byte
}

var balanceDoneChannel = sync.Map{}

func (s *Server) GetBalance(ctx context.Context, in *sphinxproxy.GetBalanceRequest) (out *sphinxproxy.GetBalanceResponse, err error) {
	logger.Sugar().Infof("get balance info coinType: %v address: %v", in.GetName(), in.GetAddress())
	if in.GetAddress() == "" {
		logger.Sugar().Errorf("GetBalance Address: %v invalid", in.GetAddress())
		return out, status.Error(codes.InvalidArgument, "Address Invalid")
	}

	if in.GetName() == "" {
		logger.Sugar().Errorf("GetBalance Name: %v empty", in.GetName())
		return out, status.Error(codes.InvalidArgument, "Name empty")
	}

	// query coininfo
	coinExist, err := coincli.GetCoinOnly(ctx, &coinpb.Conds{
		Name: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: in.GetName(),
		},
	})
	if err != nil {
		logger.Sugar().Errorf("check coin info %v error %v", in.GetName(), err)
		return out, status.Error(codes.Internal, "internal server error")
	}

	if coinExist == nil {
		logger.Sugar().Errorf("check coin info %v not exist", in.GetName())
		return out, status.Errorf(codes.NotFound, "coin %v not found", in.GetName())
	}
	// old parse method
	coinType := utils.CoinName2Type(in.GetName())
	// this is new information
	pcoinInfo := getter.GetTokenInfo(in.GetName())
	if pcoinInfo != nil && coinType != pcoinInfo.CoinType {
		coinType = pcoinInfo.CoinType
	}

	pluginProxy, err := getProxyPlugin(coinType)
	if err != nil {
		logger.Sugar().Errorf("Get PluginProxy client not found for coinType: %v", coinType)
		return out, status.Error(codes.Internal, "internal server error")
	}

	var (
		uid     = uuid.NewString()
		done    = make(chan balanceDoneInfo)
		puid    = uuid.NewString()
		pdone   = make(chan balanceDoneInfo)
		payload = make([]byte, 0)
	)

	now := time.Now()
	name := ""
	withPreBalance := false

	switch coinType {
	case sphinxplugin.CoinType_CoinTypealeo, sphinxplugin.CoinType_CoinTypetaleo:
		name = "aleo"
		withPreBalance = true
	case sphinxplugin.CoinType_CoinTypeironfish, sphinxplugin.CoinType_CoinTypetironfish:
		name = "ironfish"
		withPreBalance = true
	default:
		withPreBalance = false
	}

	if withPreBalance {
		balanceDoneChannel.Store(puid, pdone)

		signProxy, err := getProxySign(name)
		if err != nil {
			logger.Sugar().Errorf("Get ProxySign client not found")
			return out, status.Error(codes.Internal, "internal server error")
		}

		fromByte, err := json.Marshal(struct {
			From string `json:"from"`
		}{From: in.GetAddress()})
		if err != nil {
			logger.Sugar().Errorf("Marshal pre balance request Addr: %v error %v", in.GetAddress(), err)
			return out, status.Error(codes.Internal, "internal server error")
		}

		signProxy.preBalance <- &sphinxproxy.ProxySignRequest{
			Name:            in.GetName(),
			CoinType:        coinType,
			TransactionType: sphinxproxy.TransactionType_PreBalance,
			TransactionID:   puid,
			Payload:         fromByte,
		}

		select {
		case <-time.After(constant.GrpcTimeout * 6):
			balanceDoneChannel.Delete(puid)
			return out, status.Error(codes.Internal, "internal server error")
		case info := <-pdone:
			balanceDoneChannel.Delete(puid)
			if !info.success {
				logger.Sugar().Errorf("wait get wallet:%v pre balance done error: %v", in.GetAddress(), info.message)
				return out, status.Error(codes.Internal, "internal server error")
			}

			payload = info.payload
		}
	} else {
		payload, err = json.Marshal(ct.WalletBalanceRequest{
			Name:    in.GetName(),
			Address: in.GetAddress(),
		})
		if err != nil {
			logger.Sugar().Errorf("Marshal balance request Addr: %v error %v", in.GetAddress(), err)
			return out, status.Error(codes.Internal, "internal server error")
		}
	}

	balanceDoneChannel.Store(uid, done)
	pluginProxy.pluginReq <- &sphinxproxy.ProxyPluginRequest{
		Name:            in.GetName(),
		CoinType:        coinType,
		TransactionType: sphinxproxy.TransactionType_Balance,
		TransactionID:   uid,
		Payload:         payload,
		Address:         in.GetAddress(),
	}

	// timeout, block wait done
	select {
	case <-time.After(constant.GrpcTimeout * 6):
		balanceDoneChannel.Delete(uid)
		logger.Sugar().Errorf("get transactionID: %v wallet: %v balance wait response timeout", uid, in.GetAddress())
		return out, status.Error(codes.Internal, "internal server error")
	case info := <-done:
		elasp := time.Since(now).Seconds()
		logger.Sugar().Warnf("get transactionID: %v wallet: %v balance wait response use: %vs", uid, in.GetAddress(), elasp)
		balanceDoneChannel.Delete(uid)
		if !info.success {
			logger.Sugar().Errorf("wait get wallet:%v balance done error: %v", in.GetAddress(), info.message)
			return out, status.Error(codes.Internal, "internal server error")
		}
		balance := ct.WalletBalanceResponse{}
		if err := json.Unmarshal(info.payload, &balance); err != nil {
			logger.Sugar().Errorf("Unmarshal balance info Addr: %v error: %v", in.GetAddress(), err)
			return out, status.Error(codes.Internal, "internal server error")
		}
		out = &sphinxproxy.GetBalanceResponse{
			Info: &sphinxproxy.BalanceInfo{
				Balance:    balance.Balance,
				BalanceStr: balance.BalanceStr,
			},
		}
	}

	return out, nil
}
