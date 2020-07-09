package main

import (
	"encoding/json"
	"fmt"
	"github.com/app"
	"github.com/go-redis/redis/v7"
)

func steam1(r *redis.Client) {

	var message struct{
		UserId uint64
		Username string
	}
	message.UserId = 100
	message.Username = "abc3"

	// 压缩消息
	j,err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(j,&m)
	if err != nil {
		panic(err)
	}

	// 添加消息，并创建steam
	cmd := r.XAdd(&redis.XAddArgs{
		Stream:       "steam1",
		MaxLen:       0,
		MaxLenApprox: 0,
		ID:           "",
		Values:       m,
	})
	if cmd.Err() != nil {
		panic(cmd.Err())
	}

	// 读取steam
	// res 是一个数组，可以同时读取多个steam
	// 1593959676193-0 是开始的消息ID
	res,err := r.XReadStreams("steam1","1593959676193-0").Result()
	if err != nil {
		panic(err)
	}
	for _,v := range res {
		fmt.Println(v.Stream)
		for _,message := range v.Messages {
			// 打印是一个json数据，可以Unmarshal到结构体
			fmt.Println(message)
		}
	}
}

func main(){
	r := app.Redis()

	steam1(r)
}
