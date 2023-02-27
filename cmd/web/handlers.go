package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/homepage.tmpl",
		"./ui/html/baselayout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	// defining a URL path, if the URL path is not like this , it will show a HTTP Not found error
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	// This is used to show a Webpage on the website
	ts, err := template.ParseFiles(files...)
	if err != nil {

		app.serverError(w, err)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
		return

	}

}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		app.errorlog.Println(err)
		return
	}
	fmt.Fprintf(w, "\nWelcome to SnippyShow, here is what you asked for, %d...", id)

}

func (app *application) create(w http.ResponseWriter, r *http.Request) {

	w.Header()["Date"] = nil
	if r.Method != http.MethodPost {
		w.Header().Set("Allowed :", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)

		return
	}
	w.Write([]byte("Here you can create Snippetppts"))

}
