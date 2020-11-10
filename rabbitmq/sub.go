package main

import (
	"github.com/golang/glog"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@192.168.56.101:5672/")
	if err != nil {
		glog.Errorln(err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		glog.Errorln(err)
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello-queue", // name
		false,         // durable
		false,         // delete when usused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		glog.Errorln(err)
		return
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		glog.Errorln(err)
		return
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			glog.Errorf("Received a message: %s", d.Body)
		}
	}()

	glog.Errorf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
