package types

import (
	"encoding/json"
	"meetspace_backend/client/models"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type UpdateUserData struct {
	FirstName string `form:"first_name" validate:"omitempty,max=50" json:"first_name"`
    LastName  string `form:"last_name" validate:"omitempty,max=50" json:"last_name"`
    ProfilePic *multipart.FileHeader `form:"profile_pic" validate:"omitempty" json:"profile_pic"`
	ClientID uuid.UUID  `json:"client_id,omitempty"`
	CreatedBy models.Client `json:"created_by,omitempty"`
    UpdatedBy  models.Client `json:"updated_by,omitempty"`
}

type CreateUserData struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Role string `json:"role,omitempty"`
	ClientID string `json:"client_id,omitempty"`
	CreatedBy *models.Client `json:"created_by,omitempty"`
    UpdatedBy  *models.Client `json:"updated_by,omitempty"`
}

type UserResponse struct {
	ID uuid.UUID `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	ProfilePic json.RawMessage `json:"profile_pic"` 
	IsActive bool `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
