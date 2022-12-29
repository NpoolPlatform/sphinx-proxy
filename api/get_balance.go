package api

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/message/npool"
	coinpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"
	"github.com/NpoolPlatform/message/npool/sphinxproxy"

	coincli "github.com/NpoolPlatform/chain-middleware/pkg/client/coin"
	ct "github.com/NpoolPlatform/sphinx-plugin/pkg/types"
	sconst "github.com/NpoolPlatform/sphinx-proxy/pkg/message/const"
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
	if in.GetName() == "" {
		logger.Sugar().Errorf("GetBalance Name: %v empty", in.GetName())
		return out, status.Error(codes.InvalidArgument, "Name empty")
	}

	// query coininfo
	_, err = coincli.GetCoinOnly(ctx, &coinpb.Conds{
		Name: &npool.StringVal{
			Op:    cruder.EQ,
			Value: in.GetName(),
		},
	})
	if err != nil {
		logger.Sugar().Errorf("check coin info %v error %v", in.GetName(), err)
		return out, status.Error(codes.Internal, "internal server error")
	}

	coinType := utils.CoinName2Type(in.GetName())

	if in.GetAddress() == "" {
		logger.Sugar().Errorf("GetBalance Address: %v invalid", in.GetAddress())
		return out, status.Error(codes.InvalidArgument, "Address Invalid")
	}

	pluginProxy, err := getProxyPlugin(coinType)
	if err != nil {
		logger.Sugar().Errorf("Get PluginProxy client not found for coinType: %v", coinType)
		return out, status.Error(codes.Internal, "internal server error")
	}

	var (
		uid  = uuid.NewString()
		done = make(chan balanceDoneInfo)
	)

	payload, err := json.Marshal(ct.WalletBalanceRequest{
		Name:    in.GetName(),
		Address: in.GetAddress(),
	})
	if err != nil {
		logger.Sugar().Errorf("Marshal balance request Addr: %v error %v", in.GetAddress(), err)
		return out, status.Error(codes.Internal, "internal server error")
	}

	now := time.Now()
	balanceDoneChannel.Store(uid, done)
	pluginProxy.balance <- &sphinxproxy.ProxyPluginRequest{
		Name:            in.Name,
		CoinType:        coinType,
		TransactionType: sphinxproxy.TransactionType_Balance,
		TransactionID:   uid,
		Payload:         payload,
		Address:         in.GetAddress(),
	}

	// timeout, block wait done
	select {
	case <-time.After(sconst.GrpcTimeout * 6):
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
