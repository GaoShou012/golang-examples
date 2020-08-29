package main

import "fmt"

type News map[string]*string

func (n News) IsNewMessageId(key string, id string) {
	_, ok := n[key]
	if ok {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
func (n News) Add(key string, val string) {
	n[key] = &val
}
func (n News) Sub(key string) {
	n[key] = nil
}
func (n News) IsSub(key string) bool {
	_, ok := n[key]
	return ok
}

func main() {
	n1 := &News{}
	n2 := n1
	n2.Add("123", "12313")
	n2.IsNewMessageId("123", "12331")
	//n2.Sub("3333")
	fmt.Println(n2.IsSub("3333"))
}
