package repository

import (
	"github.com/Favoree-Team/server-user-api/entity"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetAll(offset int, limit int) (entity.ListTransaction, int64, error)
	GetByID(id string) (entity.Transaction, error)
	GetByUserID(userId string) (entity.ListTransaction, error)
	Insert(transaction entity.Transaction) error
	UpdateByID(id string, updates map[string]interface{}) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) GetAll(offset int, limit int) (entity.ListTransaction, int64, error) {
	var transactions []entity.Transaction

	if err := r.db.Limit(limit).Offset(offset).Order("created_at DESC").Find(&transactions).Error; err != nil {
		return nil, 0, err
	}

	var count int64
	if err := r.db.Model(&entity.Transaction{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return transactions, count, nil
}

func (r *transactionRepository) GetByID(id string) (entity.Transaction, error) {
	var transaction entity.Transaction

	if err := r.db.Where("id = ?", id).First(&transaction).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity.Transaction{}, nil
		} else {
			return entity.Transaction{}, err
		}
	}

	return transaction, nil
}

func (r *transactionRepository) GetByUserID(userId string) (entity.ListTransaction, error) {
	var transactions []entity.Transaction

	if err := r.db.Where("user_id = ?", userId).Order("created_at DESC").Find(&transactions).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return []entity.Transaction{}, nil
		} else {
			return nil, err
		}
	}

	return transactions, nil
}

func (r *transactionRepository) Insert(transaction entity.Transaction) error {
	if err := r.db.Create(&transaction).Error; err != nil {
		return err
	}

	return nil
}

func (r *transactionRepository) UpdateByID(id string, updates map[string]interface{}) error {
	if err := r.db.Model(&entity.Transaction{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return err
	}

	return nil
}
