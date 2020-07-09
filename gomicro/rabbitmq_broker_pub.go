package main

import (
	"github.com/app"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-plugins/broker/rabbitmq/v2"
)

func main(){
	b := rabbitmq.NewBroker(
		// DNS
		broker.Addrs(app.RabbitMqUrl()),
		// 交换机
		rabbitmq.ExchangeName("orders"),
		// 持久化
		rabbitmq.DurableExchange(),
		)
	b.Init()
	b.Connect()
	b.Publish("orders.new",&broker.Message{
		Header: nil,
		Body:   []byte("abc111"),
	},
		// DeliveryMode = 2 消息持久化
		rabbitmq.DeliveryMode(2),
	)

}
