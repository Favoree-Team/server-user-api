package controller

import (
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

func (uc *userController) RegisterUser(c *gin.Context) {

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

// func (cs *userController) GetAllUser(c *gin.Context) {
// 	datas, err := cs.userService.GetAllUser()
// 	if err != nil {
// 		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
// 		return
// 	}

// 	c.JSON(200, datas)
// }

// func (cs *userController) GetUserById(c *gin.Context) {
// 	param := c.Param("id")

// 	if param == "" {
// 		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, errors.New("parameter not valid")))
// 		return
// 	}

// 	data, err := cs.userService.GetUserById(param)
// 	if err != nil {
// 		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
// 		return
// 	}

// 	c.JSON(200, data)
// }

// func (cs *userController) CreateUser(c *gin.Context) {
// 	var input entity.UserInput

// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, err))
// 		return
// 	}

// 	err := cs.userService.CreateUser(input)
// 	if err != nil {
// 		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
// 		return
// 	}

// 	c.JSON(201, gin.H{
// 		"message": "success create user",
// 	})
// }

// func (cs *userController) UpdateUserById(c *gin.Context) {
// 	var input entity.UserInput

// 	param := c.Param("id")
// 	if param == "" {
// 		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, errors.New("parameter not valid")))
// 		return
// 	}

// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, err))
// 		return
// 	}

// 	err := cs.userService.UpdateUserById(param, input)
// 	if err != nil {
// 		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
// 		return
// 	}

// 	c.JSON(200, gin.H{
// 		"message": "success update user id: " + param,
// 	})
// }

// func (cs *userController) DeleteUserById(c *gin.Context) {
// 	param := c.Param("id")
// 	if param == "" {
// 		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, errors.New("parameter not valid")))
// 		return
// 	}

// 	err := cs.userService.DeleteUserById(param)
// 	if err != nil {
// 		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
// 		return
// 	}

// 	c.JSON(200, gin.H{
// 		"message": "success delete user id: " + param,
// 	})
// }
