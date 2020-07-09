package main

import (
	"fmt"
	"sync"
	"time"
)


func main(){
	var wg sync.WaitGroup

	/*
		执行效果
		begin 2020-07-06 09:05:03.2373343 +0800 CST m=+0.001952001
		finish 2020-07-06 09:05:06.2458134 +0800 CST m=+3.010431101

		总结：
		wg.Add() 添加需要等待的任务数量
		wg.Done() 结束一个任务，等待的任务数量就减1
	*/
	fmt.Println("begin",time.Now())
	wg.Add(1)

	go func() {
		defer wg.Done()
		time.Sleep(time.Second*3)
	}()

	wg.Wait()
	fmt.Println("finish",time.Now())
}
