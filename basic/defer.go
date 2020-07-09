package main

import "fmt"

func test(){
	// 输出结果是:
	// defer2
	// defer1
	// 后入先出，堆栈
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
}

func main(){
	test()
}
