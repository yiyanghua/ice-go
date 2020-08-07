package main

import (
	"ice-go/it"
)

func main()  {
	h1 := &it.Hello1{
		Name: "1",
	}
	hh1 := interface{}(*h1).(it.Hello)


	h2 := &it.Hello2{
		Name: "2",
	}
	hh2 := interface{}(*h2).(it.Hello)


	s := make([]*it.Hello, 0)
	s = append(s, &hh1,&hh2)

	for h,_ := range s {
		println(h)
	}
}