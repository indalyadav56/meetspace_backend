package services

import (
	"meetspace_backend/auth/types"
	"meetspace_backend/user/models"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
}


func NewAuthService() *AuthService {
    return &AuthService{
    }
}


func (us *AuthService) Login(reqData types.LoginRequest) (models.User, error) {
	return models.User{}, nil
}


func (us *AuthService) Register(reqData types.RegisterRequest) (models.User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(reqData.Password), bcrypt.DefaultCost)

	user := models.User{
		FirstName: reqData.FistName,
		LastName: reqData.LastName,
		Email: reqData.Email,
		Password: string(hashedPassword),
	}
	return user, nil
    // return us.AuthRepository.Register(user)
}
