package user

import (
	"net/http"
	"strconv"
	"user-res-api/dto"
	"user-res-api/service"
	e "user-res-api/utils/errors"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetUserById(c *gin.Context) {
	log.Debug("User id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var userDto *dto.UserDto

	userDto, err := service.UserService.GetUserById(id)

	if err != nil {
		apiErr, ok := err.(e.ApiError)
		if !ok {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	c.JSON(http.StatusOK, userDto)
}

func GetUsers(c *gin.Context) {
	var usersDto dto.UsersDto
	usersDto, err := service.UserService.GetUsers()

	if err != nil {
		apiErr, ok := err.(e.ApiError)
		if !ok {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	c.JSON(http.StatusOK, usersDto)
}

func UserInsert(c *gin.Context) {
	var userDto dto.UserDto
	err := c.BindJSON(&userDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userDtoPtr, er := service.UserService.InsertUser(&userDto)
	// Error del Insert
	if er != nil {
		apiErr, ok := er.(e.ApiError)
		if !ok {
			c.JSON(http.StatusInternalServerError, er.Error())
			return
		}
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	c.JSON(http.StatusCreated, userDtoPtr)
}

func Login(c *gin.Context) {
	var loginDto dto.LoginDto
	er := c.BindJSON(&loginDto)

	if er != nil {
		log.Error(er.Error())
		c.JSON(http.StatusBadRequest, er.Error())
		return
	}
	log.Debug(loginDto)

	var loginResponseDto *dto.LoginResponseDto
	loginResponseDto, err := service.UserService.Login(&loginDto)
	if err != nil {
		apiErr, ok := err.(e.ApiError)
		if !ok {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		if apiErr.Status() == 400 {
			c.JSON(http.StatusBadRequest, apiErr.Error())
			return
		}
		c.JSON(http.StatusForbidden, apiErr.Error())
		return
	}

	c.JSON(http.StatusOK, loginResponseDto)
}

func Refresh(c *gin.Context) {
	var refreshTokenDto dto.RefreshTokenDto
	err := c.BindJSON(&refreshTokenDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	response, err := service.UserService.Refresh(&refreshTokenDto)
	if err != nil {
		apiErr, ok := err.(e.ApiError)
		if !ok {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	c.JSON(http.StatusOK, response)
}
