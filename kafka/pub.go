package main

import (
	"examples/config"
	"examples/utils"
	"fmt"
	"github.com/Shopify/sarama"
	"strconv"
	"time"
)

func main() {
	utils.Micro.Init(nil)
	utils.Micro.LoadSource()
	utils.Micro.LoadConfigMust(config.KafkaClusterConfig)

	topic := "im-testing"
	addr := config.KafkaClusterConfig.Addr
	conf := sarama.NewConfig()
	conf.Producer.Retry.Max = 5
	conf.Producer.RequiredAcks = sarama.WaitForAll
	producer, err := sarama.NewAsyncProducer(addr, conf)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 100; i++ {
		val := []byte(fmt.Sprintf("%d", i))
		strTime := strconv.Itoa(int(time.Now().Unix()))
		m := &sarama.ProducerMessage{
			Topic:     topic,
			Key:       sarama.StringEncoder(strTime),
			Value:     sarama.StringEncoder(val),
			Headers:   nil,
			Metadata:  nil,
			Offset:    0,
			Partition: 0,
			Timestamp: time.Time{},
		}
		producer.Input() <- m
	}
	for {
		time.Sleep(time.Second)
	}
}
