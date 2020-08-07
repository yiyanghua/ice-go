package it

type  Hello2 struct{
	Name string
}

func (h Hello2) say() string {
	return h.Name
}