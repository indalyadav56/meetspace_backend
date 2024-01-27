package models

import (
	"meetspace_backend/user/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type Verification struct {
    gorm.Model         `json:"-"`
    ID uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID   uuid.UUID `gorm:"type:uuid;foreignKey:UserID;references:ID;default:null" json:"-"`
	User *models.User `gorm:"foreignKey:UserID;references:ID;default:null" json:"-"`
    Email              string    `gorm:"default:null" json:"email"`
    PhoneNumber        string    `gorm:"default:null" json:"phone_number"`
    Otp                string    `json:"-"`
    VerificationType   string    `gorm:"default:EMAIL" json:"verification_type"`
	IsVerified bool `gorm:"default:false" json:"is_verified"`
    ExpiryAt           time.Time  `json:"expiry_at"`
}

const defaultExpiry = 1 * time.Minute

func (v *Verification) BeforeCreate(tx *gorm.DB) error {
    if v.ExpiryAt.IsZero() {
       v.ExpiryAt = time.Now().Add(defaultExpiry) 
    }
    return nil
}
