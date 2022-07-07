package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Username string
	Password string
	Host     string
	DBName   string
	Port     string
	SSL      string
}

func FailOnError(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}

func ConnectDB() *gorm.DB {
	var cred = GetCredentialDB()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", cred.Username, cred.Password, cred.Host, cred.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	FailOnError(err)

	return db
}
