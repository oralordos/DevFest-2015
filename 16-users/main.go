package main

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"
)

type indexModel struct {
	Login     bool
	LoginURL  string
	LogoutURL string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))

	http.HandleFunc("/", index)
	http.HandleFunc("/profile", profile)
}

func index(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	// user.Current gives data about what the requester is
	// logged in as, or nil if they are not logged in.
	u := user.Current(ctx)

	var model indexModel

	// If they are not nil, they are logged in.
	if u != nil {
		// So let the template know, and get the logout url.
		model.Login = true
		logoutURL, err := user.LogoutURL(ctx, "/")
		if err != nil {
			log.Errorf(ctx, err.Error())
			http.Error(res, "Server Error", http.StatusInternalServerError)
			return
		}
		model.LogoutURL = logoutURL
	} else {
		// Otherwise, get the login url.
		loginURL, err := user.LoginURL(ctx, "/")
		if err != nil {
			log.Errorf(ctx, err.Error())
			http.Error(res, "Server Error", http.StatusInternalServerError)
			return
		}
		model.LoginURL = loginURL
	}
	tpl.ExecuteTemplate(res, "index", model)
}

func profile(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)
	// We don't have to check if u is nil since the app.yaml says that we
	// cannot get to this link without being logged in.
	// The ID field usually has the google username, but the test account will not care.
	tpl.ExecuteTemplate(res, "profile", u.ID)
}
