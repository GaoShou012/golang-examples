package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"github.com/mitchellh/mapstructure"
)

type Head struct {
	Id int
}

type ImMessage struct {
	Head *Head
	Body interface{}
}

type Body struct {
	Name string
}

func main() {
	body := &Body{Name: "123"}
	head := &Head{Id: 333}
	message := &ImMessage{
		Head: head,
		Body: body,
	}
	j,err := json.Marshal(message)
	if err != nil{
		glog.Errorln(err)
		return
	}
	fmt.Println(j)

	msg1 := &ImMessage{}
	if err := json.Unmarshal(j,msg1); err != nil {
		glog.Errorln(err)
		return
	}
	fmt.Println(msg1)
	b := &Body{}
	if err := mapstructure.WeakDecode(msg1.Body,b); err != nil {
		glog.Errorln(err)
		return
	}
	fmt.Println(b)


	str := "abc1"
	j,err = json.Marshal(str)
	if err != nil {
		glog.Errorln(err)
		return
	}
	fmt.Println(j)

	fmt.Println(string(j))
	fmt.Println(str)
}
