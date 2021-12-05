package api

import "github.com/NpoolPlatform/message/npool/sphinxproxy"

func (s *Server) ProxySign(stream sphinxproxy.SphinxProxy_ProxySignServer) error {
	newSignStream(stream)
	return nil
}
