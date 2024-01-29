package types

import "meetspace_backend/user/models"

type RegisterRequest struct {
	FirstName string `json:"first_name" validate:"required,not_blank"`
	LastName string `json:"last_name" validate:"required,not_blank"`
	Email string `json:"email" validate:"required,email"`
    Password  string `json:"password" validate:"required,min=6,not_blank"`
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