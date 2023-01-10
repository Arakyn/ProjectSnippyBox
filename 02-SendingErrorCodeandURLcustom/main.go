package main

// Let’s now update our application so that the /snippet/create route only responds to HTTP
// requests which use the POST method, like so

// Making this change is important because — later in our application build — requests to the
// /snippet/create route will result in a new snippet being created in a database. Creating a
// new snippet in a database is a non-idempotent action that changes the state of our server,
// so we should follow HTTP good practice and restrict this route to act on POST requests only.

// Let’s begin by updating our createSnippet() handler function so that it sends a 405
// (method not allowed) HTTP status code unless the request method is POST. To do this we’ll
// need to use the w.WriteHeader() method like so

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("welcome to the home page"))
}
func showsnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the home page bitchass"))
}
func create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {

		// Use r.Method to check whether the request is using POST or not. Note that
		// http.MethodPost is a constant equal to the string "POST".

		// If it's not, use the w.WriteHeader() method to send a 405 status
		// code and the w.Write() method to write a "Method Not Allowed"
		// response body. We then return from the function so that the
		// subsequent code is not executed.
		w.Header().Set("Allowed", http.MethodPost)
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return

	}
	w.Write([]byte("Here ya make snippetsssssss"))
}
func main() {
	servemux := http.NewServeMux()
	servemux.HandleFunc("/", home)
	servemux.HandleFunc("/snippet1.com", showsnippet)
	servemux.HandleFunc("/snippet1.com/create", create)
	log.Println("Starting the server")
	err := http.ListenAndServe(":6060", servemux)
	log.Fatal(err)
}
