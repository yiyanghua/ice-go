package main

import (
	"sync"
	"fmt"
)

type Host interface {
	// Health checks whether the host is healthy or not
	Health() bool
}

// types.ClusterManager
type clusterManager struct {
	name string
	mux  sync.Mutex
}

func (cm *clusterManager) test() {

}

func (cm *clusterManager) Health(name string) bool {
	fmt.Println(name)
	return false
}

type A struct {
	clusterManager
}

func main() {
	a := &A{}
	a.Health("aa")
}
