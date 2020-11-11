package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/**
	假设有一个超长的切片，切片的元素类型为int，切片中的元素为乱序排序。限时5秒，使用多个goroutine查找切片中是否存在给定的值，在查找到目标值或者超时后立刻结束所有goroutine的执行。

	比如，切片 [23,32,78,43,76,65,345,762,......915,86]，查找目标值为 345 ，如果切片中存在，则目标值输出"Found it!"并立即取消仍在执行查询任务的goroutine。
	*/
	s := generate()
	ctx, cancel := context.WithCancel(context.Background())

	r := make(chan bool)
	go func() {
		if find(ctx, s[0:len(s)/2], 23) {
			r <- true
		}
	}()

	go func() {
		if find(ctx, s[len(s)/2:], 23) {
			r <- true
		}
	}()

	timer := time.NewTimer(time.Second * 2)

	select {
	case <-timer.C:
		cancel()
		fmt.Println("timeout")
		break
	case <-r:
		cancel()
		fmt.Println("find")
		break
	}
}

func find(ctx context.Context, data []int, target int) bool {
	for _,v := range data {
		if v == target {
			return true
		}
		select {
		case <-ctx.Done():
			return false
		default:
			break
		}
	}
	return false
}

func generate() []int {
	var s []int
	for i := 0; i < 100; i++ {
		s = append(s, rand.Intn(1000))
	}
	s = append(s, 23)
	return s
}
