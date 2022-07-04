package middleware

import (
	"errors"
	"net/http"

	"github.com/Favoree-Team/server-user-api/auth"
	"github.com/Favoree-Team/server-user-api/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Middleware(authService auth.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || len(tokenString) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorHandler(utils.CreateErrorMsg(http.StatusUnauthorized, errors.New("error user unauthorized"))))
			return
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorHandler(utils.CreateErrorMsg(http.StatusUnauthorized, errors.New("error user unauthorized"))))
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user_id", claims["user_id"])
			c.Set("role", claims["role"])
			c.Set("active", claims["active"])
			c.Set("is_subscribe_blog", claims["is_subscribe_blog"])

		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorHandler(utils.CreateErrorMsg(http.StatusUnauthorized, errors.New("error user unauthorized"))))
			return
		}

		c.Next()
	}
}
