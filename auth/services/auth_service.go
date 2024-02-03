package services

import (
	"fmt"
	"meetspace_backend/auth/constants"
	"meetspace_backend/auth/types"
	gloablConstants "meetspace_backend/common/constants"
	commonServices "meetspace_backend/common/services"
	"meetspace_backend/user/services"
	userTypes "meetspace_backend/user/types"
	"meetspace_backend/utils"
	"strings"
	"time"
)

type AuthService struct {
	LoggerService *commonServices.LoggerService
	RedisService *commonServices.RedisService
	TokenService *TokenService
	UserService *services.UserService
}

func NewAuthService(
	loggerService *commonServices.LoggerService, 
	redisService *commonServices.RedisService,
	ts *TokenService, 
	us *services.UserService,
) *AuthService {

    return &AuthService{
		LoggerService: loggerService,
        RedisService:  redisService,
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

	// tokens store in redis
	as.RedisService.Set("user_access_token", "access_token", time.Minute * 15)

	// tokens store in redis
	accessTokenKey := fmt.Sprintf(gloablConstants.ACCESS_TOKEN_KEY, user.ID.String())
	refreshTokenKey := fmt.Sprintf(gloablConstants.REFRESH_TOKEN_KEY, user.ID.String())

	go as.RedisService.Set(accessTokenKey, tokens["access"], time.Minute * 15)
	go as.RedisService.Set(refreshTokenKey, tokens["refresh"], time.Hour * 48)
	
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

	// tokens store in redis
	accessTokenKey := fmt.Sprintf(gloablConstants.ACCESS_TOKEN_KEY, createdUser.ID.String())
	refreshTokenKey := fmt.Sprintf(gloablConstants.REFRESH_TOKEN_KEY, createdUser.ID.String())

	go as.RedisService.Set(accessTokenKey, tokens["access"], time.Minute * 15)
	go as.RedisService.Set(refreshTokenKey, tokens["refresh"], time.Hour * 48)
	
	data, _ := utils.StructToMap(reqData)
	data["token"] = tokens
	delete(data, "password")
	
    return utils.SuccessResponse("success", data)
}

// user logout
func (as *AuthService) UserLogout(userID string, reqData types.LogoutRequest) *utils.Response {
	// validate request struct data
	if err := utils.GetValidator().Struct(reqData); err != nil {
		data := utils.ParseError(err, reqData)
		return utils.ErrorResponse(constants.AUTH_REQUEST_VALIDATION_ERROR_MSG, data)
    }

	// verify refresh token
	_, err := as.TokenService.VerifyToken(reqData.RefreshToken, gloablConstants.REFRESH_TOKEN_KEY)
	if err != nil {
		return utils.ErrorResponse(err.Error(), nil)
	}

	// tokens store in redis
	accessTokenKey := fmt.Sprintf(gloablConstants.ACCESS_TOKEN_KEY, userID)
	refreshTokenKey := fmt.Sprintf(gloablConstants.REFRESH_TOKEN_KEY, userID)
	
	// remove tokens from redis
	go as.RedisService.Del(accessTokenKey)
	go as.RedisService.Del(refreshTokenKey)

    return utils.NoContentResponse()
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
	
	resp, err := as.UserService.UserRepository.UpdateUserByEmail(reqData.Email, updateData)
	if err != nil {
		return utils.ErrorResponse(err.Error(), nil)
	}

    return utils.SuccessResponse("success", resp)
}

// forgot password
func (as *AuthService) RefreshToken(refreshToken string) *utils.Response {
	if strings.TrimSpace(refreshToken) == "" {
		return utils.ErrorResponse("refresh token cannot be blank.", nil)
	}
	
	accessToken, err := as.TokenService.RefreshToken(refreshToken)
	if err != nil {
		return utils.ErrorResponse(err.Error(), nil)
	}
    return utils.SuccessResponse("success", map[string]string{"access": accessToken})
}
