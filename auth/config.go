package auth

import (
	"errors"
	"time"

	"github.com/Favoree-Team/server-user-api/config"
	"github.com/dgrijalva/jwt-go"
)

type AuthService interface {
	GenerateToken(UserId string, Role string, IsSubscribeBlog bool, active bool) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type authService struct{}

func NewAuthService() *authService {
	return &authService{}
}
func (s *authService) GenerateToken(UserId string, Role string, IsSubscribeBlog bool, active bool) (string, error) {
	claims := jwt.MapClaims{
		"user_id":           UserId,
		"role":              Role,
		"is_subscribe_blog": IsSubscribeBlog,
		"active":            active,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(config.GetKeyJWT()))

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *authService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(encodedToken *jwt.Token) (interface{}, error) {
		_, ok := encodedToken.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(config.GetKeyJWT()), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

func (s *authService) GenerateTokenValidate(userId string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":    userId,
		"expired_at": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(config.GetKeyJWT()))

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
