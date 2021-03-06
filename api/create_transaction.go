package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/sphinxproxy"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/crud"
	sconst "github.com/NpoolPlatform/sphinx-proxy/pkg/message/const"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateTransaction(ctx context.Context, in *sphinxproxy.CreateTransactionRequest) (out *sphinxproxy.CreateTransactionResponse, err error) {
	out = &sphinxproxy.CreateTransactionResponse{}

	// args check
	if in.GetName() == "" {
		logger.Sugar().Errorf("CreateTransaction Name: %v empty", in.GetName())
		return out, status.Error(codes.InvalidArgument, "Name empty")
	}

	coinType, err := utils.ToCoinType(in.GetName())
	if err != nil {
		logger.Sugar().Errorf("CreateTransaction Name: %v invalid", in.GetName())
		return out, status.Error(codes.InvalidArgument, "Name Invalid")
	}

	if in.GetTransactionID() == "" {
		logger.Sugar().Errorf("CreateTransaction TransactionID: %v invalid", in.GetTransactionID())
		return out, status.Error(codes.InvalidArgument, "TransactionID Invalid")
	}

	if in.GetFrom() == "" {
		logger.Sugar().Errorf("CreateTransaction From: %v invalid", in.GetFrom())
		return out, status.Error(codes.InvalidArgument, "From Invalid")
	}

	if in.GetTo() == "" {
		logger.Sugar().Errorf("CreateTransaction To: %v invalid", in.GetTo())
		return out, status.Error(codes.InvalidArgument, "To Invalid")
	}

	if in.GetAmount() <= 0 {
		logger.Sugar().Errorf("CreateTransaction Amount: %v invalid", in.GetAmount())
		return out, status.Error(codes.InvalidArgument, "Amount Invalid")
	}

	ctx, cancel := context.WithTimeout(ctx, sconst.GrpcTimeout)
	defer cancel()

	exist, err := crud.GetTransactionExist(ctx, crud.GetTransactionExistParam{TransactionID: in.GetTransactionID()})
	if err != nil {
		logger.Sugar().Errorf("CreateTransaction cal GetTransactionExist error: %v", err)
		return out, status.Error(codes.Internal, "internal server error")
	}

	if exist {
		logger.Sugar().Errorf("CreateTransaction TransactionID: %v already exist", in.GetTransactionID())
		return out, status.Errorf(codes.AlreadyExists, "TransactionID: %v already exist", in.GetTransactionID())
	}

	// store to db
	if err := crud.CreateTransaction(ctx, crud.CreateTransactionParam{
		CoinType:      coinType,
		TransactionID: in.GetTransactionID(),
		From:          in.GetFrom(),
		To:            in.GetTo(),
		Value:         in.GetAmount(),
	}); err != nil {
		logger.Sugar().Errorf("CreateTransaction save to db error: %v", err)
		return out, status.Error(codes.Internal, "internal server error")
	}

	return out, nil
}
