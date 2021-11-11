package api

import (
	"github.com/NpoolPlatform/sphinx-proxy/message/npool/signproxy"
	"github.com/NpoolPlatform/sphinx-proxy/message/npool/version"
	"google.golang.org/grpc"
)

// https://github.com/grpc/grpc-go/issues/3794
// require_unimplemented_servers=false
type Server struct {
	version.UnimplementedVersionServiceServer
	signproxy.UnimplementedSignProxyServer
}

func Register(server grpc.ServiceRegistrar) {
	version.RegisterVersionServiceServer(server, &Server{})
	signproxy.RegisterSignProxyServer(server, &Server{})
}
