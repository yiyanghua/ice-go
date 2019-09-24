package main

import "fmt"

type PersonInfo struct {
	ID string
	Name string
}

func main()  {
	var personDb map[string] PersonInfo
	personDb = make(map[string] PersonInfo)

	personDb["sd"] = PersonInfo{"sd","sd"}

	if person, ok := personDb["sd"]; ok{
		fmt.Println(person)
	}

}