package main

import (
	"github.com/NpoolPlatform/sphinx-proxy/api"
	msgcli "github.com/NpoolPlatform/sphinx-proxy/pkg/message/message"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	cli "github.com/urfave/cli/v2"

	"google.golang.org/grpc"
)

var runCmd = &cli.Command{
	Name:    "run",
	Aliases: []string{"s"},
	Usage:   "Run Sphinx Proxy daemon",
	After: func(c *cli.Context) error {
		grpc2.GShutdown()
		return logger.Sync()
	},
	Action: func(c *cli.Context) error {
		if err := msgcli.Init(); err != nil {
			return err
		}

		return grpc2.RunGRPC(rpcRegister)
	},
}

func rpcRegister(server grpc.ServiceRegistrar) error {
	api.Register(server)
	return nil
}