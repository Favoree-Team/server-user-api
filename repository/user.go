package repository

import (
	"github.com/Favoree-Team/server-user-api/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(email string) (entity.User, error)
	GetUserById(id string) (entity.User, error)
	GetUserByUsername(username string) (entity.User, error)
	Insert(user entity.User) error
	UpdateById(id string, updates map[string]interface{}) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByEmail(email string) (entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity.User{}, nil
		} else {
			return entity.User{}, err
		}
	}
	return user, nil
}

func (r *userRepository) GetUserById(id string) (entity.User, error) {
	var user entity.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity.User{}, nil
		} else {
			return entity.User{}, err
		}
	}
	return user, nil
}

func (r *userRepository) GetUserByUsername(username string) (entity.User, error) {
	var user entity.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity.User{}, nil
		} else {
			return entity.User{}, err
		}
	}
	return user, nil
}

func (r *userRepository) Insert(user entity.User) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) UpdateById(id string, updates map[string]interface{}) error {
	return r.db.Model(&entity.User{}).Where("id = ?", id).Updates(updates).Error
}
