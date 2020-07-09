package main

import (
	"fmt"
	"github.com/app"
)

func main(){
	conf,err := app.MicroConfig()
	if err != nil {
		panic(err)
	}
	src,err := app.MicroSource()
	if err != nil {
		panic(err)
	}

	// 如果取消这里的 load
	// 导致etcd配置中心，变更value的时候，下面的代码没有任何触发
	err = conf.Load(src)
	if err != nil {
		panic(err)
	}

	// conf.Map() 可以加载所有的配置
	//confMap := conf.Map()
	//fmt.Printf("conf map=%v\n",confMap)

	for {
		watcher,err := conf.Watch("micro","config","database")
		if err != nil {
			panic(err)
		}

		watcher.Next()

		// 如果取消这里的 load
		// 导致etcd配置中心，变更value后，conf.Get 收到的值是 变更前的值
		err = conf.Load(src)
		if err != nil {
			panic(err)
		}

		val := conf.Get("micro","config","database")
		fmt.Printf("the val is %s\n",string(val.Bytes()))

		watcher.Stop()
	}
}
