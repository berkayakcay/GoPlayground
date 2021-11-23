package main

import (
	"io"
	"net/http"
)

type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	io.WriteString(res, "CAT")
}

type hotdog int

func (c hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	io.WriteString(res, "DOG")
}

func main() {

	var c hotcat
	var d hotdog

	mux := http.NewServeMux()
	mux.Handle("/cat/", c)
	mux.Handle("/dog/", d)

	http.ListenAndServe(":8080", mux)
}
