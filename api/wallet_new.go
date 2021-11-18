package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/sphinxsign"
	constant "github.com/NpoolPlatform/sphinx-proxy/pkg/message/const"
	sconstant "github.com/NpoolPlatform/sphinx-sign/pkg/message/const"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) WalletNew(ctx context.Context, in *sphinxsign.WalletNewRequest) (*sphinxsign.WalletNewResponse, error) {
	conn, err := grpc.GetGRPCConn(config.GetStringValueWithNameSpace("", ""))
	if err != nil {
		logger.Sugar().Errorf("[%s] call GetGRPCConn ServiceName:%s error: %v",
			constant.FormatServiceName(),
			config.GetStringValueWithNameSpace("", sconstant.ServiceName),
			err,
		)
		return nil, status.Error(codes.Internal, "interval server error")
	}

	client := sphinxsign.NewSignClient(conn)

	// set deadline
	ctx, cancel := context.WithTimeout(ctx, constant.GrpcTimeout)
	defer cancel()

	resp, err := client.WalletNew(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("[%s] call WalletNew ServiceName:%s error: %v",
			constant.FormatServiceName(),
			config.GetStringValueWithNameSpace("", sconstant.ServiceName),
			err,
		)
		return nil, status.Error(codes.Internal, "interval server error")
	}

	return resp, nil
}
