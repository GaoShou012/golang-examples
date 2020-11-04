package main

import (
	"fmt"
	"runtime"
	"time"
)

// 确定map的ID长度越大，占用的内存越大，不能释放

// key type is string
// 1000000 =  88376KB
// 10000000 = 799660KB
// 20000000 = 1599161KB
// 30000000 = 3395109KB

// key type is int
// 10000000 = 192580KB
// 20000000 = 384980KB

var m map[int]bool

func main() {
	m = make(map[int]bool)
	for i := 0; i < 20000000; i++ {
		m[i] = true
	}

	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		printMemStats("当前内存")
		runtime.GC()
	}
}
func printMemStats(mag string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%v:分配的内存 = %vKB, GC的次数 = %v\n", mag, m.Alloc/1024, m.NumGC)
}
