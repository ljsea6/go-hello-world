package controllers

import (
	"github.com/ljsea6/go-hello-world/app/services"

	"github.com/gin-gonic/gin"
)

type IUserController interface {
	GetAllUserData(c *gin.Context)
	AddUserData(c *gin.Context)
	GetUserByID(c *gin.Context)
	UpdateUserData(c *gin.Context)
	DeleteUserByID(c *gin.Context)
}

type UserControllerImpl struct {
	svc services.IUserService
}

func (ctrl UserControllerImpl) GetAllUserData(c *gin.Context) {
	ctrl.svc.GetAllUser(c)
}

func (ctrl UserControllerImpl) AddUserData(c *gin.Context) {
	ctrl.svc.AddUserData(c)
}

func (ctrl UserControllerImpl) GetUserByID(c *gin.Context) {
	ctrl.svc.GetUserByID(c)
}

func (ctrl UserControllerImpl) UpdateUserData(c *gin.Context) {
	ctrl.svc.UpdateUserData(c)
}

func (ctrl UserControllerImpl) DeleteUserByID(c *gin.Context) {
	ctrl.svc.DeleteUserByID(c)
}

func UserControllerInit(userService services.IUserService) *UserControllerImpl {
	return &UserControllerImpl{
		svc: userService,
	}
}
