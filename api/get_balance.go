package api

import (
	"context"
	"sync"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/sphinxproxy"
	sconst "github.com/NpoolPlatform/sphinx-proxy/pkg/message/const"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/utils"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type balanceDoneInfo struct {
	success    bool
	message    string
	balance    float64
	balanceStr string
}

var balanceDoneChannel = sync.Map{}

func (s *Server) GetBalance(ctx context.Context, in *sphinxproxy.GetBalanceRequest) (out *sphinxproxy.GetBalanceResponse, err error) {
	if in.GetName() == "" {
		logger.Sugar().Errorf("GetBalance Name: %v empty", in.GetName())
		return out, status.Error(codes.InvalidArgument, "Name empty")
	}

	coinType, err := utils.ToCoinType(in.GetName())
	if err != nil {
		logger.Sugar().Errorf("GetBalance Name: %v invalid", in.GetName())
		return out, status.Error(codes.InvalidArgument, "Name Invalid")
	}

	if in.GetAddress() == "" {
		logger.Sugar().Errorf("GetBalance Address: %v invalid", in.GetAddress())
		return out, status.Error(codes.InvalidArgument, "Address Invalid")
	}

	pluginProxy, err := getProxyPlugin(coinType)
	if err != nil {
		logger.Sugar().Errorf("Get PluginProxy client not found")
		return out, status.Error(codes.Internal, "internal server error")
	}

	var (
		uid  = uuid.NewString()
		done = make(chan balanceDoneInfo)
	)

	balanceDoneChannel.Store(uid, done)
	pluginProxy.balance <- &sphinxproxy.ProxyPluginRequest{
		CoinType:        coinType,
		TransactionType: sphinxproxy.TransactionType_Balance,
		TransactionID:   uid,
		Address:         in.GetAddress(),
	}

	// timeout, block wait done
	select {
	case <-time.After(sconst.GrpcTimeout):
		balanceDoneChannel.Delete(uid)
		logger.Sugar().Errorf("get wallet: %v balance wait response timeout", in.GetAddress())
		return out, status.Error(codes.Internal, "internal server error")
	case info := <-done:
		if !info.success {
			logger.Sugar().Errorf("wait get wallet:%v balance done error: %v", in.GetAddress(), info.message)
			return out, status.Error(codes.Internal, "internal server error")
		}
		out = &sphinxproxy.GetBalanceResponse{
			Info: &sphinxproxy.BalanceInfo{
				Balance:    info.balance,
				BalanceStr: info.balanceStr,
			},
		}
		balanceDoneChannel.Delete(uid)
	}

	return out, nil
}
