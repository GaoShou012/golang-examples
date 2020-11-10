package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
)

/*
	JSON 返回空数组
*/
func main() {
	var data interface{}

	data = make([]interface{},0)
	j,err := json.Marshal(data)
	if err != nil {
		glog.Errorln(err)
		return
	}
	fmt.Println(string(j))
}
