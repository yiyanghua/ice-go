package main

import (
	"errors"
	"fmt"
)

func main()  {
	defer func() {
		if err := recover();  err != nil {
			fmt.Printf("sss")
		}
	}()


	go func() {
		c :=make(chan int,1)
		c <- 1
		close(c)
		//close(c)
		panic(errors.New("ss"))
	}()

	for {
		fmt.Print("")
	}

}