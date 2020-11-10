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
	ch, err := conn.Channel()
	if err != nil {
		glog.Errorln(err)
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello-queue", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		glog.Errorln(err)
		return
	}

	body := "{name:arvind, message:hello}"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})
	if err != nil {
		glog.Errorln(err)
		return
	}
}
