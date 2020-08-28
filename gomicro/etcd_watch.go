package main

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/utils"
)

func main(){
	utils.Micro.Init()
	utils.Micro.LoadSource()

	// watch
	watcher,err := utils.Micro.Config.Watch("micro","config","redis-cluster")
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
		val := utils.Micro.Config.Get("micro","config","redis-cluster")
		fmt.Printf("This is an old val : %s\n",string(val.Bytes()))

		// 获取改变后的值
		val = utils.Micro.Config.Get("micro","config","redis-cluster")
		fmt.Printf("This is an new val : %s\n",string(val.Bytes()))
	}
}
