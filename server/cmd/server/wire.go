//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/hephhay/go-graphql/app"
	"github.com/hephhay/go-graphql/app/controller"
	"github.com/hephhay/go-graphql/src/repository"
	"github.com/hephhay/go-graphql/src/service"
)

func InitializeApp(appCfg *config.AppCfg, dbCfg *config.DBCfg) (*Application, error) {
	wire.Build(
		app.NewDBConnection,
		repository.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
		app.NewRoute,
		NewApplication,
	)

	return &Application{}, nil
}
