// AppEngine will deal with loading our code, so we do not have to name this main.
package main

import (
	"io"
	"net/http"
)

// All web requests will go through a function with these parameters.
func hello(res http.ResponseWriter, req *http.Request) {
	// WriteString takes in a io.Writer interface, and a string, and writes the string to the writer.
	io.WriteString(res, "Hello World!")
}

// We must use init instead of main, AppEngine have it's own main, we have to make due.
func init() {
	// This registers our hello function to every request.
	http.HandleFunc("/", hello)
}
