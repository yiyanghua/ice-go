package main

import (
	"sync"
)

type Host interface {
	// Health checks whether the host is healthy or not
	Health(name string) bool
}

// types.ClusterManager
type clusterManager struct {
	name string
	mux  sync.Mutex
}

func (cm *clusterManager) Health(name string) bool {
	cm.name = name
	return false
}

func main() {
	//a := new(clusterManager)
	a := clusterManager{name: "aa"}
	a.Health("aaa")

	println(a.name)

	//a :=clusterManager{name: "aa"}
	b := (interface{})(a).(Host)
	b.Health("aa")

	// Host已经是接口了， 指针都隐藏了
	// 接口内部会存储对象的实际地址， 因为接口在做类型判断比对完，就把实际地址存起来，方便通过接口对象调用方法。

	/**
	// i.(T)
	a := &clusterManager{}
	_, ok  := (interface{})(a).(*Host)// 错误
	if !ok {
		fmt.Println("convert type err")
	}
	_, ok  = (interface{})(a).(*clusterManager)// 对
	if !ok {
		fmt.Println("convert type err")

	}
	*/
}
