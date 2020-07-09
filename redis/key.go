package main

import (
	"fmt"
	"github.com/app"
	"time"
)

func main(){
	r := app.Redis()

	// 没有超时时间
	cmd := r.Set("ke1","val1",0)
	if cmd.Err() != nil {
		fmt.Println(cmd.Err())
	}
	// 带超时时间
	cmd = r.Set("abc","777",time.Second*5)
	if cmd.Err() != nil {
		fmt.Println(cmd.Err())
	}
}