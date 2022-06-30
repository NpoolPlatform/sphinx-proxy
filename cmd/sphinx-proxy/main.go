package main

import (
	"fmt"
	"os"

	"github.com/NpoolPlatform/go-service-framework/pkg/app"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	mysqlconst "github.com/NpoolPlatform/go-service-framework/pkg/mysql/const"
	rabbitmqconst "github.com/NpoolPlatform/go-service-framework/pkg/rabbitmq/const"
	redisconst "github.com/NpoolPlatform/go-service-framework/pkg/redis/const"
	coinconst "github.com/NpoolPlatform/sphinx-coininfo/pkg/message/const"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/servicename"
	cli "github.com/urfave/cli/v2"
)

const serviceName = servicename.ServiceName

func main() {
	commands := cli.Commands{
		runCmd,
	}

	description := fmt.Sprintf("my %v service cli\nFor help on any individual command run <%v COMMAND -h>\n",
		serviceName, serviceName)
	err := app.Init(
		serviceName,
		description,
		"",
		"",
		"./",
		nil,
		commands,
		config.ServiceNameToNamespace(coinconst.ServiceName),
		config.ServiceNameToNamespace(mysqlconst.MysqlServiceName),
		config.ServiceNameToNamespace(redisconst.RedisServiceName),
		config.ServiceNameToNamespace(rabbitmqconst.RabbitMQServiceName),
	)
	if err != nil {
		logger.Sugar().Errorf("fail to create %v: %v", serviceName, err)
		return
	}
	if err := db.Init(); err != nil {
		logger.Sugar().Errorf("fail to init db %v", err)
		return
	}
	err = app.Run(os.Args)
	if err != nil {
		logger.Sugar().Errorf("fail to run %v: %v", serviceName, err)
	}
}
