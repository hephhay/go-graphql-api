package repositories

import (
	"gorm.io/gorm"
	"github.com/hephhay/go-graphql/src/abstracts"
	"github.com/hephhay/go-graphql/src/models"
)

type UserRepositoryInterface interface {
	abstracts.BaseRepositoryInterface
}

type UserRepository struct {
	*abstracts.BaseRepository
}

// @Summary UserRepository constructor
func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	var model *models.User
	baseRepo := abstracts.NewBaseRepository(db, &model)
	return &UserRepository{baseRepo}
}
