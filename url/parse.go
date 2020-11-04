package main

import (
	"fmt"
	"github.com/golang/glog"
	"net/url"
	"strings"
)

func main() {
	uri := []byte("/?token=123")
	uri = append(uri, 0x7f)
	var nUri []byte
	for i:=0;i<len(uri);i++{
		b := uri[i]
		if b < ' ' || b == 0x7f {
			continue
		}
		nUri = append(nUri,b)
	}

	//uri := "/?mid=1597053778951-0&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb29tQ29kZSI6Ijg4OWQiLCJ0ZW5hbnRDb2RlIjoieGdjcCIsInVzZXJJZCI6MzIxLCJ1c2VyTmFtZSI6ImFiY2NjIiwidXNlclRodW1iIjoiZGRkIiwidXNlclR5cGUiOiJtYW5hZ2VyIn0.QNSZ3_9sn1VK4ZANAMRoOQMvVnCZLfb2_quF9-dcO7E"
	u, err := url.ParseRequestURI(strings.TrimSpace(string(nUri)))
	if err != nil {
		glog.Errorln(err)
		return
	}

	if str := u.Query().Get("token"); str != "" {
		fmt.Println(str)
	} else {
		err = fmt.Errorf("token is empty")
		return
	}
	if str := u.Query().Get("mid"); str != "" {
		fmt.Println(str)
	} else {
		err = fmt.Errorf("")
	}
}
