package message

import (
	"golang.org/x/xerrors"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	msgcli "github.com/NpoolPlatform/go-service-framework/pkg/rabbitmq/client"
	ssconst "github.com/NpoolPlatform/sphinx-service/pkg/message/const"
	msgproducer "github.com/NpoolPlatform/sphinx-service/pkg/message/message"

	"github.com/streadway/amqp"
)

var vhost = config.ServiceNameToNamespace(ssconst.ServiceName)

type client struct {
	*msgcli.Client
	consumers map[string]<-chan amqp.Delivery
}

var myClients = map[string]*client{}

func Init() error {
	_myClient, err := msgcli.New(vhost)
	if err != nil {
		return err
	}
	queue := msgproducer.GetQueueName()
	err = _myClient.DeclareQueue(queue)
	if err != nil {
		return err
	}

	client := &client{
		Client:    _myClient,
		consumers: map[string]<-chan amqp.Delivery{},
	}

	deliveries, err := _myClient.Consume(queue)
	if err != nil {
		return xerrors.Errorf("fail to construct %s consume: %v", msgproducer.GetQueueName(), err)
	}

	client.consumers[queue] = deliveries
	myClients[vhost] = client

	return nil
}

func ConsumeTrans() (<-chan amqp.Delivery, error) {
	consumers, ok := myClients[vhost].consumers[msgproducer.GetQueueName()]
	if !ok {
		return nil, xerrors.Errorf("consumer is not constructed")
	}

	return consumers, nil
}
