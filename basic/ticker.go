package main

import (
	"fmt"
	"time"
)

func main(){
	// 每秒输出时间
	// 例如：2020-07-06 08:56:44.1820076 +0800 CST m=+14.003714101
	ticker := time.NewTicker(time.Second*1)
	for {
		select {
		case t := <- ticker.C:
			fmt.Println(t)
			break
		}
	}
}