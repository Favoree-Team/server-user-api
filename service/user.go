package service

import (
	"errors"
	"net/http"
	"time"

	"github.com/Favoree-Team/server-user-api/auth"
	"github.com/Favoree-Team/server-user-api/entity"
	"github.com/Favoree-Team/server-user-api/notification"
	"github.com/Favoree-Team/server-user-api/repository"
	"github.com/Favoree-Team/server-user-api/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(input entity.UserRegisterInput) (entity.UserResponse, error)
	LoginUser(input entity.UserLoginInput) (entity.UserResponse, error)
	SubscribeBlog(id string) error

	GetUserID(id string) (entity.UserDetailResponse, error)
	UserEditbyID(id string, input entity.UserProfileEdit) error
	// Verified()
	// InactiveUser(id string) error
	// ActivateUser(id string) error

	// ip record
	CheckIPRecord(input entity.IPRecordRequest) (entity.IPRecordResponse, error)
	InsertIPRecord(input entity.IPRecordInput) error
}

type userService struct {
	userRepository repository.UserRepository
	ipRecordRepo   repository.IPRecordRepository
	authService    auth.AuthService
	emailNotif     notification.EmailNotification
}

func NewUserService(userRepository repository.UserRepository, ipRecordRepo repository.IPRecordRepository, authService auth.AuthService, emailNotif notification.EmailNotification) *userService {
	return &userService{
		userRepository: userRepository,
		ipRecordRepo:   ipRecordRepo,
		authService:    authService,
		emailNotif:     emailNotif,
	}
}

// this case in for active and inactive user
func (s *userService) RegisterUser(input entity.UserRegisterInput) (entity.UserResponse, error) {
	// get user and check email in database
	userResp := entity.UserResponse{}

	checkEmail, err := s.userRepository.GetUserByEmail(input.Email)

	if err != nil {
		return userResp, utils.CreateErrorMsg(http.StatusInternalServerError, err)
	}

	// if found, return error information
	if checkEmail.ID != "" && checkEmail.Email == input.Email {
		return userResp, utils.CreateErrorMsg(http.StatusBadRequest, errors.New("email already exists"))
	} else if input.Username != "" && checkEmail.Username == input.Username {
		return userResp, utils.CreateErrorMsg(http.StatusBadRequest, errors.New("username already exists"))
	}

	if input.Username != "" {
		checkUname, err := s.userRepository.GetUserByUsername(input.Username)
		if err != nil {
			return userResp, utils.CreateErrorMsg(http.StatusInternalServerError, err)
		}

		if checkUname.ID != "" && checkUname.Username == input.Username {
			return userResp, utils.CreateErrorMsg(http.StatusBadRequest, errors.New("username already exists"))
		}
	}

	// generate password to hash
	hashPass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return userResp, utils.CreateErrorMsg(http.StatusInternalServerError, err)
	}

	// if not found, create new user
	user := entity.User{
		ID:              utils.NewUUID(),
		Email:           input.Email,
		Username:        input.Username,
		PhoneNumber:     input.PhoneNumber,
		Password:        string(hashPass),
		Role:            "user",
		Status:          true,
		IsSubscribeBlog: false,
	}

	// insert user to database
	err = s.userRepository.Insert(user)
	if err != nil {
		return userResp, utils.CreateErrorMsg(http.StatusInternalServerError, err)
	}

	// create token
	token, err := s.authService.GenerateToken(entity.ClaimData{
		UserId:          user.ID,
		Role:            user.Role,
		IsSubscribeBlog: user.IsSubscribeBlog,
		Active:          user.Status,
		Email:           user.Email,
	})
	if err != nil {
		return userResp, utils.CreateErrorMsg(http.StatusInternalServerError, err)
	}

	//TODO: send email verification to user email with goroutine

	// return userResponse
	userResp.ID = user.ID
	userResp.Email = user.Email
	userResp.FullName = user.FullName
	userResp.Role = user.Role
	userResp.Username = user.Username
	userResp.Token = token

	return userResp, nil
}

// this case success for user active only
func (s *userService) LoginUser(input entity.UserLoginInput) (entity.UserResponse, error) {
	// get user and check email in database
	userResp := entity.UserResponse{}

	user, err := s.userRepository.GetUserByEmail(input.Email)
	if err != nil {
		return userResp, utils.CreateErrorMsg(http.StatusInternalServerError, err)
	}

	// if not found, return error information
	if user.ID == "" || user.Email != input.Email {
		return userResp, utils.CreateErrorMsg(http.StatusBadRequest, errors.New("email or password is incorrect"))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return userResp, utils.CreateErrorMsg(http.StatusBadRequest, errors.New("email or password is incorrect"))
	}

	if !user.Status {
		return userResp, utils.CreateErrorMsg(http.StatusBadRequest, errors.New("user is inactive"))
	}

	// if found, create token
	token, err := s.authService.GenerateToken(entity.ClaimData{
		UserId:          user.ID,
		Role:            user.Role,
		IsSubscribeBlog: user.IsSubscribeBlog,
		Active:          user.Status,
		Email:           user.Email,
	})
	if err != nil {
		return userResp, utils.CreateErrorMsg(http.StatusInternalServerError, err)
	}

	// send data
	userResp.ID = user.ID
	userResp.Email = user.Email
	userResp.FullName = user.FullName
	userResp.Role = user.Role
	userResp.Username = user.Username
	userResp.Token = token

	return userResp, nil
}

func (s *userService) SubscribeBlog(id string) error {
	// get user and check email in database

	user, err := s.userRepository.GetUserById(id)
	if err != nil {
		return utils.CreateErrorMsg(http.StatusInternalServerError, err)
	}

	if !user.IsSubscribeBlog {
		updates := map[string]interface{}{
			"is_subscribe_blog": true,
			"updated_at":        time.Now(),
		}

		err = s.userRepository.UpdateById(id, updates)
		if err != nil {
			return utils.CreateErrorMsg(http.StatusInternalServerError, err)
		}
	}

	return nil
}

func (s *userService) GetUserID(id string) (entity.UserDetailResponse, error) {
	user, err := s.userRepository.GetUserById(id)
	if err != nil {
		return entity.UserDetailResponse{}, utils.CreateErrorMsg(http.StatusInternalServerError, err)
	}

	return user.ToUserDetailResponse(), nil
}

func (s *userService) UserEditbyID(id string, input entity.UserProfileEdit) error {
	var edit = map[string]interface{}{}

	edit["updated_at"] = time.Now()

	if input.FullName != "" || len(input.FullName) > 0 {
		edit["full_name"] = input.FullName
	}

	if input.PictureUrl != "" || len(input.PictureUrl) > 0 {
		edit["picture_url"] = input.PictureUrl
	}

	if input.Username != "" || len(input.Username) > 0 {
		edit["username"] = input.Username
	}

	if input.PhoneNumber != "" || len(input.PhoneNumber) > 0 {
		edit["phone_number"] = input.PhoneNumber
	}

	err := s.userRepository.UpdateById(id, edit)
	if err != nil {
		return utils.CreateErrorMsg(http.StatusInternalServerError, err)
	}

	return nil
}

func (s *userService) CheckIPRecord(input entity.IPRecordRequest) (entity.IPRecordResponse, error) {
	ipRecord, err := s.ipRecordRepo.GetByIPAddress(input.IPAddress)
	if err != nil {
		return entity.IPRecordResponse{}, utils.CreateErrorMsg(http.StatusInternalServerError, err)
	}

	if ipRecord.ID == "" || ipRecord.IPAddress != input.IPAddress {
		return entity.IPRecordResponse{
			IPAddress: input.IPAddress,
			Status:    entity.RejectedStatus,
			Message:   "IP address not found",
		}, nil
	}

	if ipRecord.IPAddress == input.IPAddress {
		if ipRecord.Role == "admin" {
			return entity.IPRecordResponse{
				IPAddress: input.IPAddress,
				Status:    entity.AccessibleStatus,
				Message:   "IP address already accepted",
			}, nil
		}
	}

	return entity.IPRecordResponse{
		IPAddress: input.IPAddress,
		Status:    entity.RejectedStatus,
		Message:   "IP address rejected",
	}, nil
}

func (s *userService) InsertIPRecord(input entity.IPRecordInput) error {
	var ipRecord = entity.IPRecord{
		ID:        utils.NewUUID(),
		IPAddress: input.IPAddress,
		Role:      input.Role,
	}

	return s.ipRecordRepo.Create(ipRecord)
}
