package repositories

import (
	"errors"
	authModel "meetspace_backend/auth/models"
	"meetspace_backend/config"
	userConstant "meetspace_backend/user/constants"
	"meetspace_backend/user/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository struct {
    db *gorm.DB
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{
		db:      config.DB,
	}
}

func (userRepo *AuthRepository) Login(email string, password string) (models.User, error) {
    var user models.User
	roles := []string{userConstant.ROLE_ADMIN, userConstant.ROLE_USER, userConstant.ROLE_SUPER_ADMIN}
	
	result := config.DB.Where("email = ? AND role IN (?)", email, roles).First(&user)
	if result.Error != nil {
        return user, result.Error
    }
	
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("invalid credentials")
	}
	
    return user, nil
}

func (userRepo *AuthRepository) Register(user models.User) (models.User, error) {
    config.DB.Create(&user)
    return user, nil
}

func (userRepo *AuthRepository) CreateVerificationRecord(data authModel.Verification) (*authModel.Verification, error) {
	config.DB.Create(&data)
    return &data, nil
}
