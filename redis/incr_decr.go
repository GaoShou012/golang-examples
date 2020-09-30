package main

import (
	"examples/config"
	"examples/utils"
	"fmt"
	"github.com/go-redis/redis/v7"
	"runtime"
)

var RedisClusterClient *redis.ClusterClient

func main() {
	utils.Micro.Init(nil)
	utils.Micro.LoadSource()
	utils.Micro.LoadConfigMust(config.RedisClusterConfig)

	addr := config.RedisClusterConfig.Addr
	password := config.RedisClusterConfig.Password
	RedisClusterClient = redis.NewClusterClient(&redis.ClusterOptions{
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
		PoolSize:           runtime.NumCPU() * 10,
		MinIdleConns:       runtime.NumCPU(),
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	})

	// res 就是incr后的数值
	// 如果key不存在 默认数值是0，所以会返回1
	res,err := RedisClusterClient.Incr("num").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	// rse 就是decr后的数值
	// 如果key不存在，默认数值是0，所以返回-1
	res,err = RedisClusterClient.Decr("num").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
