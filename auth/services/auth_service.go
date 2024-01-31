package services

import (
	"meetspace_backend/auth/constants"
	"meetspace_backend/auth/types"
	"meetspace_backend/user/services"
	userTypes "meetspace_backend/user/types"
	"meetspace_backend/utils"
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
func (authService *AuthService) Login(reqData types.LoginRequest) *utils.Response {
	// request data struct validation
	if err := utils.GetValidator().Struct(reqData); err != nil {
		data := utils.ParseError(err, reqData)
		return utils.ErrorResponse("error", data)
    }

	// get user by email and compare the password against the user's password
	user, err := authService.UserService.GetUserByEmail(reqData.Email)
	if err != nil {
		return utils.ErrorResponse(err.Error(), nil)
	}
	
	isValid := utils.ComparePassword(user.Password, reqData.Password)
	if !isValid{
		return utils.ErrorResponse("invalid password", nil)
	}

	// generate token for user
	tokenData, _ := utils.GenerateUserToken(user.ID.String())
	resData := types.AuthResponse{
		User: user,
		Token: tokenData,
	}
	return utils.SuccessResponse("success",resData)
}

// register new user
func (us *AuthService) Register(reqData types.RegisterRequest) *utils.Response {
	// validate request struct data
	if err := utils.GetValidator().Struct(reqData); err != nil {
		data := utils.ParseError(err, reqData)
		return utils.ErrorResponse(constants.AUTH_REQUEST_VALIDATION_ERROR_MSG, data)
    }

	hashedPassword, err := utils.EncryptPassword(reqData.Password)
	if err != nil {
		return utils.ErrorResponse(err.Error(), nil)
	}

	userCreateData := userTypes.CreateUserData{
		FirstName: reqData.FirstName,
		LastName: reqData.LastName,
		Email: reqData.Email,
		Password: string(hashedPassword),
	}
	createdUser, _ := us.UserService.CreateUser(userCreateData)
    return utils.SuccessResponse("success", createdUser)
}
