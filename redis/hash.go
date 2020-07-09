package main

import (
	"fmt"
	"github.com/app"
)

func main(){
	r := app.Redis()

	// hmset 同时设置多个字段
	// 设置成功 isOK = true
	isOK,err := r.HMSet("hkey1","num1",10,"num2",11).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(isOK)

	// hmget 返回一个数组
	// 索引与读取的字段排序需要对应
	// 不存在的字段，会返回 nil
	interfaceRes,err := r.HMGet("hkey1","num1","num2","num3").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(interfaceRes)
	fmt.Println(interfaceRes[0])
	fmt.Println(interfaceRes[1])
	// 不存在的字段，可以判断是否等于nil
	if interfaceRes[2] == nil {
		fmt.Println("i am nil")
	}

	// hash不存在，自动创建hash结构，字段的默认数值是0，所以第一次运行返回 res = 1
	res,err := r.HIncrBy("hkey1","num1",1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("hincrby",res)


	// hash结构，只有hincr,没有hdecr



	// 检查字段是否存在
	// 如果字段存在 bRes = true
	// 否则 bRes = false
	bRes,err := r.HExists("hkey1","num3").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(bRes)

	// key存在，返回一个map结构 map[string]string
	// 如果key不存在 返回一个 map[]
	all,err := r.HGetAll("hkey3").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(all)

	// 如果key存在，num 等于 字段的数量
	// 如果key不存在，num = 0 ，不会报错
	num,err := r.HLen("hkey1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(num)

	// 如果字段不存在，设置成功，bRes = true
	// 如果字段存在，设置不成功，bRes = false
	// 如果key不存在，自动创建key，再创建字段，并且返回true
	bRes,err = r.HSetNX("hkeyt1","nx1","1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(bRes)
}
