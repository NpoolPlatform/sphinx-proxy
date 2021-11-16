package api

import (
	"github.com/NpoolPlatform/message/npool/signproxy"
	"google.golang.org/grpc"
)

// https://github.com/grpc/grpc-go/issues/3794
// require_unimplemented_servers=false
type Server struct {
	signproxy.UnimplementedSignProxyServer
}

func Register(server grpc.ServiceRegistrar) {
	signproxy.RegisterSignProxyServer(server, &Server{})
}
