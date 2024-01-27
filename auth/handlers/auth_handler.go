package handlers

import (
	"errors"
	"fmt"
	"meetspace_backend/auth/constants"
	"meetspace_backend/auth/models"
	"meetspace_backend/auth/types"
	"meetspace_backend/config"
	"meetspace_backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
    Field string `json:"field"`
    Message   string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
    switch fe.Tag() {
        case "required":
            return "This field is required"
        case "lte":
            return "Should be less than " + fe.Param()
        case "gte":
            return "Should be greater than " + fe.Param()
    }
    return "Unknown error"
}

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
		resp:= utils.ErrorResponse("Invalid JSON", err.Error())
		c.JSON(resp.StatusCode, resp)
        return 
    }

	validate := validator.New()
    if err := validate.Struct(req); err != nil {
		var ve validator.ValidationErrors
        if errors.As(err, &ve) {
            out := make([]ErrorMsg, len(ve))
            for i, fe := range ve {
                out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
            }
			resp := utils.ErrorResponse("Invalid Data", out)
			c.JSON(http.StatusBadRequest, resp)
			return
        }
		
    }

	user, err := config.AuthService.Register(req)
	if err != nil {
		return 
	}

	accessToken, refreshToken, _ := utils.GenerateTokenPair(user.ID.String())

	tokenData := map[string]interface{}{
		"access": accessToken,
		"refresh": refreshToken,
	}

	resData := types.AuthResponse{
		User: user,
		Token: tokenData,
	}
	successResponse := utils.SuccessResponse("Successfully logged in!!", resData)
	c.JSON(200, successResponse)
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
	resData := types.AuthResponse{
		User: user,
		Token: tokenData,
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
