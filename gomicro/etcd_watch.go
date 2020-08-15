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
	//confMap := conf.Map()
	//fmt.Printf("conf map=%v\n",confMap)

	for{
		// 进行阻塞
		fmt.Println("watching")
		watcher.Next()

		//if err := conf.Load(src); err != nil {
		//	glog.Errorln(err)
		//	return
		//}

		val := config.Get("micro","config","redis-cluster")
		fmt.Println("the val is %s\n",string(val.Bytes()))

		val = config.Get("micro","config","redis-cluster")
		fmt.Println("the val is %s\n",string(val.Bytes()))
	}

	//for {
	//	watcher,err := conf.Watch("micro","config","database")
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	watcher.Next()
	//
	//	// 如果取消这里的 load
	//	// 导致etcd配置中心，变更value后，conf.Get 收到的值是 变更前的值
	//	err = conf.Load(src)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	val := conf.Get("micro","config","database")
	//	fmt.Printf("the val is %s\n",string(val.Bytes()))
	//
	//	watcher.Stop()
	//}
}
