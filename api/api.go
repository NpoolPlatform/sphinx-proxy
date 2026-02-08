package api

import (
	"context"

	"github.com/NpoolPlatform/message/npool/sphinxproxy"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	_ "github.com/NpoolPlatform/sphinx-plugin/pkg/types"
)

// Server ..
type Server struct {
	sphinxproxy.UnimplementedSphinxProxyServer
}

func Register(server grpc.ServiceRegistrar) {
	sphinxproxy.RegisterSphinxProxyServer(server, &Server{})
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return sphinxproxy.RegisterSphinxProxyHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
