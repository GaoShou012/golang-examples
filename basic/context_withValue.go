package main

import (
	"context"
	"fmt"
)

/*
	实战场景：
	携带关键信息，为全链路提供线索
	比如接入elk等系统，需要来一个 trace_id
	那withValue就非常适合做这个
*/

func withValue(ctx context.Context) {
	session, ok := ctx.Value("session").(int)
	if !ok {
		fmt.Println("session is wrong")
		return
	}

	if session != 1 {
		fmt.Println("session 不通过")
		return
	}

	traceId := ctx.Value("trace_id").(string)

	// 输出 trace_id: 123456  -session: 1
	fmt.Println("trace_id:", traceId, " -session:", session)
}

func main() {
	ctx := context.WithValue(context.Background(), "trace_id", "123456")
	// 携带session到后面的程序中去
	ctx = context.WithValue(ctx, "session", 1)

	withValue(ctx)
}
