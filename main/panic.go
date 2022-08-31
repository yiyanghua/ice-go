package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go func() {
		defer func() {
			fmt.Println("111")
			if err := recover(); err != nil {
				fmt.Println("sss")
			}
		}()
		c := make(chan int, 1)
		c <- 1
		close(c)
		panic(errors.New("sfs"))
	}()

	// wait
	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Kill, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	fmt.Println("dddd")
	<-quitCh
}
