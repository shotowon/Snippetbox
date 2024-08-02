package main

import (
	"log"
	"net/http"
)

func main() {
	addr := ":4000"

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Printf("Starting server on %s", addr)
	server := &http.Server{
		Handler: mux,
		Addr:    addr,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
