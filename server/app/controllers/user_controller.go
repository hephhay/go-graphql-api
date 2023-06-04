package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hephhay/go-graphql/src/abstracts"
	"github.com/hephhay/go-graphql/src/services"
)

type UserController struct {
	*abstracts.BaseController
	service service.UserServiceInterface
}

// @Summary UserController constructor
func NewUserController(service service.UserServiceInterface) *UserController {
	controller := abstracts.NewBaseController(service)
	return &UserController{BaseController: controller, service: service}
}
