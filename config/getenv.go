package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	_ = godotenv.Load()
)

func GetCredentialDB() Config {
	var cred Config

	cred.Username = os.Getenv("DB_USER")
	cred.Password = os.Getenv("DB_PASS")
	cred.Host = os.Getenv("DB_HOST")
	cred.DBName = os.Getenv("DB_NAME")
	cred.Port = os.Getenv("PORT")

	return cred
}

func GetKeyJWT() string {
	return os.Getenv("JWT_SECRET")
}

// expired in minute
func GetExpiredTime() (int64, error) {
	time := os.Getenv("EXPIRED_TIME")

	expTime, err := strconv.Atoi(time)
	if err != nil {
		return 0, err
	}

	return int64(expTime), nil

}

func GetEmailCredential() (email string, password string) {
	email = os.Getenv("EMAIL_USER")
	password = os.Getenv("EMAIL_PASSWORD")

	return
}
