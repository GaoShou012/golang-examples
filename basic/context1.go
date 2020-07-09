package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func chiHanBao(ctx context.Context) <-chan int {
	c := make(chan int)
	// 个数
	n := 0
	// 时间
	t := 0

	go func() {
		for {
			// time.Sleep(time.seconds)
			select {
			case <-ctx.Done():
				// 收到ctx->Done() 时执行
				fmt.Printf("耗时 %d 秒，吃了 %d 个汉堡包\n", t, n)
				return
			case c <- n:
				// 随机增量
				incr := rand.Intn(5)
				n += incr
				// 如果N大于10，就修正为10
				//if n >= 10 {
				//	n = 10
				//}
				t++
				fmt.Printf("我吃了 %d 个汉堡包\n", n)
			}
		}
	}()

	return c
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	eatNum := chiHanBao(ctx)
	for n := range eatNum {
		if n >= 100 {
			cancel()
			break
		}
	}

	fmt.Println("正在统计结果...")
	time.Sleep(time.Second * 1)
}
