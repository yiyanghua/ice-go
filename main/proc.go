package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 写出以下逻辑，要求每秒钟调用一次proc并保证程序不退出?
	timer := time.NewTimer(time.Second * 1)
	go func() {
		// 1 在这里需要你写算法
		// 2 要求每秒钟调用一次proc函数
		// 3 要求程序不能退出
		for {
			select {
			case <-timer.C:
				go func() {
					defer func() {
						if err := recover(); err != nil {
							fmt.Println(err)
						}
					}()
					proc()
				}()
				timer.Reset(time.Second * 1)
			}
		}

	}()
	select {}
}

func proc() {
	panic("ok")
}
