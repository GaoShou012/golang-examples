package main

import "fmt"

func main() {
	str := "abc"

	// 不支持abc
	// 支持abc1
	switch str {
	case "abc":
	case "abc1":
		fmt.Println(str)
		break
	}

	// 不支持abc
	// 支持abc1
	switch {
	case str == "abc":
	case str == "abc1":
		fmt.Println(str)
	}
}
