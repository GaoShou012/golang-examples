package main

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/utils"
)

func main() {
	utils.Micro.Init("192.168.1.38", 2379)

	config, source := utils.Micro.Config, utils.Micro.Source
	if err := config.Load(source); err != nil {
		glog.Errorln(err)
		return
	}

	val := config.Get("micro", "config", "redis-cluster")
	fmt.Printf("the val is %s\n", string(val.Bytes()))
}
