package main

import (
	"fmt"
	"sync"
)

// 使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：
// 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728

func main() {
	wait := &sync.WaitGroup{}
	wait.Add(1)

	letter, number := make(chan bool), make(chan bool)

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
				break
			}
		}
	}()

	go func() {
		i := 65
		for {
			if i > 90 {
				wait.Done()
				break
			}

			select {
			case <-letter:
				fmt.Print(string(i))
				i++
				fmt.Print(string(i))
				i++
				number <- true
				break
			}
		}
	}()
	number <- true
	wait.Wait()
}
