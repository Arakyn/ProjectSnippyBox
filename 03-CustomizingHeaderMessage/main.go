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
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("welcome to the home page"))
}
func showsnippet(w http.ResponseWriter, r *http.Request) {

	// Extract the value of the id parameter from the query string and try to
	// convert it to an integer using the strconv.Atoi() function. If it can't
	// be converted to an integer, or the value is less than 1, we return a 404 page
	// not found response.

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// Use the fmt.Fprintf() function to interpolate the id value with our response
	// and write it to the http.ResponseWriter

	fmt.Fprintf(w, "This is the following thing you asked for %d..\n", id)
	w.Write([]byte("This is the home page bitchass\n"))
}
func create(w http.ResponseWriter, r *http.Request) {

	// Suppressing System-Generated Headers
	// The Del() method doesn’t remove system-generated headers. To suppress these, you need
	// to access the underlying header map directly and set the value to nil. If you want to
	// suppress the Date header, for example, you need to write:
	w.Header()["Date"] = nil
	w.Header()["Content-Length"] = nil
	w.Header()["X-Content-Type-Options"] = nil

	if r.Method != http.MethodPost {

		w.Header().Set("Allowed", http.MethodPost)

		// The http.Error Shortcut
		// If you want to send a non-200 status code and a plain-text response body (like we are in the
		// 	code above) then it’s a good opportunity to use the http.Error() shortcut. This is a
		// 	lightweight helper function which takes a given message and status code, then calls the
		// 	w.WriteHeader() and w.Write() methods behind-the-scenes for us.
		w.Header()["Date"] = nil
		w.Header()["Content-Length"] = nil
		w.Header()["X-Content-Type-Options"] = nil

		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))
		http.Error(w, "Method Not Allowed BitchAss", 405)

		return

		// // Use the http.Error() function to send a 405 status code and "Method Not
		// // Allowed" string as the response body.
		//
	}
	w.Header().Set("Content-type", "application/json")

	w.Write([]byte(`{"name":"Alex"}`))
	w.Write([]byte("\nCreate Snippets Here"))
}
func main() {
	servemux := http.NewServeMux()
	servemux.HandleFunc("/", home)
	servemux.HandleFunc("/snippet1", showsnippet)
	servemux.HandleFunc("/snippet1/create", create)
	log.Println("Starting the server")
	err := http.ListenAndServe(":6060", servemux)
	log.Fatal(err)
}
