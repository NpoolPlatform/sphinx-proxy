package api

import (
	"github.com/NpoolPlatform/message/npool/signproxy"
	"google.golang.org/grpc"
)

// Server ..
type Server struct {
	signproxy.UnimplementedSignProxyServer
}

func Register(server grpc.ServiceRegistrar) {
	signproxy.RegisterSignProxyServer(server, &Server{})
}
