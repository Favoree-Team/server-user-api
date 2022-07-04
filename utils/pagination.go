package utils

func GetOffsite(page int, limit int) int {
	return (page - 1) * limit
}
