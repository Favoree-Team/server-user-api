package utils

import (
	"strings"

	"github.com/google/uuid"
)

func NewUUID() string {
	// change uuid without dash
	uuidStr := uuid.New().String()
	uuidStr = strings.Replace(uuidStr, "-", "", -1)
	return uuidStr
}
