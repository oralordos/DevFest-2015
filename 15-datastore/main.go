package main

import (
	"html/template"
	"net/http"
	"time"

	// There are a couple different valid imports, these ones don't mess up my auto complete.
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

type message struct {
	Name, Message string
	Posted        time.Time
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))

	http.HandleFunc("/", index)
	// This will redirect every request starting with /static/ into the static directory on the disk.
	// This will also make precautions against malformed requests to prevent the system from attacks.
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
}

func index(res http.ResponseWriter, req *http.Request) {
	// Most web browsers will look for a favicon.ico automatically for the website's icon.
	// This will make sure that everything besides the basic web page will get a 404.
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	// An AppEngine context is basically everything AppEngine needs to know about the request,
	// and nothing more.
	ctx := appengine.NewContext(req)

	// If we are getting a post request, instead of a get request, we will also want to add
	// the new message to our database.
	if req.Method == "POST" {
		// This will get the values from a form, it will work for both get and post.
		name := req.FormValue("name")
		val := req.FormValue("message")
		// This is how we get a key to the datastore when we don't care about the actual key value.
		// The arguements are the context, the kind (table) in the datastore, and the parent.
		// The parent is used to chain related data together, and has some benefits.
		key := datastore.NewIncompleteKey(ctx, "messages", nil)
		// An easy way to timestamp data is to add a time.Time field to the struct, and
		// fill it with time.Now before saving it to the datastore.
		m := message{name, val, time.Now()}
		// This will put data in the datastore, overwriting whatever was at that key before.
		// It returns the full key, and an error.
		_, err := datastore.Put(ctx, key, &m)
		if err != nil {
			// If there has been an error, we should log the error for our use, inform the
			// user that there was a problem, and exit our function so we don't keep going.
			log.Errorf(ctx, err.Error())
			http.Error(res, "Server Error", http.StatusInternalServerError)
			return
		}
	}

	// This will look stuff up in the datastore.
	// NewQuery takes the kind (table) to look stuff up in.
	// Order takes a field in the struct to order by, with a - prefix to reverse the order.
	// Limit takes a number and will get a maximum of that many results.
	q := datastore.NewQuery("messages").Order("-Posted").Limit(30)
	// We need a spot to put the data when it is found.
	var messages []message
	// The GetAll method takes the context, and a pointer to a slice to put the results.
	// It will return the keys to each item, and an error.
	_, err := q.GetAll(ctx, &messages)
	// Don't forget to handle the errors.
	if err != nil {
		log.Errorf(ctx, err.Error())
		http.Error(res, "Server Error", http.StatusInternalServerError)
		return
	}
	// ExecuteTemplate also returns an error that we should handle.
	err = tpl.ExecuteTemplate(res, "index", messages)
	if err != nil {
		log.Errorf(ctx, err.Error())
		http.Error(res, "Server Error", http.StatusInternalServerError)
		return
	}
}
