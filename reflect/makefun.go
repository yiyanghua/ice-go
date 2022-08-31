package main

import (
	"fmt"
	"reflect"
	"time"
)

func testMakeFunc(count int) {
	sum := 0
	for i := 0; i < count; i++ {
		sum += 1
	}
	fmt.Println(sum)
}

func main() {
	funcType := reflect.TypeOf(testMakeFunc)
	funcValue := reflect.ValueOf(testMakeFunc)

	newFunc := reflect.MakeFunc(funcType, func(args []reflect.Value) (results []reflect.Value) {
		start := time.Now()
		out := funcValue.Call(args)
		end := time.Now()
		fmt.Println(end.Sub(start))
		return out
	})
	var count int = 1e8
	newFunc.Call([]reflect.Value{reflect.ValueOf(count)})
}
