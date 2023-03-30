package main

import (
	"fmt"
	s "golang.org/x/sync/singleflight"
	"sync"
	"sync/atomic"
)

func main() {
	g := s.Group{}
	var t uint32 = 0
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			g.Do("test", func() (interface{}, error) {
				atomic.AddUint32(&t, 1)
				fmt.Println("times", t)
				return "value", nil
			})
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("done,times count", t)
}
