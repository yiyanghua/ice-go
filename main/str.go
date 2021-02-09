package main

import "fmt"

type s string

func main() {
	// https://www.slideshare.net/haoel/go-programming-patterns
	a := make([]int,32)
	b := a[1:16]
	a = append(a,1)
	a[2] =42

	fmt.Println(b[0])
	fmt.Println(b[1])
	fmt.Println(b[2])

	var d string
	d = "aa"
	fmt.Println(s(d))
}
