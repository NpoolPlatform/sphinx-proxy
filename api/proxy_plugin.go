package api

import "github.com/NpoolPlatform/message/npool/signproxy"

func (s *Server) ProxyPlugin(stream signproxy.SignProxy_ProxyPluginServer) error {
	newPluginStream(stream)
	return nil
}
