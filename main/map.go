package main

import (
	"fmt"
	"strings"
)

type TestFun func(url string) string

type DefaultFactory struct {
	Test map[string]TestFun
}

func (factory *DefaultFactory) registry(key string, fun TestFun) {
	factory.Test[key] = fun
}

func main() {

	m := make(map[string]TestFun, 1)
	factory := DefaultFactory{Test: m}

	factory.registry("1111", func(url string) string {
		return strings.Join([]string{"http://", url}, "")
	})

	test := "www.github.com"
	if testFun, ok := factory.Test["1111"]; ok {
		value := testFun(test)
		fmt.Printf(value)
	}

}
