package main

import (
	"fmt"
	"github.com/app"
	"time"
)

func main() {
	r := app.Redis()

	pipe := r.TxPipeline()

	// 返回incr1 = 0
	incr1 := pipe.Incr("incr1")
	fmt.Println("incr1",incr1)
	// 返回incr2 = 0
	incr2 := pipe.Incr("incr2")
	fmt.Println("incr2",incr2)

	// 设置incr1 超时
	pipe.Expire("incr1",time.Second*10)

	// 返回 [incr incr1:1 incr incr2:2 expire incr1 10:true]
	res,err := pipe.Exec()
	if err != nil {
		panic(err)
	}
	// 遍历返回的结果
	// 返回的Name() = "incr" "incr" "expire"
	// 返回的String() = "incr incr1:1" "incr incr2:8" "expire incr1 10:true"
	for _,v := range res {
		fmt.Printf("name=%v : string=%v\n",v.Name(),v.String())
	}
}