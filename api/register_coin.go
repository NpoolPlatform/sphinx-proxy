package api

import (
	"context"
	"fmt"
	"net"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/message/npool/signproxy"
	coinconst "github.com/NpoolPlatform/sphinx-coininfo/pkg/message/const"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/check"
	constant "github.com/NpoolPlatform/sphinx-proxy/pkg/message/const"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) RegisterCoin(ctx context.Context, in *signproxy.RegisterCoinRequest) (*signproxy.RegisterCoinResponse, error) {
	// args check
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

	// get service conn
	svc, err := config.PeekService(config.GetStringValueWithNameSpace("", coinconst.ServiceName), grpc.GRPCTAG)
	if err != nil {
		logger.Sugar().Errorf("[%s] call PeekService ServiceName:%s error: %v",
			constant.FormatServiceName(),
			config.GetStringValueWithNameSpace("", coinconst.ServiceName),
			err,
		)
		return nil, status.Error(codes.Internal, "interval server error")
	}
	conn, err := grpc.GetGRPCConn(net.JoinHostPort(svc.Address, fmt.Sprintf("%d", svc.Port)))
	if err != nil {
		logger.Sugar().Errorf("[%s] call GetGRPCConn error: %v",
			constant.FormatServiceName(),
			err,
		)
		return nil, status.Error(codes.Internal, "interval server error")
	}

	// get service client
	client := coininfo.NewSphinxCoininfoClient(conn)

	// set deadline
	ctx, cancel := context.WithTimeout(ctx, constant.GrpcTimeout)
	defer cancel()

	// new coin exists ?
	coinInfo, err := client.GetCoinInfo(ctx, &coininfo.GetCoinInfoRequest{})
	if err != nil {
		logger.Sugar().Errorf("[%s] call GetCoinInfo error: %v",
			constant.FormatServiceName(),
			err,
		)
		return nil, status.Error(codes.Internal, "interval server error")
	}
	if coinInfo != nil {
		return nil, status.Error(codes.Internal, "")
	}

	if _, err := client.RegisterCoin(ctx, &coininfo.RegisterCoinRequest{}); err != nil {
		return nil, status.Error(codes.Internal, "interval server error")
	}

	return &signproxy.RegisterCoinResponse{}, nil
}
