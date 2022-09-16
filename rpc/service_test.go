package rpc

import (
	"fmt"
	"reflect"
	"testing"
)

type A struct {
	i string
}

func TestMethodType_Call(t *testing.T) {
	//fmt.Println(reflect.TypeOf((*error)(nil)).Elem())
	a := make(map[string]A)
	a["a"] = A{
		i: "a",
	}

	a["b"] = A{
		i: "b",
	}

	fmt.Print(fmt.Println(reflect.TypeOf(a).Elem()))
}
