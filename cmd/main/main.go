package main

import (
	"log"
	"net/http"
)

func main() {
	addr := ":4000"

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("Starting server on %s", addr)
	server := &http.Server{
		Handler: mux,
		Addr:    addr,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
