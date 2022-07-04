package utils

import (
	"fmt"
	"strconv"
	"strings"
)

type ErrorMessage error

var (
	errorErr = map[int]string{
		400: "Bad Request",
		401: "Unauthorized",
		404: "Not Found",
		500: "Internal Server Error",
	}
)

type errMessageModel struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

func FindErr(code int) string {
	return errorErr[code]
}

func ErrorHandler(err ErrorMessage) *errMessageModel {
	msg := strings.Split(err.Error(), "::")

	errCode, _ := strconv.Atoi(msg[0])

	return &errMessageModel{
		Status:  errCode,
		Error:   FindErr(errCode),
		Message: msg[1],
	}
}

func GetErrorCode(err ErrorMessage) int {
	msg := strings.Split(err.Error(), "::")

	errCode, _ := strconv.Atoi(msg[0])

	return errCode
}

func CreateErrorMsg(code int, message error) ErrorMessage {
	return fmt.Errorf("%d::%s", code, message)
}
