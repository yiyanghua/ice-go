package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var p sync.Pool
var t int32

func init() {
	p = sync.Pool{
		New: func() interface{} {
			atomic.AddInt32(&t, 1)
			return new(string)
		},
	}
}

func main() {

	times := 100
	wg := sync.WaitGroup{}
	wg.Add(times)

	for i := 0; i < times; i++ {
		go func() {
			var s *string
			defer func() {
				if s != nil {
					p.Put(s)
				}
				wg.Done()
			}()
			s = p.Get().(*string)
			*s = "xxxx"
		}()
	}

	wg.Wait()
	fmt.Printf("time %v", t)
}
