package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type future struct {
	result string
}

var f chan future

func fuc() chan future {
	return f
}

func main() {
	var t <-chan time.Time
	t = time.After(time.Second * 1)

	go func() {
		for {
			select {
			case d := <-t:
				fmt.Println(d.Second())
			case r := <-fuc():
				fmt.Println(r.result)
				return
			default:
				//fmt.Println("default")
			}
		}
	}()

	time.Sleep(time.Second * 2)
	f = make(chan future, 1)
	s := future{
		result: "aa",
	}
	f <- s

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, os.Kill, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quitChan
}
