package api

import (
	"github.com/NpoolPlatform/message/npool/signproxy"
)

func (s *Server) ProxySign(stream signproxy.SignProxy_ProxySignServer) error {
	newSignStream(stream)
	return nil
}
