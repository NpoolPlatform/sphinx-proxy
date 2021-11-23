module github.com/NpoolPlatform/sphinx-proxy

go 1.16

require (
	github.com/NpoolPlatform/go-service-framework v0.0.0-20211121104402-9abc32fd422a
	github.com/NpoolPlatform/message v0.0.0-20211123064021-293d02c62a52
	github.com/NpoolPlatform/sphinx-coininfo v0.0.0-20211118034400-6a6ad84fc182
	github.com/NpoolPlatform/sphinx-service v0.0.0-20211121143758-f8fdd9e91a0c
	github.com/filecoin-project/specs-actors v0.9.14
	github.com/streadway/amqp v1.0.0
	github.com/urfave/cli/v2 v2.3.0
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/grpc v1.42.0
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.41.0
