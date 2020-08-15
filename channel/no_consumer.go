package main

import (
	"fmt"
	"sync"
)

/*
	带缓存的channel，可以不带消费组
*/

var ch chan int

func main(){
	ch = make(chan int,1000)

	ch <- 10

	wg := sync.WaitGroup{}
	go func() {
		wg.Add(1)
		defer wg.Done()
		num := <- ch
		fmt.Println(num)
	}()
	wg.Wait()
	fmt.Println("finished")
}
