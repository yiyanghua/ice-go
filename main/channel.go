package main

import "fmt"

func Count(ch chan int) {
	ch <- 1
	fmt.Println("Counting")
}

func main() {
	chs := initChan2().([]chan int)

	for i := 0; i < 10; i++ {
		go Count(chs[i])
	}

	for i, ch := range chs {
		fmt.Printf("i: %d v: %d \n", i, <-ch)
	}
}

func initChan2() interface{} {
	chs := []chan int{
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
	}
	return chs
}

func initChan() interface{} {
	chs := make([]chan int, 10)
	for i := range chs {
		chs[i] = make(chan int)
	}
	return chs
}

func initChan3()  {
	flag := make(chan bool)

	go func() {
		flag <- true
		a :=<-flag
		fmt.Printf("1 %v\n", a)
	}()

	go func() {
		a := <-flag
		flag <- false
		fmt.Printf("2 %v",a)
	}()
}