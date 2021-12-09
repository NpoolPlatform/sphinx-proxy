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

func (s *Server) ReviewTransaction(ctx context.Context, in *sphinxproxy.ReviewTransactionRequest) (out *sphinxproxy.ReviewTransactionResponse, err error) {
	if in.GetTransactionID() == "" {
		logger.Sugar().Errorf("ReviewTransaction TransactionID empty")
		return &sphinxproxy.ReviewTransactionResponse{}, status.Error(codes.InvalidArgument, "TransactionID empty")
	}

	if in.GetTransactionState() != sphinxproxy.TransactionState_TransactionStateRejected &&
		in.GetTransactionState() != sphinxproxy.TransactionState_TransactionStateWait {
		logger.Sugar().Errorf("ReviewTransaction TransactionState must Reject or Wait")
		return &sphinxproxy.ReviewTransactionResponse{}, status.Error(codes.InvalidArgument, "TransactionState must Reject or Wait")
	}

	ctx, cancel := context.WithTimeout(ctx, sconst.GrpcTimeout)
	defer cancel()

	// TODO paral condition
	exist, err := crud.GetTransactionExist(ctx, crud.GetTransactionExistParam{
		TransactionID:    in.GetTransactionID(),
		TransactionState: in.GetTransactionState(),
	})
	if err != nil {
		logger.Sugar().Errorf("ReviewTransaction cal GetTransactionExist error: %v", err)
		return out, status.Error(codes.Internal, "internal server error")
	}

	if !exist {
		logger.Sugar().Errorf("ReviewTransaction TransactionID: %v TransactionState: %v not found", in.GetTransactionID(), in.GetTransactionState())
		return out, status.Errorf(codes.NotFound, "TransactionID: %v TransactionState: %v not found", in.GetTransactionID(), in.GetTransactionState())
	}

	switch in.GetTransactionState() {
	case sphinxproxy.TransactionState_TransactionStateRejected:
		err = crud.RejectTransaction(ctx, in.GetTransactionID())
	case sphinxproxy.TransactionState_TransactionStateWait:
		err = crud.ConfirmTransaction(ctx, in.GetTransactionID())
	}

	if err != nil {
		logger.Sugar().Errorf("ReviewTransaction error: %v", err)
		return &sphinxproxy.ReviewTransactionResponse{}, status.Error(codes.Internal, "internal server error")
	}

	return &sphinxproxy.ReviewTransactionResponse{}, nil
}
