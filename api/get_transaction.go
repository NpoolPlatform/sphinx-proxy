package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/sphinxproxy"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/crud"
	sconst "github.com/NpoolPlatform/sphinx-proxy/pkg/message/const"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetTransaction(ctx context.Context, in *sphinxproxy.GetTransactionRequest) (out *sphinxproxy.GetTransactionResponse, err error) {
	if in.GetTransactionID() == "" {
		logger.Sugar().Errorf("GetTransaction TransactionID empty")
		return &sphinxproxy.GetTransactionResponse{}, status.Error(codes.InvalidArgument, "TransactionID empty")
	}

	ctx, cancel := context.WithTimeout(ctx, sconst.GrpcTimeout)
	defer cancel()

	transInfo, err := crud.GetTransaction(ctx, in.GetTransactionID())
	if err != nil {
		logger.Sugar().Errorf("GetTransaction call GetTransaction error: %v", err)
		return &sphinxproxy.GetTransactionResponse{}, status.Error(codes.Internal, "internal server error")
	}

	return &sphinxproxy.GetTransactionResponse{
		Info: &sphinxproxy.TransactionChainInfo{
			ExitCode: transInfo.ExitCode,
			CID:      transInfo.Cid,
			State:    string(transInfo.State),
		},
	}, nil
}
