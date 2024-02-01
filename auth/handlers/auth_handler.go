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
// @Success      201 "Register user successfully"
// @Failure      400 "Bad request"
// @Failure      500 "Internal server error"
func (handler *AuthHandler) UserRegister(c *gin.Context){
	var req types.RegisterRequest
	
	if errResp := utils.BindJsonData(c, &req); errResp != nil {
		c.JSON(errResp.StatusCode, errResp)
        return
    }
	
	respData := handler.AuthService.Register(req)
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
	
	if errorResp := utils.BindJsonData(c, &req); errorResp != nil {
		c.JSON(errorResp.StatusCode, errorResp)
        return 
    }
	
	respData := handler.AuthService.Login(req)
	c.JSON(respData.StatusCode, respData)
	return 
}

// UserLogout godoc
//	@Summary		user-logout
//	@Description	User logout User
//	@Tags			Auth
//	@Produce		json
// 	@Param user body types.LoginRequest true "User login details"
//	@Router			/v1/auth/logout [post]
// @Success      200 "success"
// @Failure      400 "Bad request"
// @Failure      500 "Internal server error"
func (handler *AuthHandler) UserLogout(c *gin.Context) {
	var req types.LogoutRequest
	if err := utils.BindJsonData(c, &req); err != nil {
		resp:= utils.ErrorResponse(constants.REQUEST_BODY_ERROR_MSG, nil)
		c.JSON(resp.StatusCode, resp)
        return 
    }
	successResponse := utils.SuccessResponse("Success", "test")
	c.JSON(200, successResponse)
	return
}

// ForgotPassword godoc
//	@Summary		forgot-password
//	@Description	Forgot password
//	@Tags			Auth
//	@Produce		json
// 	@Param user body types.ForgotPasswordRequest true "forgot password request body"
//	@Router			/v1/auth/forgot-password [post]
// @Success      200 {object} utils.Response "Success"
// @Failure      400 "Bad request"
// @Failure      500 "Internal server error"
func (handler *AuthHandler) ForgotPassword(c *gin.Context){
	var req types.ForgotPasswordRequest
	if errResp := utils.BindJsonData(c, &req); errResp != nil {
		c.JSON(errResp.StatusCode, errResp)
        return 
    }
	
	successResponse := utils.SuccessResponse("Successfully logged in!!", "test")
	c.JSON(200, successResponse)
	return
}

// SendEmail godoc
// @Summary		send-email
// @Description	Send email to user
// @Tags			Auth
// @Produce		json
// @Param user body types.SendEmailRequest true "send email request body"
// @Router			/v1/auth/send-email [post]
// @Success      200 "Success"
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
//	@Summary		verify-email
//	@Description	Verify email otp.
//	@Tags			Auth
//	@Produce		json
// 	@Param user body types.VerifyEmailRequest true "verify email request body"
//	@Router			/v1/auth/verify-email [post]
// @Success      200 "Success"
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
