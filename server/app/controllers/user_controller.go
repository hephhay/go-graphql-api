package controllers

import (
	"github.com/hephhay/go-graphql/src/abstracts"
	"github.com/hephhay/go-graphql/src/services"
)

type UserController struct {
	*abstracts.BaseController
	service services.UserServiceInterface
}

// @Summary UserController constructor
func NewUserController(service services.UserServiceInterface) *UserController {
	controller := abstracts.NewBaseController(service)
	return &UserController{BaseController: controller, service: service}
}
