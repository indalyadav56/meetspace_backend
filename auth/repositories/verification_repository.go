package repositories

import (
	"meetspace_backend/auth/models"

	"gorm.io/gorm"
)

type VerificationRepository struct {
    db *gorm.DB
}

func NewVerificationRepository(db *gorm.DB) *VerificationRepository {
	return &VerificationRepository{
		db:      db,
	}
}

func (repo *VerificationRepository) CreateRecord(v models.Verification) (*models.Verification, error) {
	result := repo.db.Create(&v)
	if result.Error != nil {
        return nil, result.Error
    }
    return &v, nil
}

func (repo *VerificationRepository) GetRecordByEmailAndOtp(email, otp string,) (models.Verification, error) {
	var model models.Verification
	
	result := repo.db.Where("email = ? AND otp = ?", email, otp).Order("updated_at DESC").First(&model)
	
	if result.Error != nil {
        return model, result.Error
    }
    return model, nil
}
