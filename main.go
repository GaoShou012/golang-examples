package main

import (
	"fmt"
	"github.com/app"
)

func main() {
	fmt.Println("ok")
	gorm,err := app.Gorm()
	if err != nil {
		panic(err)
	}
	fmt.Println(gorm)
}
