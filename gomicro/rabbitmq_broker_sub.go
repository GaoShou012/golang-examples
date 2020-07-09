package main

import (
	"fmt"
	"github.com/app"
	_ "github.com/app"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-plugins/broker/rabbitmq/v2"
	"time"
)

/*
	rabbitmq.ExchangeName 以事件角度考虑命名
*/
func main(){
	b := rabbitmq.NewBroker(
		broker.Addrs(app.RabbitMqUrl()),
		rabbitmq.ExchangeName("orders"),
		)
	b.Init()
	b.Connect()
	b.Subscribe("orders.new", func(event broker.Event) error {
		fmt.Println(string(event.Message().Body))
		event.Ack()
		time.Sleep(time.Second * 1)
		return nil
	},
		// 消息队列名称
		broker.Queue("orders.new.sub1"),
		// 不要自动ACK
		broker.DisableAutoAck(),
		// 消息队列持久化
		rabbitmq.DurableQueue(),
	)

	// 启动服务
	service := micro.NewService(
		micro.Name("orders.new.sub1"),
		micro.Registry(app.EtcdRegistry()),
		micro.RegisterTTL(time.Second*10),
		)
	service.Init()
	service.Run()
}
