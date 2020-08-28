package main

import (
	"examples/config"
	"examples/utils"
	"fmt"
	sarama_cluster "github.com/bsm/sarama-cluster"
	"github.com/golang/glog"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	utils.Micro.Init(nil)
	utils.Micro.LoadSource()
	utils.Micro.LoadConfigMust(config.KafkaClusterConfig)

	conf := sarama_cluster.NewConfig()
	conf.Consumer.Offsets.AutoCommit.Enable = true
	conf.Consumer.Offsets.AutoCommit.Interval = time.Second
	conf.Consumer.Offsets.CommitInterval = time.Second
	conf.Consumer.Return.Errors = true
	conf.Group.Return.Notifications = true

	addr, topics := config.KafkaClusterConfig.Addr, []string{"im-testing"}
	consumer, err := sarama_cluster.NewConsumer(addr, "1", topics, conf)
	if err != nil {
		panic(err)
	}

	// trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// consume errors
	go func() {
		for err := range consumer.Errors() {
			log.Printf("Error: %s\n", err.Error())
		}
	}()

	// consume notifications
	go func() {
		for ntf := range consumer.Notifications() {
			log.Printf("Rebalanced: %+v\n", ntf)
		}
	}()

	// consume messages, watch signals
	go func() {
		for {
			select {
			case message, ok := <-consumer.Messages():
				if ok {
					fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\t%s\n", message.Topic, message.Partition, message.Offset, message.Key, message.Value)
					consumer.MarkOffset(message, "") // mark message as processed
				}
			case <-signals:
				return
			}
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		switch s := <-c; s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			glog.Infof("got signal %s; stop server", s)
		case syscall.SIGHUP:
			glog.Infof("got signal %s; go to deamon", s)
			continue
		}
		break
	}
}
