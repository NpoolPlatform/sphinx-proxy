package api

import (
	"github.com/NpoolPlatform/message/npool/sphinxproxy"
	"google.golang.org/grpc"
)

// Server ..
type Server struct {
	sphinxproxy.UnimplementedSphinxProxyServer
}

func Register(server grpc.ServiceRegistrar) {
	sphinxproxy.RegisterSphinxProxyServer(server, &Server{})
}
