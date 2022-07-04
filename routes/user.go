package routes

import (
	"github.com/Favoree-Team/server-user-api/controller"
	"github.com/Favoree-Team/server-user-api/repository"
	"github.com/Favoree-Team/server-user-api/service"

	"github.com/gin-gonic/gin"
)

var (
	userRepository = repository.NewUserRepository(DB)
	userService    = service.NewUserService(userRepository, authService, emailNotif)
	userController = controller.NewUserController(userService)
)

func UserRoute(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		user := v1.Group("/users")
		{
			user.POST("/register")
			user.POST("/login", userController.LoginUser)

			// user detail, with middleware
			user.GET("/:id")

			// user edit, with middleware
			user.PUT("/:id")

			// subscribe blog, automatic set to subscribe
			user.POST("/:id/blog_subscribe")

			// verification, using jwt code
			user.POST("/:id/verification/")
			/*
				body {
					"verification_code": "sdlajdkajdklasdjklajdkladjl"
				}
				// generate with jwt
				// get claim "code"
			*/

			user.POST("/:id/password_reset")
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
