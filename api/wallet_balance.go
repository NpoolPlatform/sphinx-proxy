package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	pconstant "github.com/NpoolPlatform/sphinx-plugin/pkg/message/const"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/check"
	constant "github.com/NpoolPlatform/sphinx-proxy/pkg/message/const"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) WalletBalance(ctx context.Context, in *sphinxplugin.WalletBalanceRequest) (*sphinxplugin.WalletBalanceResponse, error) {
	// args check
	if in.GetAddress() == "" {
		logger.Sugar().Errorf("[%s] check Address:%v empty",
			constant.FormatServiceName(),
			in.GetAddress(),
		)
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Address:%v empty",
			in.GetAddress(),
		)
	}

	if err := check.CoinType(in.GetCoinType()); err != nil {
		logger.Sugar().Errorf("[%s] check CoinType:%v error: %v",
			constant.FormatServiceName(),
			in.GetCoinType(),
			err,
		)
		return nil, status.Errorf(
			codes.InvalidArgument,
			"CoinType:%v not support",
			in.GetCoinType(),
		)
	}

	conn, err := grpc.GetGRPCConn(config.GetStringValueWithNameSpace("", ""))
	if err != nil {
		logger.Sugar().Errorf("[%s] call GetGRPCConn ServiceName:%s error: %v",
			constant.FormatServiceName(),
			config.GetStringValueWithNameSpace("", pconstant.ServiceName),
			err,
		)
		return nil, status.Error(codes.Internal, "interval server error")
	}

	client := sphinxplugin.NewPluginClient(conn)

	// set deadline
	ctx, cancel := context.WithTimeout(ctx, constant.GrpcTimeout)
	defer cancel()

	resp, err := client.WalletBalance(ctx, &sphinxplugin.WalletBalanceRequest{
		CoinType: in.GetCoinType(),
		Address:  in.GetAddress(),
	})
	if err != nil {
		logger.Sugar().Errorf("[%s] call WalletBalance ServiceName:%s CoinType:%v Address:%v error: %v",
			constant.FormatServiceName(),
			config.GetStringValueWithNameSpace("", pconstant.ServiceName),
			in.GetCoinType(),
			in.GetAddress(),
			err,
		)
		return nil, status.Error(codes.Internal, "interval server error")
	}

	return &sphinxplugin.WalletBalanceResponse{
		Info: &sphinxplugin.WalletBalanceInfo{
			Balance: resp.GetInfo().GetBalance(),
		},
	}, nil
}
