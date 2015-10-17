package main

import (
	"io"
	"net/http"
)

func init() {
	// If we have more handlers, we will go to the more specific page.
	http.HandleFunc("/", hello)
	http.HandleFunc("/foo", foo)
}

func hello(res http.ResponseWriter, req *http.Request) {
	// Web-browsers look at this header to determine what the response is.
	// Since we are sending back html, we need to let the browser know.
	res.Header().Set("Content-Type", "text/html")
	io.WriteString(res, "<a href=\"/foo\">This is a link</a>")
}

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Bar")
}
