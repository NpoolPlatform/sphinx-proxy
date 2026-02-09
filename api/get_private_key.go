package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/sphinxproxy"
    "github.com/NpoolPlatform/go-service-framework/pkg/oss"

	"github.com/NpoolPlatform/sphinx-plugin/pkg/coins/getter"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetPrivateKey(ctx context.Context, in *sphinxproxy.GetPrivateKeyRequest) (out *sphinxproxy.GetPrivateKeyResponse, err error) {
	logger.Sugar().Infof("get private key info coinType: %v address: %v", in.GetName(), in.GetAddress())
	if in.GetAddress() == "" {
		logger.Sugar().Errorf("GetPrivateKey Address: %v invalid", in.GetAddress())
		return out, status.Error(codes.InvalidArgument, "Address Invalid")
	}

	if in.GetName() == "" {
		logger.Sugar().Errorf("GetPrivateKey Name: %v empty", in.GetName())
		return out, status.Error(codes.InvalidArgument, "Name empty")
	}

	// old parse method
	coinType := utils.CoinName2Type(in.GetName())
	// this is new information
	pcoinInfo := getter.GetTokenInfo(in.GetName())
	if pcoinInfo != nil && coinType != pcoinInfo.CoinType {
		coinType = pcoinInfo.CoinType
	}

    privateKey, err := oss.GetObject(ctx, in.GetName() + "/" + in.GetAddress(), true)
    if err != nil {
        logger.Sugar().Errorf("GetObject name %v, address %v: %v", in.GetName(), in.GetAddress(), err)
        return nil, status.Error(codes.Internal, "internal server error")
    }

	return &sphinxproxy.GetPrivateKeyResponse {
        Info: string(privateKey),
    }, nil
}
