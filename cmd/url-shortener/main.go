package main

import (
	"os"
	"url-shortener/internal/config"

	"golang.org/x/exp/slog"
)

const (
	envLocal       = "local"
	envDevelopment = "development"
	envProduction  = "production"
)

func main() {
	config := config.MustLoad()

	logger := setupLogger(config.Env)

	logger.Info("Starting url-shortener", slog.String("env", config.Env))
	logger.Debug("Debug messages are enabled")
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	case envDevelopment:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	case envProduction:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	}

	return logger
}
