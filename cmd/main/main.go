package main

import (
	"context"
	"log/slog"
	"snippetbox/internal/app"
	"snippetbox/internal/config"
	"snippetbox/internal/logger"
)

func main() {
	cfg := config.Load()

	_ = logger.SetAndReturn(cfg.Env, cfg.LogFile)
	slog.Info("Configuration loaded")

	app := app.New(cfg)
	app.Run(context.Background())
}
