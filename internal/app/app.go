package app

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"snippetbox/internal/config"
)

type application struct {
	cfg    *config.Config
	server *http.Server
}

func (a *application) Run(ctx context.Context) {
	slog.Info("Starting server", slog.String("address", a.cfg.Addr))
	if err := a.server.ListenAndServe(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
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
