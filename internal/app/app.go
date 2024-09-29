package app

import (
	"context"
	"log/slog"
	"net/http"
	"snippetbox/internal/config"
	"snippetbox/internal/gears/db"
	"snippetbox/internal/storage"
)

type application struct {
	cfg     *config.Config
	server  *http.Server
	storage *storage.Storage
}

func New(cfg *config.Config) *application {
	app := &application{
		cfg: cfg,
	}

	router := app.router()
	server := &http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	app.server = server
	return app
}

func (a *application) Run(ctx context.Context) error {
	db, err := db.OpenPostgres(a.cfg.MainDSN)
	if err != nil {
		return err
	}
	slog.Info("connected to the database", slog.String("db", a.cfg.MainDSN))
	a.storage = storage.New(db)

	slog.Info("Starting server", slog.String("address", a.cfg.Addr))
	return a.server.ListenAndServe()
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
