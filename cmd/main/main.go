package main

import (
	"log/slog"
	"net/http"
	"os"
	"snippetbox/internal/config"
	"snippetbox/internal/logger"
)

func main() {
	cfg := config.Load()

	_ = logger.SetAndReturn(cfg.Env, cfg.LogFile)
	slog.Info("Configuration loaded")

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir(cfg.StaticDirectory))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	slog.Info("Starting server", slog.String("address", cfg.Addr))
	server := &http.Server{
		Handler: mux,
		Addr:    cfg.Addr,
	}

	if err := server.ListenAndServe(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
