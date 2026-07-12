package main

import (
	"fmt"
	"log/slog"

	"ginlayout/internal/config"
	"ginlayout/internal/router"
	"ginlayout/pkg/logger"
)

// @title Gin Web Framework API
// @version 1.0
// @description This is a sample server.
// @host localhost:8080
// @BasePath /
func main() {
	app, cleanup, err := InitializeApp()
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize app: %v", err))
	}
	defer cleanup()

	logger.InitLogger(&app.Config.Log)
	slog.Info("Starting server...", slog.Int("port", app.Config.Server.Port))

	engine := app.Router.Setup()

	if err := engine.Run(fmt.Sprintf(":%d", app.Config.Server.Port)); err != nil {
		slog.Error("Server failed to start", slog.String("error", err.Error()))
	}
}

type App struct {
	Router *router.Router
	Config *config.Config
}

func NewApp(r *router.Router, cfg *config.Config) *App {
	return &App{
		Router: r,
		Config: cfg,
	}
}
