package app

import (
	"fmt"
	"github.com/streadway/amqp"
)

func RabbitMq() (*amqp.Connection,error){
	user := Configs.RabbitMq.User
	pass := Configs.RabbitMq.Password
	host := Configs.RabbitMq.Host
	port := Configs.RabbitMq.Port
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/",user,pass,host,port)
	// 连接
	conn,err := amqp.Dial(url)
	if err != nil {
		return nil,fmt.Errorf("MQ连接失败 %s\n",err)
	}

	return conn,nil
}

func RabbitMqUrl() string {
	user := Configs.RabbitMq.User
	pass := Configs.RabbitMq.Password
	host := Configs.RabbitMq.Host
	port := Configs.RabbitMq.Port
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/",user,pass,host,port)
	return url
}