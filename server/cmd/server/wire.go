//+build wireinject

package server

import (
	"github.com/google/wire"
	"github.com/hephhay/go-graphql/config"
	"github.com/hephhay/go-graphql/internal/app"
	"github.com/hephhay/go-graphql/internal/repository"
	"github.com/hephhay/go-graphql/internal/service"
	"github.com/hephhay/go-graphql/internal/resolver"
)

func InitializeApp(dbCfg *config.DBCfg) (*Application, error) {
	wire.Build(
		app.NewDBConnection,
		repository.NewUserRepository,
		service.NewUserService,
		resolver.NewUserResolver,
		app.NewRoute,
		NewApplication,
	)

	return &Application{}, nil
}
