package main

import (
	"examples/utils"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/golang/glog"
)


var RedisClient *redis.ClusterClient


func init(){
	RedisClient = utils.NewRedis()
}

func main(){
	stream := "testing:stream:a"

	{
		fmt.Println("读取最新的消息")
		res,err := RedisClient.XRevRangeN(stream,"+","-",1).Result()
		if err != nil {
			glog.Errorln(err)
			return
		}
		fmt.Println(res)
	}
}
