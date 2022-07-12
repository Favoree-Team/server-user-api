package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Favoree-Team/server-user-api/entity"
)

// example : F-202206110001

func GetOrderNow(last entity.Transaction) (string, error) {
	var (
		mm string
		dd string
	)

	now := time.Now()

	if int(now.Month()) < 10 {
		mm = fmt.Sprintf("0%d", int(now.Month()))
	} else {
		mm = fmt.Sprintf("%d", int(now.Month()))
	}

	if now.Day() < 10 {
		dd = fmt.Sprintf("0%d", now.Day())
	} else {
		dd = fmt.Sprintf("%d", now.Day())
	}

	if last.OrderID == "" && last.ID == "" {
		return fmt.Sprintf("F-%d%s%s0001", now.Year(), mm, dd), nil
	} else {
		newUnique, err := getRowUnique(last.OrderID, 4, 4)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("F-%d%s%s%s", now.Year(), mm, dd, newUnique), nil
	}

}

func getRowUnique(lastOrderID string, lastUniqueLen int, lenUnique int) (string, error) {
	getUnique := lastOrderID[len(lastOrderID)-lastUniqueLen:]

	unique, err := strconv.Atoi(getUnique)
	if err != nil {
		return "", err
	}

	unique += 1

	newUnique := strconv.Itoa(unique)

	if len(newUnique) < lenUnique {
		for i := 0; i < lenUnique-len(newUnique); i++ {
			newUnique = "0" + newUnique
		}
	}

	return newUnique, nil
}
