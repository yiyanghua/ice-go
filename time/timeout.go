package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	c := make(chan struct{})
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int, close <-chan struct{}) {
			defer func() {
				wg.Done()
			}()
			<-close
			fmt.Println(num)
		}(i, c)
	}

	if waitTimeout(&wg, time.Second*2) {
		close(c)
		fmt.Println("timeout exit")
	}
	time.Sleep(time.Second * 10)
}

func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	// 要求手写代码
	// 要求sync.WaitGroup支持timeout功能
	// 如果timeout到了超时时间返回true
	// 如果WaitGroup自然结束返回false
	ch := make(chan bool)
	go func() {
		wg.Wait()
		ch <- false
	}()

	go time.AfterFunc(timeout, func() {
		ch <- true
	})

	return <-ch
}
