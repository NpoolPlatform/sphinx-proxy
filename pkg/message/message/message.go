package message

import (
	msgsrv "github.com/NpoolPlatform/go-service-framework/pkg/rabbitmq/server"
	msgproducer "github.com/NpoolPlatform/sphinx-service/pkg/message/message"
)

func InitQueues() error {
	err := msgsrv.DeclareQueue(msgproducer.GetQueueName())
	if err != nil {
		return err
	}
	return nil
}
