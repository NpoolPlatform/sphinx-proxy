package api

import (
	"context"

	"github.com/NpoolPlatform/message/npool/sphinxproxy"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// Server ..
type Server struct {
	sphinxproxy.UnimplementedSphinxProxyServer
}

func Register(server grpc.ServiceRegistrar) {
	sphinxproxy.RegisterSphinxProxyServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return sphinxproxy.RegisterSphinxProxyHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
