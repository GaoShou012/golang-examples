package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

/*
	吃汉堡包比赛
	用时10秒，10秒可以吃多少个

	时间到时后，自动触发 ctx.Done()
*/

func chiHanBao1(ctx context.Context) {
	n := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop")
			return
		default:
			incr := rand.Intn(5)
			n += incr
			fmt.Printf("我吃了 %d 个汉堡包\n", n)
		}
		time.Sleep(time.Second)
	}
}

func main() {
	ctx,cancel := context.WithTimeout(context.Background(),time.Second*10)
	chiHanBao1(ctx)
	defer cancel()
}
