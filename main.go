package main

import (
	"github.com/Favoree-Team/server-user-api/entity"
	"github.com/Favoree-Team/server-user-api/notification"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("status_enum", entity.ValidateTransStatusEnum)
	}

	// r := gin.Default()

	// r.Use(middleware.CORSMiddleware())

	// routes.UserRoute(r)
	// routes.TransactionRoute(r)

	// r.Run()

	var (
		emailNotif = notification.NewEmailNotification()
	)

	emailNotif.SendVerification("afistapratama@gmail.com", "https://google.com")

}
