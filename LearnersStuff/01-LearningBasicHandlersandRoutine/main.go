package main

import (
	"log"
	"net/http"
)

// Writing a home handler function that writes a
// response as the response body
func home(w http.ResponseWriter, r *http.Request) {

	// Check if the current request URL path exactly matches "/". If it doesn't, use
	// the http.NotFound() function to send a 404 response to the client.
	// Importantly, we then return from the handler. If we don't return the handler
	// would keep executing and also write the "Hello from SnippetBox" message.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippet Box"))
}

// writing another route for your webpage, basically another handler function
func showSnippet(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("This is another sniipet, we show snippets here, we do shit here ya know ehehhehehehe"))
}
func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet here heheheheehheehe"))
}

func main() {
	// using the http NewSurvemux we initialize a new survemux and then registering the home handler func to it for the "/" URL path
	mux1 := http.NewServeMux()

	mux1.HandleFunc("/", home)

	// using the mux, we register the showSnippet handler func using a "/showsnippet" URL path

	mux1.HandleFunc("/showSnippet", showSnippet)
	mux1.HandleFunc("/showSnippet/create", createSnippet)

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil.
	log.Println("Starting a server on :6060")
	err := http.ListenAndServe(":6060", mux1)
	log.Fatal(err)

}

// func main() {
// 	http.HandleFunc("/", home)
// 	http.HandleFunc("/snippet", showSnippet)
// 	http.HandleFunc("/snippet/create", createSnippet)
// 	log.Println("Starting server on :4000")
// 	err := http.ListenAndServe(":4000", nil)
// 	log.Fatal(err)
// 	}

// 	Behind the scenes, these functions register their routes with something called the
// 	DefaultServeMux. There’s nothing special about this — it’s just regular servemux like we’ve
// 	already been using, but which is initialized by default and stored in a net/http global
// 	variable. Here’s the relevant line from the Go source code:

// Although this approach can make your code slightly shorter, I don’t recommend it for
// production applications.

// Because DefaultServeMux is a global variable, any package can access it and register a
// route — including any third-party packages that your application imports. If one of those
// third-party packages is compromised, they could use DefaultServeMux to expose a
// malicious handler to the web.

// So, for the sake of security, it’s generally a good idea to avoid DefaultServeMux and the
// corresponding helper functions. Use your own locally-scoped servemux instead, like we
// have been doing in this project so far.
