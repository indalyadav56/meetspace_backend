package services

import (
	"fmt"
	"meetspace_backend/auth/constants"
	"meetspace_backend/auth/types"
	"meetspace_backend/user/services"
	userTypes "meetspace_backend/user/types"
	"meetspace_backend/utils"
)

type AuthService struct {
	TokenService *TokenService
	UserService *services.UserService
}

func NewAuthService(ts *TokenService, us *services.UserService) *AuthService {
    return &AuthService{
		TokenService: ts,
		UserService: us,
    }
}

// user login
func (as *AuthService) Login(reqData types.LoginRequest) *utils.Response {
	// request data struct validation
	if err := utils.GetValidator().Struct(reqData); err != nil {
		data := utils.ParseError(err, reqData)
		return utils.ErrorResponse("error", data)
    }

	// get user by email and compare the password against the user's password
	user, err := as.UserService.GetUserByEmail(reqData.Email)
	if err != nil {
		return utils.ErrorResponse(err.Error(), nil)
	}
	
	isValid := utils.ComparePassword(user.Password, reqData.Password)
	if !isValid{
		return utils.ErrorResponse("invalid password", nil)
	}

	// generate new user tokens
	tokens, err := as.TokenService.GenerateToken(user.ID.String())
	
	data, _ := utils.StructToMap(reqData)
	data["token"] = tokens
	delete(data, "password")
	return utils.SuccessResponse("success", data)
}

// register new user
func (as *AuthService) Register(reqData types.RegisterRequest) *utils.Response {
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
	createdUser, err := as.UserService.CreateUser(userCreateData)
	if err != nil {
		var errData []utils.ErrorMsg
		errData = append(errData, utils.ErrorMsg{
				Field: "email",
				Message: err.Error(),
			})
		resp := utils.ErrorResponse(err.Error(), errData)
		return resp
	}
	
	// generate new user tokens
	tokens, err := as.TokenService.GenerateToken(createdUser.ID.String())
	
	data, _ := utils.StructToMap(reqData)
	data["token"] = tokens
	delete(data, "password")
	
    return utils.SuccessResponse("success", data)
}

// user logout
func (as *AuthService) UserLogout(reqData types.LogoutRequest) *utils.Response {
    return utils.SuccessResponse("success", nil)
}

// forgot password
func (as *AuthService) ForgotPassword(reqData types.ForgotPasswordRequest) *utils.Response {
	// validate request struct data
	if err := utils.GetValidator().Struct(reqData); err != nil {
		data := utils.ParseError(err, reqData)
		return utils.ErrorResponse(constants.AUTH_REQUEST_VALIDATION_ERROR_MSG, data)
    }

	hashedPassword, err := utils.EncryptPassword(reqData.NewPassword)
	if err != nil {
		return utils.ErrorResponse(err.Error(), nil)
	}

	updateData := map[string]interface{}{
		"Password": hashedPassword,
	}
	fmt.Println("updateData", updateData)
	resp, err := as.UserService.UserRepository.UpdateUserByEmail(reqData.Email, updateData)
	if err != nil {
		return utils.ErrorResponse(err.Error(), nil)
	}

    return utils.SuccessResponse("success", resp)
}
