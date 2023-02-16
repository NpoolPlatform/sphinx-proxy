package api

import (
	"github.com/NpoolPlatform/message/npool/sphinxproxy"
	"google.golang.org/grpc/metadata"
)

func (s *Server) ProxySign(stream sphinxproxy.SphinxProxy_ProxySignServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	name := ""
	if ok {
		for _, _name := range md.Get("name") {
			name = _name
		}
	}
	newSignStream(name, stream)
	return nil
}
