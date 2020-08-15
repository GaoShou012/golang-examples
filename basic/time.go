package main

import (
	"fmt"
	"time"
)

func main(){
	now := time.Now()
	fmt.Println(now)
	format := time.RFC3339
	before,err := time.Parse(format,"2020-08-12T08:37:46.58280318-04:00")
	fmt.Println(before,err)
	fmt.Println(before.Sub(now).Minutes())
	if now.Sub(before) > 40 {
		fmt.Println("ok more")
	}
}
