// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/hephhay/go-graphql/database"
)

func InitializeApp(connectionInfo string) error {
	wire.Build(database.NewDBConnection)

	return nil
}
