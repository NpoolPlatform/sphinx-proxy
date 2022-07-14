package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/price"
	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/message/npool/sphinxproxy"
	pconst "github.com/NpoolPlatform/sphinx-plugin/pkg/message/const"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/crud"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent"
	sconst "github.com/NpoolPlatform/sphinx-proxy/pkg/message/const"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetTransactions ..
func (s *Server) GetTransactions(ctx context.Context, in *sphinxproxy.GetTransactionsRequest) (out *sphinxproxy.GetTransactionsResponse, err error) {
	pluginInfo := pconst.GetPluginInfo(ctx)
	ctx, cancel := context.WithTimeout(ctx, sconst.GrpcTimeout)
	defer cancel()

	// TODO: debug plugin env info include(position and ip)
	if in.GetTransactionState() == sphinxproxy.TransactionState_TransactionStateUnKnow {
		return &sphinxproxy.GetTransactionsResponse{},
			status.Error(codes.InvalidArgument, "Invalid argument TransactionState must not empty")
	}

	if _, ok := sphinxproxy.TransactionState_name[int32(in.GetTransactionState())]; !ok {
		return &sphinxproxy.GetTransactionsResponse{},
			status.Error(codes.InvalidArgument, "Invalid argument TransactionState not support")
	}

	if in.GetCoinType() != sphinxplugin.CoinType_CoinTypeUnKnow {
		if _, ok := sphinxplugin.CoinType_name[int32(in.GetCoinType())]; !ok {
			return &sphinxproxy.GetTransactionsResponse{},
				status.Error(codes.InvalidArgument, "Invalid argument CoinType not support")
		}
	}

	if in.GetENV() != "" && in.GetENV() != "main" && in.GetENV() != "test" {
		return &sphinxproxy.GetTransactionsResponse{},
			status.Error(codes.InvalidArgument, "Invalid argument ENV only support main|test")
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
			Name:          utils.TruncateCoinTypePrefix(sphinxplugin.CoinType(info.CoinType)),
			Amount:        price.DBPriceToVisualPrice(info.Amount),
			Payload:       info.Payload,
			From:          info.From,
			To:            info.To,
		})
	}

	if len(infos) > 0 {
		logger.Sugar().Infof(
			"%v get tasks,CoinType:%v CoinTransactionState:%v Rows:%v",
			pluginInfo,
			in.GetCoinType(),
			in.GetTransactionState(),
			len(infos),
		)
	}

	return &sphinxproxy.GetTransactionsResponse{
		Infos: infos,
		Total: 0, // total no need
	}, nil
}
