package entity

type User struct {
	ID              string
	FullName        string
	Email           string
	Username        string
	PhoneNumber     string
	Password        string
	Role            string
	Status          bool
	PictureUrl      string
	IsSubscribeBlog bool
	CreatedAt       string
	UpdatedAt       string
}

type UserRegisterInput struct {
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phone_number"`
}

type UserLoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}
