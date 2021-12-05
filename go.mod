module github.com/NpoolPlatform/sphinx-proxy

go 1.16

require (
	entgo.io/ent v0.9.1
	github.com/NpoolPlatform/go-service-framework v0.0.0-20211121104402-9abc32fd422a
	github.com/NpoolPlatform/message v0.0.0-20211204095120-d48a7c636167
	github.com/NpoolPlatform/sphinx-coininfo v0.0.0-20211118034400-6a6ad84fc182
	github.com/filecoin-project/go-state-types v0.1.1 // indirect
	github.com/filecoin-project/lotus v1.13.1
	github.com/filecoin-project/specs-actors v0.9.14
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.0 // indirect
	github.com/shopspring/decimal v1.3.1
	github.com/ugorji/go v1.2.6 // indirect
	github.com/urfave/cli/v2 v2.3.0
	google.golang.org/grpc v1.42.0
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.41.0

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20181204163529-d75b2dcb6bc8
