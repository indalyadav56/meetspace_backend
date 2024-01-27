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
	"github.com/gofiber/fiber/v2"
)

var authService = services.NewAuthService()
var verificationService = services.NewVerificationService()

// 	UserRegister godoc
//	@Summary		Register User account
//	@Description	Register User account
//	@Tags			Auth
//	@Produce		json
// 	@Param user body types.RegisterRequest true "User registration details"
//	@Router			/v1/auth/register [post]
func UserRegister(c *fiber.Ctx) error{
	var req types.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
        return err
    }
	user, err := authService.Register(req)
	if err != nil {
		return err
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
	c.Status(successResponse.StatusCode).JSON(successResponse)
	return nil
}

// 	UserLogin godoc
//	@Summary		UserLogin User account
//	@Description	UserLogin User account
//	@Tags			Auth
//	@Produce		json
//	@Router			/v1/auth/login [post]
func UserLogin(c *fiber.Ctx) error {
	var req types.LoginRequest
	if err := c.BodyParser(&req); err != nil {
        return err
    }

	user, err := authService.Login(req)
	if err != nil {
		return err
	}

	tokenData, _ := utils.GenerateUserToken(user.ID.String())
	resData := types.AuthResponse{
		User: user,
		Token: tokenData,
	}
	successResponse := utils.SuccessResponse(constants.USER_LOGIN_MSG, resData)
	c.Status(successResponse.StatusCode).JSON(successResponse)
	return nil
}

// UserLogout godoc
//	@Summary		UserLogout User account
//	@Description	UserLogout User account
//	@Tags			Auth
//	@Produce		json
//	@Router			/v1/auth/logout [post]
func UserLogout(c *fiber.Ctx) error {
	successResponse := utils.SuccessResponse("Successfully logged in!!", "test")
	c.Status(successResponse.StatusCode).JSON(successResponse)
	return nil
}

// ForgotPassword godoc
//	@Summary		ForgotPassword
//	@Description	ForgotPassword
//	@Tags			Auth
//	@Produce		json
//	@Router			/v1/auth/forgot-password [post]
func ForgotPassword(c *fiber.Ctx) error{
	successResponse := utils.SuccessResponse("Successfully logged in!!", "test")
	c.Status(200).Status(successResponse.StatusCode).JSON(successResponse)
	return nil
}

// SendEmail godoc
//	@Description	Send email to user
//	@Tags			Auth
//	@Produce		json
//	@Router			/v1/auth/send-email [post]
func SendEmailHandler(c *fiber.Ctx) error {
	var reqBody types.SendEmailRequest
	err := c.BodyParser(&reqBody)
	if err != nil {
		return nil
	}
	otp := utils.GenerateOTP()

	emailBody := fmt.Sprintf("Your OTP is:- <h2> %s </h2>", otp)
	go utils.SendEmail(reqBody.Email, "Email OTP", emailBody)
	
	data := models.Verification{
		Email: reqBody.Email,
		Otp: otp,
	}
	verificationService.Create(data)
	
	resp := utils.SuccessResponse("successfully send email!", gin.H{
	})
	c.JSON(resp)
	return nil
}

// VerifyEmailHandler godoc
//	@Description	Verify email otp.
//	@Tags			Auth
//	@Produce		json
//	@Router			/v1/auth/verify-email [post]
func VerifyEmailHandler(c *fiber.Ctx)  error{
	var reqBody types.VerifyEmailRequest
	err := c.BodyParser(&reqBody)
	if err != nil {
		return nil
	}

	modelObj, _ := verificationService.GetVerificationDataByEmail(reqBody.Email)
	if modelObj.Otp == reqBody.OTP{
		config.DB.Model(&modelObj).Update("is_verified", true)
		resp := utils.SuccessResponse("successfully send email!", nil)
		c.Status(resp.StatusCode).JSON(resp)
		return nil
	}
	
	resp := utils.ErrorResponse("error", nil)
	c.JSON(resp)
	return nil
}
