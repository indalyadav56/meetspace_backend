package handlers

import (
	"meetspace_backend/auth/services"
	"meetspace_backend/auth/types"
	"meetspace_backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
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
//	@Param			user	body	types.RegisterRequest	true	"User registration details"
//	@Router			/v1/auth/register [post]
//	@Success		201	"Register user successfully"
//	@Failure		400	"Bad request"
//	@Failure		500	"Internal server error"
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
//	@Param			user	body	types.LoginRequest	true	"User login details"
//	@Router			/v1/auth/login [post]
//	@Success		200	"Login user successfully"
//	@Failure		400	"Bad request"
//	@Failure		500	"Internal server error"
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
//	@Param			user	body	types.LogoutRequest	true	"User login details"
//	@Router			/v1/auth/logout [post]
//	@Success		204	"success"
//	@Failure		400	"Bad request"
//	@Failure		500	"Internal server error"
//	@Security		Bearer
func (handler *AuthHandler) UserLogout(c *gin.Context) {
	// get user from context
	currentUser, _ := utils.GetUserFromContext(c)

	var req types.LogoutRequest
	if errorResp := utils.BindJsonData(c, &req); errorResp != nil {
		c.JSON(errorResp.StatusCode, errorResp)
        return 
    }

	resp := handler.AuthService.UserLogout(currentUser.ID.String(), req)
	c.JSON(resp.StatusCode, resp)
	return
}

// ForgotPassword godoc
//	@Summary		forgot-password
//	@Description	Forgot password
//	@Tags			Auth
//	@Produce		json
//	@Param			user	body	types.ForgotPasswordRequest	true	"forgot password request body"
//	@Router			/v1/auth/forgot-password [post]
//	@Success		200	{object}	utils.Response	"Success"
//	@Failure		400	"Bad request"
//	@Failure		500	"Internal server error"
func (handler *AuthHandler) ForgotPassword(c *gin.Context){
	var req types.ForgotPasswordRequest
	if errorResp := utils.BindJsonData(c, &req); errorResp != nil {
		c.JSON(errorResp.StatusCode, errorResp)
        return 
    }
	
	resp := handler.AuthService.ForgotPassword(req)
	c.JSON(resp.StatusCode, resp)
	return
}

// RefreshToken godoc
//	@Summary		refresh token
//	@Description	refresh jwt token
//	@Tags			Auth
//	@Produce		json
//	@Param			user	body	types.RefreshTokenRequest	true	"refresh token request body"
//	@Router			/v1/auth/refresh-token [post]
//	@Success		200	"Success"
//	@Failure		400	"Bad request"
//	@Failure		500	"Internal server error"
//	@Security		Bearer
func (handler *AuthHandler) RefreshToken(c *gin.Context) {
	var req types.RefreshTokenRequest
	if errResp := utils.BindJsonData(c, &req); errResp != nil {
		c.JSON(errResp.StatusCode, errResp)
        return 
    }
	
	resp := handler.AuthService.RefreshToken(req.RefreshToken)
	c.JSON(resp.StatusCode, resp)
	return 
}


// SendEmail godoc
//	@Summary		send-email
//	@Description	Send email to user
//	@Tags			Auth
//	@Produce		json
//	@Param			user	body	types.SendEmailRequest	true	"send email request body"
//	@Router			/v1/auth/send-email [post]
//	@Success		200	"Success"
//	@Failure		400	"Bad request"
//	@Failure		500	"Internal server error"
func (handler *AuthHandler) SendEmailHandler(c *gin.Context) {
	var req types.SendEmailRequest
	
	if errResp := utils.BindJsonData(c, &req); errResp != nil {
		c.JSON(errResp.StatusCode, errResp)
        return 
    }
	
	resp := handler.VerificationService.Create(req)
	c.JSON(resp.StatusCode, resp)
	return
}

// VerifyEmailHandler godoc
//	@Summary		verify-email
//	@Description	Verify email otp.
//	@Tags			Auth
//	@Produce		json
//	@Param			user	body	types.VerifyEmailRequest	true	"verify email request body"
//	@Router			/v1/auth/verify-email [post]
//	@Success		200	"Success"
//	@Failure		400	"Bad request"
//	@Failure		500	"Internal server error"
func (handler *AuthHandler) VerifyEmailHandler(c *gin.Context) {
	var req types.VerifyEmailRequest
	if errResp := utils.BindJsonData(c, &req); errResp != nil {
		c.JSON(errResp.StatusCode, errResp)
        return 
    }
	
	resp := handler.VerificationService.VerifyEmailOtp(req)
	c.JSON(resp.StatusCode, resp)
	return 
}

func (handler *AuthHandler) GoogleLogin(c *gin.Context) {
	q := c.Request.URL.Query()
	q.Add("provider", "google")
	c.Request.URL.RawQuery = q.Encode()
	gothic.BeginAuthHandler(c.Writer, c.Request)
	return
}

func (handler *AuthHandler) GitHubLogin(c *gin.Context) {
	q := c.Request.URL.Query()
	q.Add("provider", "github")
	c.Request.URL.RawQuery = q.Encode()
	gothic.BeginAuthHandler(c.Writer, c.Request)
	return
}

func (handler *AuthHandler) SSOCallback(c *gin.Context) {
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, user)
	return
}
