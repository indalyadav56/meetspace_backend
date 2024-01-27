package services

import (
	"fmt"
	"meetspace_backend/auth/repositories"
	"meetspace_backend/auth/types"
	"meetspace_backend/user/models"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
    AuthRepository *repositories.AuthRepository
}


func NewAuthService() *AuthService {
	fmt.Println("initializing new auth service")
	authRepo := repositories.NewAuthRepository()
    return &AuthService{
        AuthRepository: authRepo,
    }
}


func (us *AuthService) Login(reqData types.LoginRequest) (models.User, error) {
    return us.AuthRepository.Login(reqData.Email, reqData.Password)
}


func (us *AuthService) Register(reqData types.RegisterRequest) (models.User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(reqData.Password), bcrypt.DefaultCost)

	user := models.User{
		FirstName: reqData.FistName,
		LastName: reqData.LastName,
		Email: reqData.Email,
		Password: string(hashedPassword),
	}

    return us.AuthRepository.Register(user)
}
