package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Favoree-Team/server-user-api/config"
	"gorm.io/gorm"
)

func main() {
	db := config.ConnectDB()

	var checkFlag string

	for _, arg := range os.Args[1:] {
		checkFlag += arg
	}

	fmt.Println("execute code :", checkFlag)

	switch checkFlag {
	case "migrate_db":
		// excute create table
		ExecuteQueries(db, "./migration/table.sql")
		ExecuteQueries(db, "./migration/seed.sql")
	case "seed_test":
		// excute seed test
		ExecuteQueries(db, "./migration/seed_test.sql")
	case "drop_db":
		// drop tables
		ExecuteQueries(db, "./migration/drop.sql")
	default:
		break
	}
}

func Err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ExecuteQueries(db *gorm.DB, pathFile string) {
	dat, err := os.ReadFile(pathFile)
	Err(err)

	listExecs := strings.Split(string(dat), ";")

	for _, qExec := range listExecs[:len(listExecs)-1] {
		if err := db.Exec(qExec).Error; err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("success execute", qExec)
		}
	}
}
