package logger

import (
	"io"
	"log/slog"
	"os"
	"snippetbox/internal/config"
)

func SetAndReturn(env string, logFile string) *slog.Logger {
	var slogHandler slog.Handler

	switch env {
	case config.EnvLocal:
		slogHandler = slog.NewTextHandler(os.Stdout, nil)
	case config.EnvStaging, config.EnvProd:
		f, err := os.Open(logFile)
		if err != nil {
			panic(err.Error())
		}

		writer := io.MultiWriter(os.Stdout, f)
		slogHandler = slog.NewJSONHandler(writer, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelWarn,
		})
	}

	logger := slog.New(slogHandler)
	slog.SetDefault(logger)

	return logger
}
