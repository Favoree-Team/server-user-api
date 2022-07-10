package main

import (
	"github.com/Favoree-Team/server-user-api/middleware"
	"github.com/Favoree-Team/server-user-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	routes.PingRoute(r)
	routes.UserRoute(r)
	routes.TransactionRoute(r)

	r.Run()

}
