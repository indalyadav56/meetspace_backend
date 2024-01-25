package types

import (
	"encoding/json"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type UpdateUserData struct {
	FirstName string `form:"first_name" validate:"omitempty,max=50" json:"first_name"`
    LastName  string `form:"last_name" validate:"omitempty,max=50" json:"last_name"`
    ProfilePic *multipart.FileHeader `form:"profile_pic" validate:"omitempty" json:"profile_pic"`
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
