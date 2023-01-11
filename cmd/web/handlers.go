package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	// defining a URL path, if the URL path is not like this , it will show a HTTP Not found error
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	ts, err := template.ParseFiles("./ui/html/homepage.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	w.Write([]byte("Welcome to the Home Page, SnippyBox\n"))
	w.Write([]byte("We are really happy to Have you Here"))

}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "\nWelcome to SnippyShow, here is what you asked for, %d...", id)

}

func create(w http.ResponseWriter, r *http.Request) {

	w.Header()["Date"] = nil
	if r.Method != http.MethodPost {
		w.Header().Set("Allowed :", http.MethodPost)
		http.Error(w, "Method Not Allowed Please Use Another", 405)
		return
	}
	w.Write([]byte("Here you can create Snippetppts"))

}
