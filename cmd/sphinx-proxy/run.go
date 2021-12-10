package main

import (
	"os"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/sphinx-proxy/api"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	cli "github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

var runCmd = &cli.Command{
	Name:    "run",
	Aliases: []string{"s"},
	Usage:   "Run Sphinx Proxy daemon",
	After: func(c *cli.Context) error {
		if err := grpc2.HShutdown(); err != nil {
			logger.Sugar().Warnf("graceful shutdown http server error: %v", err)
		}
		grpc2.GShutdown()
		return logger.Sync()
	},
	Action: func(c *cli.Context) error {
		go api.Transaction()
		go func() {
			if err := grpc2.RunGRPC(rpcRegister); err != nil {
				logger.Sugar().Warnf("start grpc server error: %v", err)
				os.Exit(1)
			}
		}()
		return grpc2.RunGRPCGateWay(rpcGatewayRegister)
	},
}

// nolint
func rpcRegister(server grpc.ServiceRegistrar) error {
	api.Register(server)
	return nil
}

func rpcGatewayRegister(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return api.RegisterGateway(mux, endpoint, opts)
}
