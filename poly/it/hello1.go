package it

type Hello1 struct {
	Name string
}

func (h *Hello1) Say() string {
	return h.Name
}
