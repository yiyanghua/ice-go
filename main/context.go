package main

import (
	"context"
	"fmt"
	"time"
)

const (
	key = "aa"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	valueCtx := context.WithValue(ctx, key, "add value")

	go watch(valueCtx)
	time.Sleep(10 * time.Second)
	cancel()

	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context) {
	/**
	Done方法返回一个只读的chan，类型为struct{}，
	我们在goroutine中，如果该方法返回的chan可以读取，则意味着parent context已经发起了取消请求，
	我们通过Done方法收到这个信号后，就应该做清理操作，然后退出goroutine，释放资源。
	之后，Err 方法会返回一个错误，告知为什么 Context 被取消。
	*/
	for {
		select {
		case <-ctx.Done():
			//get value
			fmt.Println(ctx.Value(key), "is cancel")

			return
		default:
			//get value
			fmt.Println(ctx.Value(key), "int goroutine")

			time.Sleep(2 * time.Second)
		}
	}
}
