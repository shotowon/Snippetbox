package main

import "net/http"

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, from Snippetbox"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create snippet..."))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("View specific snippet..."))
}
