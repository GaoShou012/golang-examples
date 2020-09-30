package main

import (
	"fmt"
	"sync"
	"time"
)

// 测试close是否会阻塞
// 经过测试，channel积压数据，然后close(channel)，不会阻塞
// 如果channel还有数据，会一直等待消费完成后，ok = false

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	ch := make(chan int, 100000)
	go func() {
		defer func() {
			fmt.Println("exit goroutine")
			wg.Done()
		}()
		for {
			i, ok := <-ch
			if !ok {
				return
			}
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}()
	for i := 0; i < 10; i++ {
		ch <- i
	}
	fmt.Println("close channel begin")
	close(ch)
	fmt.Println("close channel end")
	wg.Wait()
}
