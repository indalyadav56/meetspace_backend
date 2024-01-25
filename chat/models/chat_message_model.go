package models

import (
	"meetspace_backend/user/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ChatMessage struct {
	gorm.Model `json:"-"`
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Content     string `json:"content"`
	SenderID uuid.UUID `gorm:"foreignKey:SenderID;references:ID" json:"sender_id"`
	Sender   models.User `gorm:"joinForeignKey:SenderID" json:"sender"`
	ChatRoomID uuid.UUID `gorm:"foreignKey:ChatRoomID;references:ID"`
	ChatRoom ChatRoom `gorm:"joinForeignKey:ChatRoomID" json:"chat_room"`
	IsSeen bool `gorm:"default:false" json:"is_seen"`
	SeenBy []*models.User `gorm:"many2many:SeenBy;" json:"seen_by"`
	SeenAt time.Time `json:"seen_at"`
	IsDeleted bool `gorm:"default:false" json:"is_deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
