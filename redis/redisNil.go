package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
)

func main(){
	// 输出是 redis: nil
	// 可以用作判断redis的nil
	fmt.Println(redis.Nil)
	// 输出 equal
	if "redis: nil" == redis.Nil {
		fmt.Println("equal")
	}
}
