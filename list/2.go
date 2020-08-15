package main

import (
	"container/list"
	"sync"
	"time"
)

/*
	测试同时插入，同时遍历
	是否会出错
*/

var li list.List

func main(){
	wg := sync.WaitGroup{}
	go func() {
		for{
			li.PushBack(time.Now().Unix())
			time.Sleep(time.Millisecond)
		}
	}()
	go func() {
		for{
			for ele:=li.Front();ele != nil; ele = ele.Next() {
			}
			time.Sleep(time.Millisecond)
		}
	}()
	wg.Add(1)
	wg.Wait()
}
