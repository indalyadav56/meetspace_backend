package repositories

import (
	"meetspace_backend/chat/models"

	"gorm.io/gorm"
)

type ChatMessageRepository struct {
	db *gorm.DB
}

func NewChatMessageRepository(db *gorm.DB) *ChatMessageRepository {
	return &ChatMessageRepository{
		db:      db,
	}
}

func (crr *ChatMessageRepository) CreateRecord(chatMessage models.ChatMessage) (models.ChatMessage, error) {
	err := crr.db.Create(&chatMessage).Error
	if err != nil {
	    return models.ChatMessage{}, err
	}
	return chatMessage, nil
}
