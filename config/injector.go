// go:build wireinject
//go:build wireinject
// +build wireinject

package config

import (
	"github.com/ljsea6/go-hello-world/app/controllers"
	"github.com/ljsea6/go-hello-world/app/repositories"
	"github.com/ljsea6/go-hello-world/app/services"

	"github.com/google/wire"
)

var db = wire.NewSet(ConnectToDB)

var userServiceSet = wire.NewSet(services.UserServiceInit,
	wire.Bind(new(services.IUserService), new(*services.UserServiceImpl)),
)

var userRepoSet = wire.NewSet(repositories.UserRepositoryInit,
	wire.Bind(new(repositories.IUserRepository), new(*repositories.UserRepositoryImpl)),
)

var userCtrlSet = wire.NewSet(controllers.UserControllerInit,
	wire.Bind(new(controllers.IUserController), new(*controllers.UserControllerImpl)),
)

func Init() *Initialization {
	wire.Build(NewInitialization, db, userCtrlSet, userServiceSet, userRepoSet)
	return nil
}
