package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var counter int

var cache chan int

func abc(int) {
	counter++
}

func main() {
	cache = make(chan int, 1000000)
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				<-cache
				counter++
			}
		}()
	}
	time.Sleep(time.Millisecond)

	wg := sync.WaitGroup{}
	wg.Add(runtime.NumCPU())
	now := time.Now()
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 1000000/runtime.NumCPU(); i++ {
				abc(i)
			}
		}()
	}

	//for i := 0; i < runtime.NumCPU(); i++ {
	//	go func() {
	//		defer wg.Done()
	//		for i := 0; i < 1000000/runtime.NumCPU(); i++ {
	//			cache <- i
	//		}
	//	}()
	//}

	wg.Wait()
	fmt.Print(time.Now().Sub(now))

}
