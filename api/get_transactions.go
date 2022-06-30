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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetTransactions ..
func (s *Server) GetTransactions(ctx context.Context, in *sphinxproxy.GetTransactionsRequest) (out *sphinxproxy.GetTransactionsResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, sconst.GrpcTimeout)
	defer cancel()

	// TODO: debug plugin env info include(position and ip)
	if in.GetCoinType() == sphinxplugin.CoinType_CoinTypeUnKnow ||
		in.GetTransactionState() == sphinxproxy.TransactionState_TransactionStateUnKnow ||
		(in.GetENV() != "main" && in.GetENV() != "test") {
		return &sphinxproxy.GetTransactionsResponse{}, nil
	}

	transInfos, err := crud.GetTransactions(ctx, crud.GetTransactionsParam{
		CoinType:         in.GetCoinType(),
		TransactionState: in.GetTransactionState(),
	})
	if ent.IsNotFound(err) {
		logger.Sugar().Info("GetTransactions no wait transaction")
		return &sphinxproxy.GetTransactionsResponse{}, nil
	}

	if err != nil {
		logger.Sugar().Errorf("GetTransactions call GetTransactions error: %v", err)
		return &sphinxproxy.GetTransactionsResponse{}, status.Error(codes.Internal, "internal server error")
	}

	infos := make([]*sphinxproxy.TransactionInfo, 0, len(transInfos))
	for _, info := range transInfos {
		infos = append(infos, &sphinxproxy.TransactionInfo{
			TransactionID: info.TransactionID,
			// Name             :info.
			Amount: price.DBPriceToVisualPrice(info.Amount),
			From:   info.From,
			To:     info.To,
		})
	}

	return &sphinxproxy.GetTransactionsResponse{
		Infos: infos,
		Total: 0, // total no need
	}, nil
}
