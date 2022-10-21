package main

import (
	"fmt"
	"net/http"
)

type logger struct {
	handler http.Handler
}

func (l logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("xxx")
	l.handler.ServeHTTP(w, r)
}
func NewLog(handler http.Handler) logger {
	return logger{handler: handler}
}

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
	l := NewLog(mx)
	n := &name{}
	mx.Handle("/aa", n)
	http.ListenAndServe(":8080", l)

	/*http.HandleFunc("/hf", hf)
	http.Handle("/aa", http.HandlerFunc(hf))

	http.ListenAndServe(":8080", nil)*/
}
