package main

import (
	"container/list"
	"fmt"
	"sync"
)

// 并发insert是可以的，没有报错
func main() {
	wg := sync.WaitGroup{}
	li := list.New()
	e := li.PushBack("jing")
	fmt.Println(e)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			li.InsertAfter(i, e)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1000; i < 2000; i++ {
			li.InsertAfter(i, e)
		}
	}()

	wg.Wait()

	for ele := li.Front(); ele != nil; ele = ele.Next() {
		fmt.Println(ele.Value)
	}
}
