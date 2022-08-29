package main

import (
	"os"
	"os/signal"
	"syscall"

	apimgrcli "github.com/NpoolPlatform/api-manager/pkg/client"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/sphinx-proxy/api"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	cli "github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

var runCmd = &cli.Command{
	Name:    "run",
	Aliases: []string{"r"},
	Usage:   "Run Sphinx Proxy daemon",
	After: func(*cli.Context) error {
		if err := grpc2.HShutdown(); err != nil {
			logger.Sugar().Warnf("graceful shutdown http server error: %v", err)
		}
		grpc2.GShutdown()
		return logger.Sync()
	},
	Action: func(*cli.Context) error {
		podStopSig := make(chan os.Signal, 1)
		exitChan := make(chan struct{})
		signal.Notify(podStopSig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		go func() {
			<-podStopSig
			logger.Sugar().Info("received SIGTERM to clean conn resource")
			close(exitChan)
		}()
		go func() {
			if err := grpc2.RunGRPC(rpcRegister); err != nil {
				logger.Sugar().Warnf("start grpc server error: %v", err)
				os.Exit(1)
			}
		}()
		return grpc2.RunGRPCGateWay(rpcGatewayRegister)
	},
}

func rpcRegister(server grpc.ServiceRegistrar) error {
	api.Register(server)
	return nil
}

func rpcGatewayRegister(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	err := api.RegisterGateway(mux, endpoint, opts)
	if err != nil {
		return err
	}
	apimgrcli.Register(mux)
	return nil
}
