package main

import (
	"examples/config"
	"examples/utils"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/golang/glog"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var RedisClusterClient *redis.ClusterClient

func frontierPing(frontierId string) {
	key := fmt.Sprintf("im:frontier:heartbeat")
	RedisClusterClient.HSet(key, frontierId, time.Now().Unix())
}
func frontierHeartbeatList() (map[string]string, error) {
	key := fmt.Sprintf("im:frontier:heartbeat")
	return RedisClusterClient.HGetAll(key).Result()
}
func incrUserCount(tenantCode string, roomCode string, frontierId string) (int64, error) {
	key := fmt.Sprintf("im:rooms:users:count:%s:%s:%s", tenantCode, roomCode, frontierId)
	return RedisClusterClient.Incr(key).Result()
}
func getUserCount(tenantCode string, roomCode string) (count int, err error) {
	now := time.Now().Unix()
	m, err := frontierHeartbeatList()
	if err != nil {
		return
	}
	pipe := RedisClusterClient.TxPipeline()
	for key, val := range m {
		timestamp, err := strconv.Atoi(val)
		if err != nil {
			glog.Errorln(err)
			continue
		}
		if now-int64(timestamp) > 30 {
			continue
		}
		pipe.Get(fmt.Sprintf("im:rooms:users:count:%s:%s:%s", tenantCode, roomCode, key))
	}
	res, err := pipe.Exec()
	if err != nil {
		if err == redis.Nil {
			err = nil
		} else {
			return
		}
	}
	fmt.Println(res)
	for _, val := range res {
		if val.Err() == redis.Nil {
			continue
		}

		row := strings.Split(val.String(), " ")
		num, err := strconv.Atoi(row[2])
		if err != nil {
			glog.Errorln(err)
			continue
		}
		count += num
	}

	return
}

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

	frontierPing("123")
	//incrUserCount("123","123","123")
	count, err := getUserCount("1234", "123")
	if err != nil {
		panic(err)
	}
	fmt.Println(count)

	return
	// 查询room分布的frontier列表
	tenantCode := ""
	roomCode := ""
	key := fmt.Sprintf("im:rooms:frontier:%s:%s", tenantCode, roomCode)
	res, err := RedisClusterClient.HGetAll(key).Result()
	if err != nil {
		panic(err)
	}

	var keys []string
	now := time.Now().Unix()
	for key, val := range res {
		timestamp, err := strconv.Atoi(val)
		if err != nil {
			glog.Errorln(err)
			continue
		}
		if now-int64(timestamp) >= 100 {
			continue
		}
		keys = append(keys, key)
	}
	return

	//pipe := RedisClusterClient.TxPipeline()
	////key := fmt.Sprintf("im:rooms:")
	//pipe.Incr("incr1")
	//pipe.Incr("incr2")
	//pipe.Expire("incr1", time.Second*10)
	//pipe.Expire("incr2", time.Second*10)
	//
	//res, err := pipe.Exec()
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, v := range res {
	//	fmt.Println(v.String())
	//}
}
