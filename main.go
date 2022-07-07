package main

import (
	"github.com/Favoree-Team/server-user-api/entity"
	"github.com/Favoree-Team/server-user-api/middleware"
	"github.com/Favoree-Team/server-user-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("status_enum", entity.ValidateTransStatusEnum)
	}

	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	routes.PingRoute(r)
	routes.UserRoute(r)
	routes.TransactionRoute(r)

	r.Run()

}
