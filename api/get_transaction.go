package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/price"
	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/message/npool/sphinxproxy"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/crud"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent"
	sconst "github.com/NpoolPlatform/sphinx-proxy/pkg/message/const"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/utils"
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
	if ent.IsNotFound(err) {
		logger.Sugar().Errorf("GetTransaction TransactionID: %v not found", in.GetTransactionID())
		return &sphinxproxy.GetTransactionResponse{}, status.Errorf(codes.NotFound, "TransactionID: %v not found", in.GetTransactionID())
	}

	if err != nil {
		logger.Sugar().Errorf("GetTransaction call GetTransaction error: %v", err)
		return &sphinxproxy.GetTransactionResponse{}, status.Error(codes.Internal, "internal server error")
	}

	return &sphinxproxy.GetTransactionResponse{
		Info: &sphinxproxy.TransactionInfo{
			TransactionID: transInfo.TransactionID,
			Name:          utils.TruncateCoinTypePrefix(sphinxplugin.CoinType(transInfo.CoinType)),
			Amount:        price.DBPriceToVisualPrice(transInfo.Amount),
			From:          transInfo.From,
			To:            transInfo.To,

			ExitCode:         transInfo.ExitCode,
			CID:              transInfo.Cid,
			TransactionState: sphinxproxy.TransactionState(transInfo.State),

			CreatedAt: transInfo.CreatedAt,
			UpdatedAt: transInfo.UpdatedAt,
		},
	}, nil
}
