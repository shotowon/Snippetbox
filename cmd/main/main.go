package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, from Snippetbox"))
}
