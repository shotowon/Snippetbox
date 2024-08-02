package main

import "net/http"

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, from Snippetbox"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create snippet..."))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("View specific snippet..."))
}
