package main

import (
	"fmt"
	"github.com/app"
	"time"
)

func main() {
	r := app.Redis()

	// 如果key不存在，设置成功
	// 否则设置失败
	// 0 代表不超时，永久锁
	bRes,err := r.SetNX("key2",time.Now().Unix(),0).Result()
	if err != nil {
		panic(err)
	}
	if bRes == true {
		fmt.Println("设置成功")
	}else{
		fmt.Println("设置失败")
	}
}
