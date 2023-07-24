package config

import (
	"github.com/ljsea6/go-hello-world/app/controllers"
	"github.com/ljsea6/go-hello-world/app/repositories"
	"github.com/ljsea6/go-hello-world/app/services"
)

type Initialization struct {
	userRepo repositories.IUserRepository
	userSvc  services.IUserService
	UserCtrl controllers.IUserController
}

func NewInitialization(
	userRepo repositories.IUserRepository,
	userService services.IUserService,
	userCtrl controllers.IUserController,
) *Initialization {
	return &Initialization{
		userRepo: userRepo,
		userSvc:  userService,
		UserCtrl: userCtrl,
	}
}
