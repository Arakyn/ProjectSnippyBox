package main

import (
	"log"
	"net/http"
)

func main() {
	mux1 := http.NewServeMux()
	mux1.HandleFunc("/", home)
	mux1.HandleFunc("/snippet", showSnippet)
	mux1.HandleFunc("/snippet/create", create)

	// Adding a new route using http.fileserver() func which will help us serve static files over to the servers

	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir function is relative to the project
	// directory root.

	fileSever := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.

	mux1.Handle("/static/", http.StripPrefix("/static", fileSever))

	// 	With the file server working properly, we can now update the ui/html/base.layout.tmpl
	// file to make use of the static files:

	log.Println("Server Starting Up")
	err := http.ListenAndServe(":6060", mux1)
	log.Fatal(err)

}
