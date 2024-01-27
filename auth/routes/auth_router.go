package routes

import (
	"meetspace_backend/auth/handlers"

	"github.com/gin-gonic/gin"
)


func AuthRouter(e *gin.Engine){
	authRouter := e.Group("/v1/auth")

	authRouter.POST("/register", handlers.UserRegister)
	authRouter.POST("/login", handlers.UserLogin)
	authRouter.POST("/logout", handlers.UserLogout)
	
	authRouter.POST("/send-email", handlers.SendEmailHandler)
	authRouter.POST("/verify-email", handlers.VerifyEmailHandler)

}