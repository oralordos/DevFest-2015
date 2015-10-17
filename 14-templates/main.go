package main

import (
	"html/template"
	"net/http"
)

// Hold your template data globally, it makes it easier.
var tpl *template.Template

func init() {
	// During startup, you will need to read in all the template files.
	// This will send back a pointer to the templates, and an error.
	// The Must function takes both of those, and if there is an error,
	// prints an error message and quits. If there was no error,
	// it will simply return just the template.
	tpl = template.Must(template.ParseGlob("*.gohtml"))

	http.HandleFunc("/", index)
	http.HandleFunc("/foo", foo)
	http.HandleFunc("/bar", bar)
}

func index(res http.ResponseWriter, req *http.Request) {
	// The three arguments is, a Writer that the template will be sent to,
	// the name of the template to use, and the data to put into the template.
	// ResponseWriter is smart enough to recognize a full html page and will set Content-Type by itself.
	tpl.ExecuteTemplate(res, "index", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	// Here we are sending some data into the template.
	tpl.ExecuteTemplate(res, "foo", 42)
}

func bar(res http.ResponseWriter, req *http.Request) {
	// Here we are creating an anonymous struct type and variable and filling it.
	model := struct {
		Name       string
		HasAccount bool
	}{
		"Daniel",
		false,
	}
	// Here we pass the struct to the template.
	// Note that sending a pointer in more efficient than the value for structs.
	tpl.ExecuteTemplate(res, "bar", &model)
}
