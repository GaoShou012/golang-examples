package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/golang/glog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	redisClient := newRedis()
	channel := fmt.Sprintf("testing:channel")
	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			<-ticker.C
			_, err := redisClient.Publish(channel, time.Now().String()).Result()
			if err != nil {
				glog.Errorln(err)
			}
		}
	}()

	go func() {
		for {
			msg, err := redisClient.Subscribe(channel).ReceiveMessage()
			if err != nil {
				glog.Errorln(err)
				continue
			}
			fmt.Println(msg)
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

func newRedis() *redis.ClusterClient {
	addr := []string{"192.168.56.101:9001", "192.168.56.101:9002", "192.168.56.101:9003", "192.168.56.101:9004", "192.168.56.101:9005", "192.168.56.101:9006"}
	password := ""
	minIdleConns := 10
	poolSize := 20

	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:              addr,
		MaxRedirects:       0,
		ReadOnly:           false,
		RouteByLatency:     false,
		RouteRandomly:      false,
		ClusterSlots:       nil,
		OnNewNode:          nil,
		Dialer:             nil,
		OnConnect:          nil,
		Username:           "",
		Password:           password,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		NewClient:          nil,
		PoolSize:           poolSize,
		MinIdleConns:       minIdleConns,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	})
	return client
}
