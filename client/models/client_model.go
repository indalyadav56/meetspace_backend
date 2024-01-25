package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type Client struct {
	gorm.Model  `json:"-"`
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	CompanyName string `gorm:"uniqueIndex;index" json:"company_name"`
	CompanyDomain string `gorm:"default:1" json:"company_domain"`
	CompanySize uint32 `gorm:"default:1" json:"company_size"`
	Country string `gorm:"default:india" json:"country"`
	ClientUserID string `json:"client_user_id"`
}