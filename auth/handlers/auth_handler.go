package handlers

import (
	"fmt"
	"meetspace_backend/auth/constants"
	"meetspace_backend/auth/models"
	"meetspace_backend/auth/services"
	"meetspace_backend/auth/types"
	"meetspace_backend/config"
	"meetspace_backend/utils"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct{
	AuthService *services.AuthService
	VerificationService *services.VerificationService
}

func NewAuthHandler(authSvc *services.AuthService, verificationSvc *services.VerificationService) *AuthHandler {
	return &AuthHandler{
		AuthService: authSvc,
		VerificationService: verificationSvc,
	}
}

// 	UserRegister godoc
//	@Summary		register-user
//	@Description	Register User account
//	@Tags			Auth
//	@Produce		json
// 	@Param user body types.RegisterRequest true "User registration details"
//	@Router			/v1/auth/register [post]
// @Success      200 "Register user successfully"
// @Failure      400 "Bad request"
// @Failure      500 "Internal server error"
func (handler *AuthHandler) UserRegister(c *gin.Context){
	var req types.RegisterRequest
	
	if err := utils.BindJsonData(c, &req); err != nil {
		resp:= utils.ErrorResponse(constants.REQUEST_BODY_ERROR_MSG, nil)
		c.JSON(resp.StatusCode, resp)
        return
    }
	
	if err := utils.GetValidator().Struct(req); err != nil {
		data := utils.ParseError(err, req)
		resp := utils.ErrorResponse(constants.AUTH_REQUEST_VALIDATION_ERROR_MSG, data)
		c.JSON(resp.StatusCode, resp)
		return
    }

	user, err := handler.AuthService.Register(req)
	if err != nil {
		var errData []utils.ErrorMsg
		errData = append(errData, utils.ErrorMsg{
			Field: "email",
			Message: err.Error(),
		})
		resp := utils.ErrorResponse(err.Error(), errData)
		c.JSON(resp.StatusCode, resp)
		return 
	}
	
	tokenData, _ := utils.GenerateUserToken(user.ID.String())
	resData := types.RegisterResponse{
		RegisterRequest: req,
		Token: tokenData,
	}
	respData := utils.SuccessResponse(constants.USER_REGISTER_MSG, resData)
	c.JSON(respData.StatusCode, respData)
	return
}

// 	UserLogin godoc
//	@Summary		login-user
//	@Description	Login user
//	@Tags			Auth
//	@Produce		json
// 	@Param user body types.LoginRequest true "User login details"
//	@Router			/v1/auth/login [post]
// @Success      200 "Login user successfully"
// @Failure      400 "Bad request"
// @Failure      500 "Internal server error"
func (handler *AuthHandler) UserLogin(c *gin.Context) {
	var req types.LoginRequest
	
	if err := utils.BindJsonData(c, &req); err != nil {
		resp:= utils.ErrorResponse(constants.REQUEST_BODY_ERROR_MSG, nil)
		c.JSON(resp.StatusCode, resp)
        return 
    }

	respData, err := handler.AuthService.Login(req)
	if err != nil {
		resp := utils.ErrorResponse(err.Error(), nil)
		c.JSON(resp.StatusCode, resp)
		return
	}

	successResponse := utils.SuccessResponse(constants.USER_LOGIN_MSG, respData)
	c.JSON(200, successResponse)
	return 
}

// UserLogout godoc
//	@Summary		UserLogout User account
//	@Description	UserLogout User account
//	@Tags			Auth
//	@Produce		json
//	@Router			/v1/auth/logout [post]
// @Success      200 "Login user successfully"
// @Failure      400 "Bad request"
// @Failure      500 "Internal server error"
func (handler *AuthHandler) UserLogout(c *gin.Context) {
	successResponse := utils.SuccessResponse("Successfully logged in!!", "test")
	c.JSON(200, successResponse)
	return
}

// ForgotPassword godoc
//	@Summary		ForgotPassword
//	@Description	ForgotPassword
//	@Tags			Auth
//	@Produce		json
//	@Router			/v1/auth/forgot-password [post]
// @Success      200 "Login user successfully"
// @Failure      400 "Bad request"
// @Failure      500 "Internal server error"
func (handler *AuthHandler) ForgotPassword(c *gin.Context){
	successResponse := utils.SuccessResponse("Successfully logged in!!", "test")
	c.JSON(200, successResponse)

	return
}

// SendEmail godoc
//	@Description	Send email to user
//	@Tags			Auth
//	@Produce		json
//	@Router			/v1/auth/send-email [post]
// @Success      200 "Login user successfully"
// @Failure      400 "Bad request"
// @Failure      500 "Internal server error"
func (handler *AuthHandler) SendEmailHandler(c *gin.Context) {
	var reqBody types.SendEmailRequest
	err := c.ShouldBindJSON(&reqBody)
	if err != nil {
		return
	}
	otp := utils.GenerateOTP()

	emailBody := fmt.Sprintf("Your OTP is:- <h2> %s </h2>", otp)
	go utils.SendEmail(reqBody.Email, "Email OTP", emailBody)
	
	data := models.Verification{
		Email: reqBody.Email,
		Otp: otp,
	}
	handler.VerificationService.Create(data)
	
	resp := utils.SuccessResponse("successfully send email!", gin.H{
	})
	c.JSON(200, resp)
	return
}

// VerifyEmailHandler godoc
//	@Description	Verify email otp.
//	@Tags			Auth
//	@Produce		json
//	@Router			/v1/auth/verify-email [post]
// @Success      200 "Login user successfully"
// @Failure      400 "Bad request"
// @Failure      500 "Internal server error"
func (handler *AuthHandler) VerifyEmailHandler(c *gin.Context) {
	var reqBody types.VerifyEmailRequest
	err := c.ShouldBindJSON(&reqBody)
	if err != nil {
		return
	}

	modelObj, _ := handler.VerificationService.GetVerificationDataByEmail(reqBody.Email)
	if modelObj.Otp == reqBody.OTP{
		config.DB.Model(&modelObj).Update("is_verified", true)
		resp := utils.SuccessResponse("successfully send email!", nil)
		c.JSON(200, resp)
		return
	}
	
	resp := utils.ErrorResponse("error", nil)
	c.JSON(200, resp)
	return 
}
