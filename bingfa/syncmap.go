package main

import (
	"sync"
	"time"
)

func main(){
	var m sync.Map

	go func() {
		for i:=0;i<100000;i++{
			m.Store(i,i)
		}
	}()
	go func() {
		for i:=0;i<100000;i++{
			m.Delete(i)
		}
	}()
	go func() {
		m.Range(func(key, value interface{}) bool {
			return true
		})
	}()

	time.Sleep(time.Second*5)
}
