package routes

import (
	"meetspace_backend/auth/handlers"

	"github.com/gofiber/fiber/v2"
)


func AuthRouter(app *fiber.App){
	authRouter := app.Group("/v1/auth")

	authRouter.Post("/register", handlers.UserRegister)
	authRouter.Post("/login", handlers.UserLogin)
	authRouter.Post("/logout", handlers.UserLogout)
	authRouter.Post("/forgot-password", handlers.ForgotPassword)
	
	authRouter.Post("/send-email", handlers.SendEmailHandler)
	authRouter.Post("/verify-email", handlers.VerifyEmailHandler)
}