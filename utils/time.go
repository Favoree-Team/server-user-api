package utils

import (
	"strings"
	"time"
)

const (
	formatStr = time.RFC3339
)

func ParseStrtoTime(date string) (time.Time, error) {
	strToTime, err := time.Parse(formatStr, strings.Split(date, " m=")[0])
	if err != nil {
		return time.Time{}, err
	}

	return strToTime, nil
}
