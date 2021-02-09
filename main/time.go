package main

import (
	"fmt"
	"time"
)

func main() {
	a := time.Now()
	b := a.Add(- time.Minute * 3)

	astr := a.Format(time.RFC3339)
	fmt.Println(a.Format(time.RFC3339))
	fmt.Println(b.Format(time.RFC3339))
	fmt.Println(time.Parse(time.RFC3339,astr))
}