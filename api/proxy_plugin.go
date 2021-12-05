package api

import "github.com/NpoolPlatform/message/npool/sphinxproxy"

func (s *Server) ProxyPlugin(stream sphinxproxy.SphinxProxy_ProxyPluginServer) error {
	newPluginStream(stream)
	return nil
}
