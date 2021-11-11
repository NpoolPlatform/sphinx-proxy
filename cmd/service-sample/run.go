package main

import (
	"time"

	"github.com/NpoolPlatform/sphinx-proxy/api"
	db "github.com/NpoolPlatform/sphinx-proxy/pkg/db"
	msgcli "github.com/NpoolPlatform/sphinx-proxy/pkg/message/client"
	msg "github.com/NpoolPlatform/sphinx-proxy/pkg/message/message"
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

		go msgSender()

		return grpc2.RunGRPC(rpcRegister)
	},
}

func rpcRegister(server grpc.ServiceRegistrar) error {
	api.Register(server)
	return nil
}

func msgSender() {
	id := 0
	for {
		logger.Sugar().Infof("send example")
		err := msgsrv.PublishExample(&msg.Example{
			ID:      id,
			Example: "hello world",
		})
		if err != nil {
			logger.Sugar().Errorf("fail to send example: %v", err)
			return
		}
		id++
		time.Sleep(3 * time.Second)
	}
}
