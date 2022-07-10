package entity

type ClaimData struct {
	UserId          string `json:"user_id"`
	Email           string `json:"email"`
	Role            string `json:"role"`
	IsSubscribeBlog bool   `json:"is_subscribe_blog"`
	Active          bool   `json:"active"`
}
