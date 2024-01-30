package handlers

import (
	"fmt"
	"meetspace_backend/auth/constants"
	"meetspace_backend/auth/models"
	"meetspace_backend/auth/types"
	"meetspace_backend/config"
	"meetspace_backend/utils"

	"github.com/gin-gonic/gin"
)

// 	UserRegister godoc
//	@Summary		Register User account
//	@Description	Register User account
//	@Tags			Auth
//	@Produce		json
// 	@Param user body types.RegisterRequest true "User registration details"
//	@Router			/v1/auth/register [post]
func UserRegister(c *gin.Context){
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

	user, err := config.AuthService.Register(req)
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
//	@Summary		UserLogin User account
//	@Description	UserLogin User account
//	@Tags			Auth
//	@Produce		json
//	@Router			/v1/auth/login [post]
func UserLogin(c *gin.Context) {
	var req types.LoginRequest
	if err := utils.BindJsonData(c, &req); err != nil {
        return 
    }

	user, err := config.AuthService.Login(req)
	if err != nil {
		return 
	}

	tokenData, _ := utils.GenerateUserToken(user.ID.String())
	fmt.Println(tokenData)
	resData := types.AuthResponse{
		// Token: tokenData,
	}
	successResponse := utils.SuccessResponse(constants.USER_LOGIN_MSG, resData)
	c.JSON(200, successResponse)
	return 
}

// UserLogout godoc
//	@Summary		UserLogout User account
//	@Description	UserLogout User account
//	@Tags			Auth
//	@Produce		json
//	@Router			/v1/auth/logout [post]
func UserLogout(c *gin.Context) {
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
func ForgotPassword(c *gin.Context){
	successResponse := utils.SuccessResponse("Successfully logged in!!", "test")
	c.JSON(200, successResponse)

	return
}

// SendEmail godoc
//	@Description	Send email to user
//	@Tags			Auth
//	@Produce		json
//	@Router			/v1/auth/send-email [post]
func SendEmailHandler(c *gin.Context) {
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
	config.VerificationService.Create(data)
	
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
func VerifyEmailHandler(c *gin.Context) {
	var reqBody types.VerifyEmailRequest
	err := c.ShouldBindJSON(&reqBody)
	if err != nil {
		return
	}

	modelObj, _ := config.VerificationService.GetVerificationDataByEmail(reqBody.Email)
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
