package main

import (
	"fmt"
	"github.com/app"
)

func main(){
	r := app.Redis()

	sub := r.Subscribe("test")
	for{
		message,err := sub.ReceiveMessage()
		if err != nil {
			panic(err)
		}
		fmt.Println(message.Payload)
	}
}
