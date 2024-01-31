package services

import (
	"meetspace_backend/auth/models"
	"meetspace_backend/auth/repositories"
)

type VerificationService struct {
    VerificationRepository *repositories.VerificationRepository
}

func NewVerificationService(repo *repositories.VerificationRepository) *VerificationService {
    return &VerificationService{
        VerificationRepository: repo,
    }
}


func (v *VerificationService) Create(verification models.Verification) (*models.Verification, error) {
    return v.VerificationRepository.CreateRecord(verification)
}


func (v *VerificationService) GetVerificationDataByEmail(email string) (models.Verification, error) {
    return v.VerificationRepository.GetRecordByEmail(email)
}
