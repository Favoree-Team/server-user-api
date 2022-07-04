package routes

import (
	"github.com/Favoree-Team/server-user-api/auth"
	"github.com/Favoree-Team/server-user-api/config"
	"github.com/Favoree-Team/server-user-api/notification"
)

var (
	DB          = config.ConnectDB()
	authService = auth.NewAuthService()
	emailNotif  = notification.NewEmailNotification()
)
