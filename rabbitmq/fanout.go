package main

import (
	"fmt"
	"github.com/app"
	"github.com/streadway/amqp"
	"log"
)

func main(){
	conn,err := app.RabbitMq()
	if err != nil {
		log.Println(err)
		return
	}

	// 打开channel
	channel,err := conn.Channel()
	if err != nil {
		log.Println("MQ打开channel失败",err)
		return
	}

	// 创建交换机
	err = channel.ExchangeDeclare("exchange1",amqp.ExchangeFanout,true,false,false,false,nil)
	if err != nil {
		log.Println("创建交换机失败",err)
		return
	}

	// 创建消息队列
	_,err = channel.QueueDeclare("queue1",true,false,false,false,nil)
	if err != nil {
		log.Println("创建消息队列失败",err)
		return
	}

	// 绑定消息队列
	err = channel.QueueBind("queue1","","exchange1",false,nil)
	if err != nil {
		log.Println("消息绑定失败",err)
		return
	}

	_,err = channel.QueueDeclare("queue2",true,false,false,false,nil)
	if err != nil {
		log.Println("创建消息队失败",err)
		return
	}
	err = channel.QueueBind("queue2","k1","exchange1",false,nil)
	if err != nil {
		log.Println("消息绑定失败",err)
		return
	}

	// 发送消息
	//message := amqp.Publishing{
	//	Headers:         nil,
	//	ContentType:     "",
	//	ContentEncoding: "",
	//	DeliveryMode:    0,
	//	Priority:        0,
	//	CorrelationId:   "",
	//	ReplyTo:         "",
	//	Expiration:      "",
	//	MessageId:       "",
	//	Timestamp:       time.Time{},
	//	Type:            "",
	//	UserId:          "",
	//	AppId:           "",
	//	Body:            []byte("112233"),
	//}
	//err = channel.Publish("exchange1","",false,false,message)
	//if err != nil {
	//	log.Println("发送消息失败",err)
	//	return
	//}

	// 拉取消息1
	msg,ok,err := channel.Get("queue1",false)
	if err != nil {
		log.Println("拉取消息失败",err)
		return
	}
	if !ok {
		fmt.Println("n ok ")
		return
	}
	fmt.Println(string(msg.Body))

	// 拉取消息2
	// 如果消息队列不存在，报错
	// 如果消息队列存在，但是没有消息， ok = false
	msg,ok,err = channel.Get("queue2",false)
	if err != nil {
		log.Println("拉取消息失败",err)
		return
	}
	if !ok {
		fmt.Println("n ok")
		return
	}
	fmt.Println(string(msg.Body))
	err = msg.Ack(true)
	if err != nil {
		log.Println("ack",err)
		return
	}
}


