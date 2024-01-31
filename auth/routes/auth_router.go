package routes

import (
	"meetspace_backend/auth/handlers"

	"github.com/gin-gonic/gin"
)


func AuthRouter(e *gin.Engine, handler *handlers.AuthHandler){
	authRouter := e.Group("/v1/auth")

	authRouter.POST("/register", handler.UserRegister)
	authRouter.POST("/login", handler.UserLogin)
	authRouter.POST("/logout", handler.UserLogout)
	
	authRouter.POST("/send-email", handler.SendEmailHandler)
	authRouter.POST("/verify-email", handler.VerifyEmailHandler)

}