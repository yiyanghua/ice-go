package main

import (
	"sync"
	"fmt"
	"runtime"
)

var counter = 0

func Counter(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println(counter)
	lock.Unlock()
}

func main() {
	lock := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		go Counter(lock)
	}

	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched()
		if (c > 10) {
			break
		}
	}

}
