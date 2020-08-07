package it

type  Hello1 struct{
	Name string
}

func (h Hello1) say() string {
	return h.Name
}