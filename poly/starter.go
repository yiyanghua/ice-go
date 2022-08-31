package main

import (
	"github.com/yiyanghua/ice-go/poly/it"
)

func main() {
	h1 := &it.Hello1{
		Name: "1",
	}
	var hh1 it.Hello = h1

	h2 := &it.Hello2{
		Name: "2",
	}
	var hh2 it.Hello = h2

	s := make([]*it.Hello, 0)
	s = append(s, &hh1, &hh2)

	for h := range s {
		println(h)
	}
}
