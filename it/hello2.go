package it

type  Hello2 struct{
	Name string
}

func (h *Hello2) Say() string {
	return h.Name
}