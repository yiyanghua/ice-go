package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Ban struct {
	visitIps map[string]time.Time
	rmx sync.Mutex
}

func NewBan(ctx context.Context) *Ban {
	o := &Ban{
		visitIps: make(map[string]time.Time),
		rmx: sync.Mutex{},
	}
	go func() {
		timer := time.NewTimer(time.Minute * 1)
		for {
			select {
			case <-timer.C:
				o.rmx.Lock()
				for k, v := range o.visitIps {
					if time.Now().Sub(v) >= time.Minute*1 {
						delete(o.visitIps, k)
					}
				}
				o.rmx.Unlock()
				timer.Reset(time.Minute * 1)
			case <-ctx.Done():
				return

			}
		}
	}()

	return o
}
func (o *Ban) visit(ip string) bool {
	o.rmx.Lock()
	defer func() {
		o.rmx.Unlock()
	}()
	if _, ok := o.visitIps[ip]; ok {
		return true
	}
	o.visitIps[ip] = time.Now()
	return false
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ban := NewBan(ctx)
	success := int64(0)
	wait := &sync.WaitGroup{}
	wait.Add(1000 * 100)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func() {
				if !ban.visit(fmt.Sprintf("192.168.1.%v", j)) {
					atomic.AddInt64(&success, 1)
				}
				wait.Done()
			}()
		}
	}
	wait.Wait()
	fmt.Println("success:", success)
}
