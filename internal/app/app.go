package app

import (
	"net/http"
	"snippetbox/internal/config"
)

type application struct {
	cfg    *config.Config
	server *http.Server
}

func (a *application) router() http.Handler {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir(a.cfg.StaticDirectory))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", a.home)
	mux.HandleFunc("/snippet/view", a.snippetView)
	mux.HandleFunc("/snippet/create", a.snippetCreate)

	return mux
}
