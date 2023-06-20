// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"github.com/hephhay/go-graphql/config"
	"github.com/hephhay/go-graphql/internal/app"
	"github.com/hephhay/go-graphql/internal/repository"
	"github.com/hephhay/go-graphql/internal/service"
)

// Injectors from wire.go:

func InitializeApp(dbCfg *config.DBCfg) (*Application, error) {
	db := app.NewDBConnection(dbCfg)
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	engine := app.NewRoute(userService)
	application := NewApplication(engine)
	return application, nil
}
