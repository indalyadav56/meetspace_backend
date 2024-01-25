package repositories

import (
	"meetspace_backend/chat/models"
	"meetspace_backend/config"

	"gorm.io/gorm"
)

type ChatMessageRepository struct {
	db *gorm.DB
}

func NewChatMessageRepository() *ChatMessageRepository {
	return &ChatMessageRepository{
		db:      config.DB,
	}
}

func (crr *ChatMessageRepository) CreateRecord(chatMessage models.ChatMessage) (models.ChatMessage, error) {
	err := crr.db.Create(&chatMessage).Error
	if err != nil {
	    return models.ChatMessage{}, err
	}
	return chatMessage, nil
}
