package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Token struct {
	Key string
	Val string
}

type Data struct {
	Username string
	Password string
	Token1 Token
	Token2 Token
}

func main() {
	data := &Data{
		Username: "adsflakjsfjklasjf;jkf;kjljw3q2jrpoqwpjfijqowjae;jkf;kjljw3q2jrpoqwpjfijqowjae;jkf;kjljw3q2jrpoqwpjfijqowjae;jkf;kjljw3q2jrpoqwpjfijqowjae",
		Password: "1239iajsdjlasdjkfajsfjk;sf;jk;jkfasjdfafkj;lafkj;sjfa;ajkf;jkf;kjljw3q2jrpoqwpjfijqowjae",
	}

	taskNum := 30000
	//noParallel(data, taskNum)
	parallel(data, taskNum)
}

func parallel(data *Data, taskNum int) {
	wg := sync.WaitGroup{}

	cache := make(chan *Data, 100000)
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				data := <-cache
				_, err := json.Marshal(data)
				if err != nil {
					panic(err)
				}
				wg.Done()
			}
		}()
	}
	wg.Add(taskNum)
	beginTime := time.Now()
	for i := 0; i < taskNum; i++ {
		cache <- data
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(beginTime))
}

func noParallel(data *Data, taskNum int) {
	wg := sync.WaitGroup{}
	wg.Add(taskNum)
	beginTime := time.Now()
	for i := 0; i < taskNum; i++ {
		go func() {
			defer wg.Done()
			_, err := json.Marshal(data)
			if err != nil {
				panic(err)
			}
		}()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(beginTime))
}
