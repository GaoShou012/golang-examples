package main

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/utils"
)

func main(){
	utils.Micro.Init("192.168.1.38",2379)

	config,source := utils.Micro.Config,utils.Micro.Source

	// load source
	if err := config.Load(source); err != nil {
		glog.Errorln(err)
		return
	}

	// watch
	watcher,err := config.Watch("micro","config","redis-cluster")
	if err != nil {
		glog.Errorln(err)
		return
	}

	// conf.Map() 可以加载所有的配置
	// confMap := conf.Map()
	// fmt.Printf("conf map=%v\n",confMap)

	for{
		// 进行阻塞
		fmt.Println("watching")
		watcher.Next()

		// 获取改变前的值
		val := config.Get("micro","config","redis-cluster")
		fmt.Printf("This is an old val : %s\n",string(val.Bytes()))

		// 获取改变后的值
		val = config.Get("micro","config","redis-cluster")
		fmt.Printf("This is an new val : %s\n",string(val.Bytes()))
	}
}
