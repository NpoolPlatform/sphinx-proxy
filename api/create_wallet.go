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

type walletDoneInfo struct {
	success bool
	message string
	address string
}

var walletDoneChannel = sync.Map{}

func (s *Server) CreateWallet(ctx context.Context, in *sphinxproxy.CreateWalletRequest) (out *sphinxproxy.CreateWalletResponse, err error) {
	logger.Sugar().Infof("create wallet info: %#v", in)

	if in.GetName() == "" {
		logger.Sugar().Errorf("CreateWallet Name: %v empty", in.GetName())
		return out, status.Error(codes.InvalidArgument, "Name empty")
	}

	coinType, err := utils.ToCoinType(in.GetName())
	if err != nil {
		logger.Sugar().Errorf("CreateWallet Name: %v invalid", in.GetName())
		return out, status.Error(codes.InvalidArgument, "Name Invalid")
	}

	signProxy, err := getProxySign()
	if err != nil {
		logger.Sugar().Errorf("Get ProxySign client not found")
		return out, status.Error(codes.Internal, "internal server error")
	}

	var (
		uid  = uuid.NewString()
		done = make(chan walletDoneInfo)
	)

	walletDoneChannel.Store(uid, done)
	signProxy.walletNew <- &sphinxproxy.ProxySignRequest{
		CoinType:        coinType,
		TransactionType: sphinxproxy.TransactionType_WalletNew,
		TransactionID:   uid,
	}

	// timeout, block wait done
	select {
	case <-time.After(sconst.GrpcTimeout):
		walletDoneChannel.Delete(uid)
		logger.Sugar().Error("create wallet wait response timeout")
		return out, status.Error(codes.Internal, "internal server error")
	case info := <-done:
		walletDoneChannel.Delete(uid)
		if !info.success {
			logger.Sugar().Errorf("wait create wallet done error: %v", info.message)
			return out, status.Error(codes.Internal, "internal server error")
		}
		out = &sphinxproxy.CreateWalletResponse{
			Info: &sphinxproxy.WalletInfo{
				Address: info.address,
			},
		}
	}

	return out, nil
}
