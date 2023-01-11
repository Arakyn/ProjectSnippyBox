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
	log.Println("Server Starting Up")
	err := http.ListenAndServe(":6060", mux1)
	log.Fatal(err)

}
