package main

import (
	"fmt"
	"time"
)

func main(){
	m := make(map[int]int)

	go func() {
		for i:=0;i<100000;i++{
			m[i] = i
		}
	}()
	go func() {
		for{
			for k,v := range m {
				fmt.Println(k,v)
			}
		}
	}()

	time.Sleep(time.Second*5)
}
