package main

import (
	"container/list"
	"fmt"
	"time"
)

/*
	多次尝试后报错，并不经常出现
	panic: runtime error: invalid memory address or nil pointer dereference
	[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x460e98]
*/

var li list.List

func main(){
	go func() {
		for{
			li.PushBack(1)
		}
	}()

	go func() {
		for{
			ele := li.Front()
			if ele == nil {
				continue
			}
			li.Remove(ele)
		}
	}()

	go func() {
		for{
			for ele := li.Front(); ele != nil; ele=ele.Next() {
				fmt.Println(ele.Value.(int))
			}
		}
	}()

	fmt.Println("running")
	time.Sleep(time.Second*5)
}
