//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/hephhay/go-graphql/app/controllers"
	"github.com/hephhay/go-graphql/configs"
	"github.com/hephhay/go-graphql/database"
	"github.com/hephhay/go-graphql/src/repositories"
	"github.com/hephhay/go-graphql/src/services"
	"github.com/hephhay/go-graphql/app"
)

func InitializeApp(dbCfg *configs.DBCfg) (*Application, error){
	wire.Build(
		database.NewDBConnection,
		repositories.NewUserRepository,
		services.NewUserService,
		controllers.NewUserController,
		app.NewRouter,
		NewApplication,
	)

	return nil, nil
}
