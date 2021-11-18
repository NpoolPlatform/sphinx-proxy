module github.com/NpoolPlatform/sphinx-proxy

go 1.16

require (
	entgo.io/ent v0.9.1
	github.com/NpoolPlatform/go-service-framework v0.0.0-20211117074545-bc1340849b08
	github.com/NpoolPlatform/message v0.0.0-20211118083406-0de35a6ec50a
	github.com/NpoolPlatform/sphinx-coininfo v0.0.0-20211118034400-6a6ad84fc182
	github.com/NpoolPlatform/sphinx-plugin v0.0.0-20211115083319-34afa2ad3d5a
	github.com/NpoolPlatform/sphinx-service v0.0.0-20211118082907-0b5459a1a579
	github.com/NpoolPlatform/sphinx-sign v0.0.0-20211110071555-4eee7ac5c56d
	github.com/go-resty/resty/v2 v2.7.0 // indirect
	github.com/shopspring/decimal v1.3.1
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/urfave/cli/v2 v2.3.0
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/grpc v1.42.0
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.41.0
