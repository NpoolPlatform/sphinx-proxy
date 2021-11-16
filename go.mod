module github.com/NpoolPlatform/sphinx-proxy

go 1.16

require (
	entgo.io/ent v0.9.1
	github.com/NpoolPlatform/go-service-framework v0.0.0-20211116114223-c078dc6a440d
	github.com/NpoolPlatform/message v0.0.0-20211116100750-cd899d4a9762
	github.com/NpoolPlatform/sphinx-coininfo v0.0.0-20211115022617-243b1ef7313b
	github.com/streadway/amqp v1.0.0
	github.com/urfave/cli/v2 v2.3.0
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/grpc v1.41.0
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.41.0
