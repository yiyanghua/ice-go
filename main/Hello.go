package main

import (
	"fmt"
	"sync"
)
var (
	serverContextMap   = make(map[string]*string, 8)
	serverContextMutex sync.Mutex
)

func main()  {
	fmt.Print("hello")
	var confFile string
	confFile = "serverdemo"
	ms := serverContextMap[confFile]
	if ms==nil {
		fmt.Println("---")
	}else {
		fmt.Println("####")
	}
}
