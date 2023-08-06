package app

import (
	"log"

	"github.com/hephhay/go-graphql-api/config"
	"github.com/hephhay/go-graphql-api/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConnection(config *config.DBCfg) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.DBCfg), &gorm.Config{})
	if err != nil {
		log.Fatalf("gorm open database connection error")
	}
	db.AutoMigrate(&model.User{})

	return db
}
