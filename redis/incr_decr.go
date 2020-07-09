package main

import (
	"fmt"
	"github.com/app"
)

func main() {
	r := app.Redis()

	// res 就是incr后的数值
	// 如果key不存在 默认数值是0，所以会返回1
	res,err := r.Incr("num").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	// rse 就是decr后的数值
	// 如果key不存在，默认数值是0，所以返回-1
	res,err = r.Decr("num1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
