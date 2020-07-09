package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// n 变量 投递到c channel
// c channel 在main里，使用 for 消费
// 由外部决定goroutine的生命周期
func ctx2(ctx context.Context) <- chan int{
	c := make(chan int)
	// 个数
	n := 0

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case c <- n:
				// 随机增量
				incr := rand.Intn(5)
				n += incr
				// 如果N大于10，就修正为10
				//if n >= 10 {
				//	n = 10
				//}
				fmt.Printf("我吃了 %d 个汉堡包\n", n)
			}
		}
	}()

	return c
}

func main() {
	ctx,cancel := context.WithCancel(context.Background())
	num := ctx2(ctx)
	for n := range num {
		if n >= 10 {
			cancel()
			break
		}
		time.Sleep(time.Second)
		fmt.Println(n)
	}
}
