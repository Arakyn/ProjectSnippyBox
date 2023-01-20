package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// Define an application struct to hold the application-wide dependencies for the
// web application. For now we'll only include fields for the two custom loggers, but
// we'll add more to it as the build progresses.

type application struct {
	errorlog *log.Logger
	infolog  *log.Logger
}

func main() {

	InfoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	ErrorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := *&application{
		errorlog: ErrorLog,
		infolog:  InfoLog,
	}

	mux1 := http.NewServeMux()
	mux1.HandleFunc("/", app.home)
	mux1.HandleFunc("/snippet", app.showSnippet)
	mux1.HandleFunc("/snippet/create", app.create)

	// Adding a new route using http.fileserver() func which will help us serve static files over to the servers

	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir function is relative to the project
	// directory root.
	addr := flag.String("addr", ":4040", "HTTP network Address")

	flag.Parse()

	fileSever := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.

	mux1.Handle("/static/", http.StripPrefix("/static", fileSever))

	// 	With the file server working properly, we can now update the ui/html/base.layout.tmpl
	// file to make use of the static files:

	// Use log.New() to create a logger for writing information messages. This takes
	// three parameters: the destination to write the logs to (os.Stdout), a string
	// prefix for message (INFO followed by a tab), and flags to indicate what
	// additional information to include (local date and time). Note that the flags
	// are joined using the bitwise OR operator |.

	srv := &http.Server{

		Addr:     *addr,
		ErrorLog: ErrorLog,
		Handler:  mux1,
	}

	InfoLog.Printf("Server Starting Up on %s", *addr)

	err := srv.ListenAndServe()
	app.errorlog.Fatal(err)

}
