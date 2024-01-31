package services

import (
	"errors"
	"meetspace_backend/auth/types"
	"meetspace_backend/user/models"
	"meetspace_backend/user/services"
	userTypes "meetspace_backend/user/types"
	"meetspace_backend/utils"

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

// user login
func (authService *AuthService) Login(reqData types.LoginRequest) (*types.AuthResponse, error) {
	// request data struct validation
	if err := utils.GetValidator().Struct(reqData); err != nil {
		return nil, err
    }

	// get user by email and compare the password against the user's password
	user, err := authService.UserService.GetUserByEmail(reqData.Email)
	if err != nil {
		return nil, err
	}
	
	isValid := utils.ComparePassword(user.Password, reqData.Password)
	if !isValid{
		return nil, errors.New("invalid password")
	}

	// generate token for user
	tokenData, _ := utils.GenerateUserToken(user.ID.String())
	resData := types.AuthResponse{
		User: user,
		Token: tokenData,
	}

	return &resData, err
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
