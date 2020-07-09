package main

import (
	"fmt"
	"github.com/app"
	"io/ioutil"
)

func main(){
	r := app.Redis()

	content,err := ioutil.ReadFile("./test.lua")
	if err != nil {
		panic(err)
	}
	sha,err := r.ScriptLoad(string(content)).Result()
	if err != nil {
		panic(err)
	}

	res,err := r.EvalSha(sha,nil,nil).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
