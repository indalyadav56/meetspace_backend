package services

import (
	"meetspace_backend/auth/types"
	"meetspace_backend/user/models"
	"meetspace_backend/user/services"
	userTypes "meetspace_backend/user/types"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserService *services.UserService
}


func NewAuthService(service *services.UserService) *AuthService {
    return &AuthService{
		UserService: service,
    }
}


func (us *AuthService) Login(reqData types.LoginRequest) (models.User, error) {
	return models.User{}, nil
}

// register new user
func (us *AuthService) Register(reqData types.RegisterRequest) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqData.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	userCreateData := userTypes.CreateUserData{
		FirstName: reqData.FirstName,
		LastName: reqData.LastName,
		Email: reqData.Email,
		Password: string(hashedPassword),
	}
	createdUser, err := us.UserService.CreateUser(userCreateData)
    return createdUser, err
}
