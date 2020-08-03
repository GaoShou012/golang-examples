package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
	压入数据前:分配的内存 = 78KB, GC的次数 = 0
	压入数据后:分配的内存 = 10608KB, GC的次数 = 2
	删除数据后:分配的内存 = 3955KB, GC的次数 = 3
*/

var m sync.Map

func main(){
	printMemStats("压入数据前")
	for i:=0;i<100000;i++{
		m.Store(i,i)
	}
	printMemStats("压入数据后")

	for i:=0;i<100000;i++{
		m.Delete(i)
	}
	runtime.GC()

	printMemStats("删除数据后")
}

func printMemStats(mag string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%v:分配的内存 = %vKB, GC的次数 = %v\n", mag, m.Alloc/1024, m.NumGC)
}
