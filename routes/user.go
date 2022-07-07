package routes

import (
	"github.com/Favoree-Team/server-user-api/controller"
	"github.com/Favoree-Team/server-user-api/repository"
	"github.com/Favoree-Team/server-user-api/service"

	"github.com/gin-gonic/gin"
)

var (
	userRepository     = repository.NewUserRepository(DB)
	ipRecordRepository = repository.NewIPRecordRepository(DB)
	userService        = service.NewUserService(userRepository, ipRecordRepository, authService, emailNotif)
	userController     = controller.NewUserController(userService)
)

func UserRoute(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		user := v1.Group("/users")
		{

			user.POST("/ip_address/check", userController.IPAddressCheck)

			user.POST("/ip_address/create", mainMiddleware, userController.IPAddressCreate)

			user.POST("/register", userController.RegisterUser)
			user.POST("/login", userController.LoginUser)

			// user detail, with middleware
			user.GET("/detail", mainMiddleware, userController.GetUser)

			// user edit, with middleware
			user.PUT("/edit", mainMiddleware, userController.UserProfileEdit)

			// subscribe blog, automatic set to subscribe
			user.POST("/blog_subscribe", mainMiddleware, userController.SubscribeBlog)

			// verification, using jwt code
			user.POST("/verification", mainMiddleware)
			/*
				body {
					"verification_code": "sdlajdkajdklasdjklajdkladjl"
				}
				// generate with jwt
				// get claim "code"
			*/

			user.POST("/password_reset", mainMiddleware)
			/*
				body {
					"new_password": "12345678"
				}
			*/

			// PHASE 1
			// with middleware
			// user.POST("/:id/deactivate")

			// // without middleware
			// user.POST("/:id/activate")
		}
	}
}
