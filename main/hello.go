package main

import "fmt"

type TestFun func(url string) string

type DefaultFactory struct {
	test map[string]TestFun
}

func (factory *DefaultFactory) registry(key string, fun TestFun) {
	factory.test[key] = fun
}

func main() {

	m := make(map[string]TestFun, 2)
	factory := DefaultFactory{test: m}

	factory.registry("1111", func(url string) string {
		return url
	})

	factory.registry("2222", func(url string) string {
		return url
	})

	test := "test12"
	if testFun, ok := factory.test["1111"]; ok {
		value := testFun(test)
		fmt.Printf("########%s\n", value)
	}

}
