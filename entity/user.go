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

func (u *User) ToUserDetailResponse() UserDetailResponse {
	return UserDetailResponse{
		ID:              u.ID,
		FullName:        u.FullName,
		Email:           u.Email,
		Username:        u.Username,
		PhoneNumber:     u.PhoneNumber,
		Role:            u.Role,
		Status:          u.Status,
		PictureUrl:      u.PictureUrl,
		IsSubscribeBlog: u.IsSubscribeBlog,
		CreatedAt:       u.CreatedAt,
		UpdatedAt:       u.UpdatedAt,
	}
}

type UserDetailResponse struct {
	ID              string `json:"id"`
	FullName        string `json:"full_name"`
	Email           string `json:"email"`
	Username        string `json:"username"`
	PhoneNumber     string `json:"phone_number"`
	Role            string `json:"role"`
	Status          bool   `json:"status"`
	PictureUrl      string `json:"picture_url"`
	IsSubscribeBlog bool   `json:"is_subscribe_blog"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

// for verification code
type VerifyEmailRequest struct {
	VerificationCode string `json:"verification_code"`
}

// for reset password
type ResetPasswordInput struct {
	NewPassword string `json:"new_password"`
}

type UserProfileEdit struct {
	FullName    string `json:"full_name"`
	PictureUrl  string `json:"picture_url"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phone_number"`
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
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}
