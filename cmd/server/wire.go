//go:build wireinject
// +build wireinject

package main

import (
	"ginlayout/internal/config"
	"ginlayout/internal/handler"
	"ginlayout/internal/repository"
	"ginlayout/internal/router"
	"ginlayout/internal/service"
	"github.com/google/wire"
)

func InitializeApp() (*App, func(), error) {
	panic(wire.Build(
		config.ProviderSet,
		repository.ProviderSet,
		service.ProviderSet,
		handler.ProviderSet,
		router.ProviderSet,
		NewApp,
	))
}
