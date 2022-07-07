package repository

import (
	"github.com/Favoree-Team/server-user-api/entity"
	"gorm.io/gorm"
)

type IPRecordRepository interface {
	GetByIPAddress(ipAddress string) (entity.IPRecord, error)
	Create(records entity.IPRecord) error
}

type ipRecordRepository struct {
	db *gorm.DB
}

func NewIPRecordRepository(db *gorm.DB) *ipRecordRepository {
	return &ipRecordRepository{db: db}
}

func (r *ipRecordRepository) GetByIPAddress(ipAddress string) (entity.IPRecord, error) {
	var record entity.IPRecord
	if err := r.db.Where("ip_address = ?", ipAddress).First(&record).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity.IPRecord{}, nil
		} else {
			return entity.IPRecord{}, err
		}
	}
	return record, nil
}

func (r *ipRecordRepository) Create(records entity.IPRecord) error {
	return r.db.Create(&records).Error
}
