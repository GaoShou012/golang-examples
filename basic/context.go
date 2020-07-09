package main

import (
	"context"
	"fmt"
)

/*
	遵循以下规则，以保持包之间的接口一致，并启用静态分析工具以检查上下文传播。

	不要将 Contexts 放入结构体，相反context应该作为第一个参数传入，命名为ctx。 func DoSomething（ctx context.Context，arg Arg）error { // ... use ctx ... }
	即使函数允许，也不要传入nil的 Context。如果不知道用哪种 Context，可以使用context.TODO()。
	使用context的Value相关方法只应该用于在程序和接口中传递的和请求相关的元数据，不要用它来传递一些可选的参数
	相同的 Context 可以传递给在不同的goroutine；Context 是并发安全的。
*/


/*
	此示例演示使用一个可取消的上下文，以防止 goroutine 泄漏。示例函数结束时，defer 调用 cancel 方法，gen goroutine 将返回而不泄漏。
*/
func c1(){
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}


func main(){
	c1()
}
