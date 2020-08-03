package main

import (
	"container/list"
	"fmt"
	"runtime"
)

/*
压入数据前:分配的内存 = 78KB, GC的次数 = 0
压入数据后:分配的内存 = 5537KB, GC的次数 = 1
删除数据后:分配的内存 = 69KB, GC的次数 = 2
*/

func main(){
	var li list.List
	printMemStats("压入数据前")
	for i:=0;i<100000;i++{
		li.PushBack(i)
	}
	printMemStats("压入数据后")

	for {
		e := li.Front()
		if e == nil {
			break
		}
		li.Remove(e)
	}

	runtime.GC()

	printMemStats("删除数据后")
}

func printMemStats(mag string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%v:分配的内存 = %vKB, GC的次数 = %v\n", mag, m.Alloc/1024, m.NumGC)
}