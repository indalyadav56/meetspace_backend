package types

import "meetspace_backend/user/models"

type RegisterRequest struct {
	FirstName string `json:"first_name" validate:"required,not_blank"`
	LastName string `json:"last_name" validate:"required,not_blank"`
	Email string `json:"email" validate:"required,email"`
    Password  string `json:"password" validate:"required,min=6,not_blank"`
}

type LoginRequest struct {
	Email string `json:"email" validate:"required,email"`
    Password  string `json:"password" validate:"required,min=6,not_blank"`
}

type RegisterResponse struct {
	RegisterRequest 
	Token map[string]string `json:"token"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
	NewPassword string `json:"new_password" validate:"required,not_blank"`
}

type LogoutRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required,not_blank"`
}

type AuthResponse struct {
	models.User
	Token map[string]string `json:"token"`
}

type SendEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type VerifyEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
	OTP string `json:"otp" validate:"required,not_blank"`
}

type SendEmailResponse struct {
	Email string `json:"email"`
}