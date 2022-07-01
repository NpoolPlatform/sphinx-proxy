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

func (s *Server) UpdateTransaction(ctx context.Context, in *sphinxproxy.UpdateTransactionRequest) (out *sphinxproxy.UpdateTransactionResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, sconst.GrpcTimeout)
	defer cancel()

	if in.GetTransactionID() == "" {
		logger.Sugar().Info("UpdateTransaction TransactionID empty")
		return &sphinxproxy.UpdateTransactionResponse{},
			status.Error(codes.InvalidArgument, "TransactionID empty")
	}

	if in.GetTransactionState() == sphinxproxy.TransactionState_TransactionStateUnKnow ||
		in.GetNextTransactionState() == sphinxproxy.TransactionState_TransactionStateUnKnow {
		logger.Sugar().Info("GetTransactions no wait transaction")
		return &sphinxproxy.UpdateTransactionResponse{},
			status.Error(codes.InvalidArgument, "TransactionState|NextTransactionState empty")
	}

	exist, err := crud.GetTransactionExist(ctx, crud.GetTransactionExistParam{
		TransactionID:    in.GetTransactionID(),
		TransactionState: in.GetTransactionState(),
	})
	if err != nil {
		return &sphinxproxy.UpdateTransactionResponse{},
			status.Error(codes.Internal, "internal server error")
	}

	if !exist {
		return &sphinxproxy.UpdateTransactionResponse{},
			status.Errorf(codes.NotFound, "TransactionID: %v and State: %v not found",
				in.GetTransactionID(),
				in.GetTransactionState(),
			)
	}

	err = crud.UpdateTransaction(ctx, &crud.UpdateTransactionParams{
		TransactionID: in.GetTransactionID(),
		State:         in.GetTransactionState(),
		NextState:     in.GetNextTransactionState(),
		Payload:       in.GetPayload(),
		Cid:           in.GetCID(),
		ExitCode:      in.GetExitCode(),
	})
	if err != nil {
		logger.Sugar().Errorf("GetTransactions call GetTransactions error: %v", err)
		return &sphinxproxy.UpdateTransactionResponse{},
			status.Error(codes.Internal, "internal server error")
	}

	return &sphinxproxy.UpdateTransactionResponse{}, nil
}
