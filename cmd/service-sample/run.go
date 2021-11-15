package main

import (
	"github.com/NpoolPlatform/sphinx-proxy/api"
	db "github.com/NpoolPlatform/sphinx-proxy/pkg/db"
	msgcli "github.com/NpoolPlatform/sphinx-proxy/pkg/message/client"
	msgsrv "github.com/NpoolPlatform/sphinx-proxy/pkg/message/server"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	cli "github.com/urfave/cli/v2"

	"google.golang.org/grpc"
)

var runCmd = &cli.Command{
	Name:    "run",
	Aliases: []string{"s"},
	Usage:   "Run the daemon",
	After: func(c *cli.Context) error {
		grpc2.GShutdown()
		return logger.Sync()
	},
	Action: func(c *cli.Context) error {
		if err := db.Init(); err != nil {
			return err
		}

		if err := msgsrv.Init(); err != nil {
			return err
		}
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
