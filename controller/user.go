package controller

import (
	"errors"
	"net/http"

	"github.com/Favoree-Team/server-user-api/entity"
	"github.com/Favoree-Team/server-user-api/service"
	"github.com/Favoree-Team/server-user-api/utils"
	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *userController {
	return &userController{
		userService: userService,
	}
}

func (uc *userController) VerifyEmailUser(c *gin.Context) {

}

func (uc *userController) RegisterUser(c *gin.Context) {
	var userInput entity.UserRegisterInput

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorHandler(utils.CreateErrorMsg(http.StatusBadRequest, err)))
		return
	}

	userRegister, err := uc.userService.RegisterUser(userInput)
	if err != nil {
		c.JSON(utils.GetErrorCode(err), utils.ErrorHandler(err))
		return
	}

	c.JSON(http.StatusCreated, userRegister)
}

func (uc *userController) LoginUser(c *gin.Context) {
	var userInput entity.UserLoginInput

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorHandler(utils.CreateErrorMsg(http.StatusBadRequest, err)))
		return
	}

	userLogin, err := uc.userService.LoginUser(userInput)
	if err != nil {
		c.JSON(utils.GetErrorCode(err), utils.ErrorHandler(err))
		return
	}

	c.JSON(http.StatusOK, userLogin)
}

func (uc *userController) GetUser(c *gin.Context) {
	userId, ok := c.Get("user_id")
	if !ok {
		errMsg := utils.CreateErrorMsg(http.StatusUnauthorized, errors.New("unauthorized user"))
		c.JSON(http.StatusUnauthorized, utils.ErrorHandler(errMsg))
		return
	}

	userDetail, err := uc.userService.GetUserID(userId.(string))
	if err != nil {
		c.JSON(utils.GetErrorCode(err), utils.ErrorHandler(err))
		return
	}

	c.JSON(http.StatusOK, userDetail)
}

func (uc *userController) UserProfileEdit(c *gin.Context) {
	userId, ok := c.Get("user_id")
	if !ok {
		errMsg := utils.CreateErrorMsg(http.StatusUnauthorized, errors.New("unauthorized user"))
		c.JSON(http.StatusUnauthorized, utils.ErrorHandler(errMsg))
		return
	}

	var userProfileEdit entity.UserProfileEdit
	if err := c.ShouldBindJSON(&userProfileEdit); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorHandler(utils.CreateErrorMsg(http.StatusBadRequest, err)))
		return
	}

	err := uc.userService.UserEditbyID(userId.(string), userProfileEdit)
	if err != nil {
		c.JSON(utils.GetErrorCode(err), utils.ErrorHandler(err))
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (uc *userController) SubscribeBlog(c *gin.Context) {
	userId, ok := c.Get("user_id")
	if !ok {
		errMsg := utils.CreateErrorMsg(http.StatusUnauthorized, errors.New("unauthorized user"))
		c.JSON(http.StatusUnauthorized, utils.ErrorHandler(errMsg))
		return
	}

	err := uc.userService.SubscribeBlog(userId.(string))
	if err != nil {
		c.JSON(utils.GetErrorCode(err), utils.ErrorHandler(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "subscribe blog success",
	})
}
