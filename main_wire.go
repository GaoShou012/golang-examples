package main

import "examples/mock"

/*
// 使用wire前
func main() {
	message := NewMessage("hello world")
	greeter := NewGreeter(message)
	event := NewEvent(greeter)
	event.Start()
}*/

// 使用wire后
func main() {
	event := mock.InitializeEvent("hello_world")
	event.Start()
}


