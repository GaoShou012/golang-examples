package main

import (
	"fmt"
	"github.com/app"
	"github.com/go-redis/redis/v7"
)

func main() {
	r := app.Redis()

	// a1 是有序集合的 key
	// z 是有序集合的数据结构，score 就是排序用的数字，member就是唯一的存在
	// ZAdd，当member不存在的时候，添加member，当member存在的时候，就修改score
	index,err := r.ZAdd("a1",&redis.Z{
		Score:  1,
		Member: "member2",
	}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(index)

	// 获取有序集合的成员数量
	// 如果集合不存在，并不会报错，返回 num = 0
	num,err := r.ZCard("a2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(num)

	// 计算在有序集合中指定区间分数的成员数
	num,err = r.ZCount("a1","-100","-1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("ZCount",num)

	// 有序集合中对指定成员的分数加上增量 increment
	// f 就是增量后的结果
	// 成员不存在，并不会报错，默认值是0
	f,err := r.ZIncrBy("a1",1,"member3").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("ZIncrBy",f)

	// 查询范围内的成员列表
	// members是一个数组，每个单元存储的就是member
	// stop = -1 就是查询所有
	members,err := r.ZRange("a1",0,-1).Result()
	if err != nil {
		panic(err)
	}
	for k,v := range members{
		fmt.Println(k,v)
	}


	// 返回成功删除的数量
	// 有序集合不存在的时候，返回的是0，不会报错
	theInt,err := r.ZRem("a1","member4").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(theInt)


}
