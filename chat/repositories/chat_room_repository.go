package repositories

import (
	"meetspace_backend/chat/models"

	"gorm.io/gorm"
)


type ChatRoomRepository struct {
	db *gorm.DB
}

func NewChatRoomRepository(db *gorm.DB) *ChatRoomRepository {
	return &ChatRoomRepository{
		db:      db,
	}
}

func (crr *ChatRoomRepository) CreateChatRoomRecord(chatRoom models.ChatRoom) (models.ChatRoom, error) {
	err := crr.db.Create(&chatRoom).Error
	if err != nil {
	    return models.ChatRoom{}, err
	}
	return chatRoom, nil
}

func (crr *ChatRoomRepository) CreateRecord(chatRoom models.ChatRoom) (models.ChatRoom, error) {
	err := crr.db.Create(&chatRoom).Error
	if err != nil {
	    return models.ChatRoom{}, err
	}
	return chatRoom, nil
}

func (crr *ChatRoomRepository) GetChatRoomByID(roomID string) (models.ChatRoom, error) {
	var chatRoom models.ChatRoom
	err := crr.db.Where("id = ?", roomID).First(&chatRoom).Error
	if err != nil {
	    return models.ChatRoom{}, err
	}
	return chatRoom, nil
}

func (crr *ChatRoomRepository) DeleteChatRoomRecord() (models.ChatRoom, error) {
	var chatRoom models.ChatRoom

	err := crr.db.Where("id = ?", "roomID").First(&chatRoom).Error
	
	if err != nil {
	    return models.ChatRoom{}, err
	}
	return chatRoom, nil
}