package main

import (
	"fmt"
	"github.com/app"
	"github.com/go-redis/redis/v7"
	"sync"
	"time"
)

func watch2(r *redis.Client) {
	var wg sync.WaitGroup
	wg.Add(2)

	r.Set("kk2", "11", 0)

	go func() {
		defer wg.Done()
		defer fmt.Println("i am finish")

		err := r.Watch(func(tx *redis.Tx) error {
			time.Sleep(time.Second * 3)
			// 获取当前kk2的值
			n, err := tx.Get("kk2").Int()
			if err != nil && err != redis.Nil {
				return err
			}

			// 实际操作
			n++
			fmt.Printf("i am n:%d\n",n)

			// 仅监视key保持不变的情况下执行
			_, err = tx.Pipelined(func(pipeliner redis.Pipeliner) error {
				pipeliner.Set("kk2", n, 0)
				return nil
			})
			return err
		}, "kk2")
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 1)

		err := r.Watch(func(tx *redis.Tx) error {
			// 获取当前kk2的值
			n, err := tx.Get("kk2").Int()
			if err != nil && err != redis.Nil {
				return err
			}

			// 实际操作
			n++
			fmt.Printf("i am n:%d\n",n)

			// 仅监视key保持不变的情况下执行
			_, err = tx.Pipelined(func(pipeliner redis.Pipeliner) error {
				pipeliner.Set("kk2", n, 0)
				return nil
			})
			return err
		}, "kk2")
		if err != nil {
			panic(err)
		}
	}()

	wg.Wait()
}

func watch3(r *redis.Client){
	var wg sync.WaitGroup
	wg.Add(100)
	r.Set("kk2", "0", 0)

	for i:=0;i<100;i++{
		go func() {
			defer wg.Done()

			err := r.Watch(func(tx *redis.Tx) error {
				// 获取当前kk2的值
				n, err := tx.Get("kk2").Int()
				if err != nil && err != redis.Nil {
					return err
				}

				// 实际操作
				n++
				fmt.Printf("i am n:%d\n",n)

				// 仅监视key保持不变的情况下执行
				_, err = tx.Pipelined(func(pipeliner redis.Pipeliner) error {
					pipeliner.Set("kk2", n, 0)
					return nil
				})
				return err
			}, "kk2")
			if err != nil {
				fmt.Println(err)
			}
		}()
	}

	wg.Wait()
}

func watch4(r *redis.Client){
	var wg sync.WaitGroup
	wg.Add(2)
	r.Set("kk2", "0", 0)
	time.Sleep(time.Second*1)

	// 这里会优先watch，但是在写入前睡眠2秒
	// 确保下一个goroutine优先写入，而另到次goroutine 失败
	go func() {
		defer wg.Done()
		err := r.Watch(func(tx *redis.Tx) error {
			// 获取当前kk2的值
			n, err := tx.Get("kk2").Int()
			if err != nil && err != redis.Nil {
				return err
			}

			// 实际操作
			n++
			fmt.Printf("i am n:%d\n",n)

			time.Sleep(time.Second*2)
			// 仅监视key保持不变的情况下执行
			_, err = tx.Pipelined(func(pipeliner redis.Pipeliner) error {
				pipeliner.Set("kk2", n, 0)
				return nil
			})
			return err
		},"kk2")
		if err != nil {
			panic(err)
		}
	}()

	// 这里睡眠1秒，确保上一个goroutine会优先watch
	go func() {
		defer wg.Done()
		time.Sleep(time.Second*1)
		err := r.Watch(func(tx *redis.Tx) error {
			// 获取当前kk2的值
			n, err := tx.Get("kk2").Int()
			if err != nil && err != redis.Nil {
				return err
			}

			// 实际操作
			n++
			fmt.Printf("i am n:%d\n",n)

			// 仅监视key保持不变的情况下执行
			_, err = tx.Pipelined(func(pipeliner redis.Pipeliner) error {
				pipeliner.Set("kk2", n, 0)
				return nil
			})
			return err
		},"kk2")
		if err != nil {
			panic(err)
		}
	}()

	wg.Wait()
}

func main() {
	r := app.Redis()

	/*
		经过实验，watch2 里面的例程试不正确的
		watch3是正确的
		watch4是正确的
	*/

	//watch2(r)
	watch4(r)
}
