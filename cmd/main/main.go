package main

import (
	"log"
	"net/http"
	"snippetbox/internal/config"
)

func main() {
	cfg := config.Load()

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir(cfg.StaticDirectory))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("Starting server on %s", cfg.Addr)
	server := &http.Server{
		Handler: mux,
		Addr:    cfg.Addr,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
