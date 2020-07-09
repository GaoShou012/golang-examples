package main

import (
	"fmt"
	"github.com/app"
)

func main(){
	r := app.Redis()

	// res 返回的是，当前list里面的数量
	res,err := r.LPush("li1","abc").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	// str返回的就是内容
	str,err := r.RPop("li1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(str)

	// 查看当前list的长度
	res,err = r.LLen("li1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("llen",res)

	// li2 没有内容，所以err = "redis: nil"
	str,err = r.RPop("li2").Result()
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(str)
	}

	// 根据索引获取元素
	// 如果元素不存在，返回 error = redis: nil
	str,err = r.LIndex("li2",10).Result()
	if err != nil {
		fmt.Println("lindex",err)
	}else{
		fmt.Println(str)
	}

	// 获取列表指定范围内的元素
	// 如果队列不存在，返回一个空的数组
	num,err := r.LRange("li2",0,-1).Result()
	if err != nil {
		fmt.Println("lrange",err)
	}else{
		fmt.Println("lrange",num)
	}
}
