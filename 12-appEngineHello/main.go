package main

import (
	"io"
	"net/http"
)

func init() {
	http.HandleFunc("/", hello)
}

func hello(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello World!")
}
