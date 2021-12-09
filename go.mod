module github.com/NpoolPlatform/sphinx-proxy

go 1.16

require (
	entgo.io/ent v0.9.1
	github.com/NpoolPlatform/go-service-framework v0.0.0-20211207121121-adb2402676f0
	github.com/NpoolPlatform/message v0.0.0-20211209105313-0e3b812f0809
	github.com/NpoolPlatform/sphinx-coininfo v0.0.0-20211206035652-888de6e20996
	github.com/filecoin-project/go-state-types v0.1.1
	github.com/filecoin-project/lotus v1.13.1
	github.com/filecoin-project/specs-actors v0.9.14
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/shopspring/decimal v1.3.1
	github.com/ugorji/go v1.2.6 // indirect
	github.com/urfave/cli/v2 v2.3.0
	golang.org/x/tools v0.1.8 // indirect
	google.golang.org/grpc v1.42.0
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.41.0

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20181204163529-d75b2dcb6bc8
