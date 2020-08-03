package app

import (
	"fmt"
	"runtime"
)

func printMemStats(mag string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%v:分配的内存 = %vKB, GC的次数 = %v\n", mag, m.Alloc/1024, m.NumGC)
}
