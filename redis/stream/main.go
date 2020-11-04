package main

import (
	"examples/utils"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/golang/glog"
)

var RedisClient *redis.ClusterClient

func init() {
	RedisClient = utils.NewRedis()
}

func main() {
	stream := "testing:stream:a"
	{
		values := make(map[string]interface{})
		values["payload"] = []byte("abc")
		xAddArgs := &redis.XAddArgs{
			Stream:       stream,
			MaxLen:       5,
			MaxLenApprox: 0,
			ID:           "*",
			Values:       values,
		}
		str, err := RedisClient.XAdd(xAddArgs).Result()
		if err != nil {
			glog.Errorln(err)
			return
		}
		fmt.Println("message id:", str)
	}

	// 返回[]，暂时不知道怎么使用
	{
		res, err := RedisClient.XInfoGroups(stream).Result()
		if err != nil {
			glog.Errorln(err)
			return
		}
		fmt.Println(res)
	}

	// 正向读取
	// count = 1 	读取数据量1
	// block = -1 	非阻塞
	{
		xReadArgs := &redis.XReadArgs{
			Streams: []string{stream, "0"},
			Count:   1,
			Block:   -1,
		}
		res, err := RedisClient.XRead(xReadArgs).Result()
		if err != nil {
			glog.Errorln(err)
			return
		}
		fmt.Println(res)
	}

	{
		// ID，从小到大 排序出来
		// 不限制数量，把所有的数据打印出来
		fmt.Println("XRange")
		res, err := RedisClient.XRange(stream, "-", "+").Result()
		if err != nil {
			glog.Errorln(err)
			return
		}
		fmt.Println(res)
	}

	{
		// ID,从小到大 排序出来
		// count 限制消息数量
		fmt.Println("XRangeN")
		res,err := RedisClient.XRangeN(stream,"-","+",2).Result()
		if err != nil {
			glog.Errorln(err)
			return
		}
		fmt.Println(res)
	}

	{
		// ID,从大到小，排序出来
		// count 限制消息数量
		fmt.Println("XRevRangeN +-")
		res,err := RedisClient.XRevRangeN(stream,"+","-",2).Result()
		if err != nil {
			glog.Errorln(err)
			return
		}
		for _,row := range res {
			fmt.Println(row)
		}
	}
}
