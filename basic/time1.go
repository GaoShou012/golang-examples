package main

import (
	"fmt"
	"github.com/golang/glog"
	"time"
)

func main() {
	format := time.RFC3339
	str := "2020-09-23T22:42:08.501934868-04:00"
	t, err := time.Parse(format, str)
	if err != nil {
		glog.Errorln(err)
		return
	}
	fmt.Println(t)
	fmt.Println(t.Year())
	fmt.Println(t.Month())
}
