package main

import (
	"fmt"
	"github.com/app"
)

func main(){
	r := app.Redis()
	res,err := r.Publish("test","abc").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
