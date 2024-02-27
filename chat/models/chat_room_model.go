package models

import (
	"meetspace_backend/user/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ChatRoom struct {
	gorm.Model `json:"-"`
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"  json:"id"`
	RoomName string `json:"room_name"`
	IsGroup bool `grom:"default:false;" json:"is_group"`
	RoomOwnerID uuid.UUID `gorm:"foreignKey:RoomOwnerID;references:ID"`
	RoomOwner *models.User `gorm:"joinForeignKey:RoomOwnerID" json:"room_owner"`
	RoomUsers []*models.User `gorm:"many2many:RoomUsers;" json:"room_users"`
	IsDeleted bool `json:"is_deleted" gorm:"default:false;"`
}

// func (r *ChatRoom) AfterSave(tx *gorm.DB) error {
// 	if tx.Statement.Changed("UpdatedAt") {
// 		tx.Model(ChatMessage{}).Where("chat_room_id = ?", r.ID).UpdateColumn("updated_at", time.Now())   
// 	}
// 	return nil
// }