package message

import (
	"golang.org/x/xerrors"

	msgcli "github.com/NpoolPlatform/go-service-framework/pkg/rabbitmq/client"
	constant "github.com/NpoolPlatform/sphinx-proxy/pkg/message/const"
	msgproducer "github.com/NpoolPlatform/sphinx-service/pkg/message/message"

	"github.com/streadway/amqp"
)

type client struct {
	*msgcli.Client
	consumers map[string]<-chan amqp.Delivery
}

var myClients = map[string]*client{}

func Init() error {
	_myClient, err := msgcli.New(constant.ServiceName)
	if err != nil {
		return err
	}
	queue := msgproducer.GetQueueName()
	err = _myClient.DeclareQueue(queue)
	if err != nil {
		return err
	}

	sampleClient := &client{
		Client:    _myClient,
		consumers: map[string]<-chan amqp.Delivery{},
	}

	deliveries, err := _myClient.Consume(queue)
	if err != nil {
		return xerrors.Errorf("fail to construct %s consume: %v", msgproducer.GetQueueName(), err)
	}

	sampleClient.consumers[queue] = deliveries
	myClients[constant.ServiceName] = sampleClient

	return nil
}

func ConsumeTrans() (<-chan amqp.Delivery, error) {
	consumers, ok := myClients[constant.ServiceName].consumers[msgproducer.GetQueueName()]
	if !ok {
		return nil, xerrors.Errorf("consumer is not constructed")
	}

	return consumers, nil
}
