package main

import (
	"context"
	"log/slog"
	"os"
	"snippetbox/internal/app"
	"snippetbox/internal/config"
	"snippetbox/internal/logger"
)

func main() {
	cfg := config.Load()

	_ = logger.SetAndReturn(cfg.Env, cfg.LogFile)
	slog.Info("Configuration loaded")

	app := app.New(cfg)
	if err := app.Run(context.Background()); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
