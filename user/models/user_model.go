package models

import (
	"encoding/json"
	"meetspace_backend/client/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type User struct {
	gorm.Model  `json:"-"`
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `gorm:"uniqueIndex,index" json:"email"`
	Password string `gorm:"not null" json:"-"`
	ProfilePic json.RawMessage `gorm:"type:jsonb" json:"profile_pic"` 
	IsActive bool `json:"is_active" gorm:"default:true"`
	Language string `json:"language"`
	TimeZone string `json:"time_zone"`
	Theme string `json:"theme"`
	PhoneNumber string `json:"phone_number"`
	IsAdmin bool `gorm:"default:false" json:"is_admin"`
	Role string `gorm:"default:user" json:"role"`
	ClientID   uuid.UUID `gorm:"type:uuid;foreignKey:ClientID;references:ID;default:null;index" json:"-"`
	CreatedBy *models.Client `gorm:"foreignKey:ClientID;references:ID;default:null" json:"-"`
    UpdatedBy  *models.Client `gorm:"foreignKey:ClientID;references:ID;default:null" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
    var profile map[string]interface{}
    json.Unmarshal([]byte(u.ProfilePic), &profile)
	
    return
}