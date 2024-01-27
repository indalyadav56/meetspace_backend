package repositories

import (
	"meetspace_backend/auth/models"
	"meetspace_backend/config"

	"gorm.io/gorm"
)

type VerificationRepository struct {
    db *gorm.DB
}

func NewVerificationRepository() *VerificationRepository {
	return &VerificationRepository{
		db:      config.DB,
	}
}

func (repo *VerificationRepository) CreateRecord(v models.Verification) (*models.Verification, error) {
	result := config.DB.Create(&v)
	if result.Error != nil {
        return nil, result.Error
    }
    return &v, nil
}

func (repo *VerificationRepository) GetRecordByEmail(email string) (models.Verification, error) {
	var model models.Verification
	
	result := config.DB.Where("email = ?", email).Order("updated_at DESC").First(&model)
	
	if result.Error != nil {
        return model, result.Error
    }
    return model, nil
}
