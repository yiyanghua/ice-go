package main

import (
	"net/http"
)

type name struct {
}

func (n name) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Cookies()
	w.Write([]byte("11111"))
}

func hf(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hf"))
}

func main() {
	mx := http.NewServeMux()
	n := &name{}
	mx.Handle("/aa", n)
	//http.ListenAndServe(":8080", mx)

	http.HandleFunc("/hf", hf)
	http.Handle("/aa", http.HandlerFunc(hf))

	http.ListenAndServe(":8080", nil)
}
