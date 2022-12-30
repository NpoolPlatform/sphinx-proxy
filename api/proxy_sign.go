package api

import (
	"github.com/NpoolPlatform/message/npool/sphinxproxy"
	"google.golang.org/grpc/metadata"
)

func (s *Server) ProxySign(stream sphinxproxy.SphinxProxy_ProxySignServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	name := ""
	if ok {
		name = md.Get("name")[0]
	}
	newSignStream(name, stream)
	return nil
}
