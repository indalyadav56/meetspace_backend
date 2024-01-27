package types

import "meetspace_backend/user/models"

type RegisterRequest struct {
	FistName string `json:"first_name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	Email string `json:"email" validate:"required,email"`
    Password  string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Email string `json:"email"`
    Password  string `json:"password"`
	UserType string `json:"user_type,omitempty"`
}

type AuthResponse struct {
	models.User
	Token map[string]interface{} `json:"token"`
}

type SendEmailRequest struct {
	Email string `json:"email"`
}

type VerifyEmailRequest struct {
	Email string `json:"email"`
	OTP string `json:"otp"`
}

type SendEmailResponse struct {
	Email string `json:"email"`
}