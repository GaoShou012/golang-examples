package main

import (
	"fmt"
	"runtime"
)

/*
	压入数据前:分配的内存 = 77KB, GC的次数 = 0
	压入数据后:分配的内存 = 4250KB, GC的次数 = 1
	删除数据后:分配的内存 = 69KB, GC的次数 = 2
*/

func main(){
	pushInt()
}

func printMemStats(mag string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%v:分配的内存 = %vKB, GC的次数 = %v\n", mag, m.Alloc/1024, m.NumGC)
}

func pushInt(){
	mapint := make(map[int]int)

	printMemStats("压入数据前")
	for i:=0;i<100000;i++{
		mapint[i] = i
	}
	printMemStats("压入数据后")

	for i:=0;i<100000;i++{
		delete(mapint,i)
	}
	runtime.GC()

	printMemStats("删除数据后")
}
