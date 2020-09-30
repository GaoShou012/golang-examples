package main

import "fmt"

func main() {
	basic := make([]byte, 10)
	for i := 0; i < len(basic); i++ {
		basic[i] = '1'
	}
	next := basic[:3]
	for i := 0; i < len(next); i++ {
		next[i] = '2'
	}

	// [50 50 50 49 49 49 49 49 49 49]
	// [50 50 50]
	fmt.Println(basic)
	fmt.Println(next)
}
