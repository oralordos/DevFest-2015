package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))

	http.HandleFunc("/", index)
	http.HandleFunc("/foo", foo)
	http.HandleFunc("/bar", bar)
}

func index(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "index", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "foo", 42)
}

func bar(res http.ResponseWriter, req *http.Request) {
	model := struct {
		Name       string
		HasAccount bool
	}{
		"Daniel",
		false,
	}
	tpl.ExecuteTemplate(res, "bar", model)
}
