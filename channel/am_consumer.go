package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	测试百万消费channel
	是否会带来卡机的现象

	验证结果
	就是消息内存，完全不消耗cpu
*/

func main() {
	max := 100000
	channels := make([]chan int, max)
	for i := 0; i < max; i++ {
		channels[i] = make(chan int, 1000)
		go func(i int) {
			for {
				num := <-channels[i]
				fmt.Println(num)
			}
		}(i)
	}
	time.Sleep(time.Second * 3)
	for {
		num := rand.Intn(max)
		//fmt.Println(num)
		channels[num] <- num
		time.Sleep(time.Second)
	}
}
