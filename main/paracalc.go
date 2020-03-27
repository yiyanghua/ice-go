package main

import "fmt"

func sum(values [] int, resultChan chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	resultChan <- sum // 讲结果写到chan
}

func main() {
	values := [] int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resultChan := make(chan int, 1)
	//resultChan := make(chan int, 2)
	//todo 这个1/2 有啥区别呢？很有意思
	go sum(values[:len(values)/2], resultChan)
	go sum(values[len(values)/2:], resultChan)

	sum1, sum2 := <-resultChan, <-resultChan

	fmt.Println(sum1, sum2, sum1+sum2)
}