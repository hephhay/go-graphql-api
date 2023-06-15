package app

import (
	"github.com/hephhay/go-graphql/config"
	"github.com/hephhay/go-graphql/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewDBConnection(config *config.DBCfg) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.DBCfg), &gorm.Config{})
	if err != nil {
		log.Fatalf("gorm open database connection error")
	}
	db.AutoMigrate(&model.User{})

	return db
}
