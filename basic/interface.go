package main

import (
	"encoding/json"
	"fmt"
)

type Head struct {
	Id           uint64
	BusinessType string
	BusinessApi  string
}

type Message struct {
	Head
	Body interface{}
}

func main() {
	str := "{'head':{'id':123,'businessType':'bt1','businessApi':'api123'},'body':{'abc':'adsfsdf'}}"

	msg := &Message{}
	err := json.Unmarshal([]byte(str), msg)
	if err != nil {
		panic(err)
	}
	fmt.Println(msg)
}
