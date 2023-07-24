package services

import (
	"net/http"
	"strconv"

	"github.com/ljsea6/go-hello-world/app/constants"
	"github.com/ljsea6/go-hello-world/app/domain/dao"
	"github.com/ljsea6/go-hello-world/app/pkg"
	"github.com/ljsea6/go-hello-world/app/repositories"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	GetAllUser(c *gin.Context)
	GetUserByID(c *gin.Context)
	AddUserData(c *gin.Context)
	UpdateUserData(c *gin.Context)
	DeleteUserByID(c *gin.Context)
}

type UserServiceImpl struct {
	userRepository repositories.IUserRepository
}

func (s UserServiceImpl) UpdateUserData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program update user data by id")
	userID, _ := strconv.Atoi(c.Param("userID"))

	var request dao.User
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constants.InvalidRequest)
	}

	data, err := s.userRepository.FindUserByID(userID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constants.DataNotFound)
	}

	data.RoleID = request.RoleID
	data.Email = request.Email
	data.Name = request.Password
	data.Status = request.Status
	s.userRepository.Save(&data)

	if err != nil {
		log.Error("Happened error when updating data to database. Error", err)
		pkg.PanicException(constants.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constants.Success, data))
}

func (s UserServiceImpl) GetUserByID(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get user by id")
	userID, _ := strconv.Atoi(c.Param("userID"))

	data, err := s.userRepository.FindUserByID(userID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constants.DataNotFound)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constants.Success, data))
}

func (s UserServiceImpl) AddUserData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program add data user")
	var request dao.User
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constants.InvalidRequest)
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 15)
	request.Password = string(hash)

	data, err := s.userRepository.Save(&request)
	if err != nil {
		log.Error("Happened error when saving data to database. Error", err)
		pkg.PanicException(constants.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constants.Success, data))
}

func (s UserServiceImpl) GetAllUser(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute get all data user")

	data, err := s.userRepository.FindAllUser()
	if err != nil {
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constants.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constants.Success, data))
}

func (s UserServiceImpl) DeleteUserByID(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute delete data user by id")
	userID, _ := strconv.Atoi(c.Param("userID"))

	err := s.userRepository.DeleteUserByID(userID)
	if err != nil {
		log.Error("Happened Error when try delete data user from DB. Error:", err)
		pkg.PanicException(constants.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constants.Success, pkg.Null()))
}

func UserServiceInit(userRepository repositories.IUserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}
