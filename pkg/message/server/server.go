package server

import (
	msg "github.com/NpoolPlatform/sphinx-proxy/pkg/message/message"
)

func Init() error {
	return msg.InitQueues()
}
