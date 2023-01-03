package api

import (
	"context"

	coincli "github.com/NpoolPlatform/chain-middleware/pkg/client/coin"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/message/npool"
	coinpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"
	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/message/npool/sphinxproxy"
	"github.com/NpoolPlatform/sphinx-plugin/pkg/coins/getter"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/crud"
	sconst "github.com/NpoolPlatform/sphinx-proxy/pkg/message/const"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/utils"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	ocodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateTransaction(ctx context.Context, in *sphinxproxy.CreateTransactionRequest) (out *sphinxproxy.CreateTransactionResponse, err error) {
	out = &sphinxproxy.CreateTransactionResponse{}

	_, span := otel.Tracer(sconst.ServiceName).Start(ctx, "CreateTransaction")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(ocodes.Error, "call CreateTransaction")
			span.RecordError(err)
		}
	}()

	span.SetAttributes(
		attribute.String("Name", in.GetName()),
		attribute.String("TransactionID", in.GetTransactionID()),
		attribute.Float64("Amount", in.GetAmount()),
		attribute.String("From", in.GetFrom()),
		attribute.String("To", in.GetTo()),
	)

	// args check
	if in.GetName() == "" {
		logger.Sugar().Errorf("CreateTransaction Name: %v empty", in.GetName())
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

	span.AddEvent("call db GetTransactionExist")
	exist, err := crud.GetTransactionExist(ctx, crud.GetTransactionExistParam{TransactionID: in.GetTransactionID()})
	if err != nil {
		logger.Sugar().Errorf("CreateTransaction cal GetTransactionExist error: %v", err)
		return out, status.Error(codes.Internal, "internal server error")
	}

	if exist {
		logger.Sugar().Errorf("CreateTransaction TransactionID: %v already exist", in.GetTransactionID())
		return out, status.Errorf(codes.AlreadyExists, "TransactionID: %v already exist", in.GetTransactionID())
	}

	coinType := utils.CoinName2Type(in.GetName())
	pcoinInfo := getter.GetTokenInfo(in.GetName())
	if pcoinInfo != nil || coinType == sphinxplugin.CoinType_CoinTypeUnKnow {
		coinType = pcoinInfo.CoinType
	}

	// store to db
	span.AddEvent("call db CreateTransaction")
	if err := crud.CreateTransaction(ctx, &crud.CreateTransactionParam{
		CoinType:      coinType,
		TransactionID: in.GetTransactionID(),
		Name:          in.GetName(),
		From:          in.GetFrom(),
		To:            in.GetTo(),
		Value:         in.GetAmount(),
	}); err != nil {
		logger.Sugar().Errorf("CreateTransaction save to db error: %v,TransactionInfo:%v", err, in)
		return out, status.Error(codes.Internal, "internal server error")
	}

	return out, nil
}
