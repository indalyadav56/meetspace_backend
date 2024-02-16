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

func (crr *ChatRoomRepository) CreateRecord(chatRoom models.ChatRoom) (models.ChatRoom, error) {
	err := crr.db.Create(&chatRoom).Error
	if err != nil {
	    return models.ChatRoom{}, err
	}
	return chatRoom, nil
}

func (crr *ChatRoomRepository) UpdateRecord(chatRoom models.ChatRoom) (models.ChatRoom, error) {
	err := crr.db.Model(&chatRoom).Where("id = ?", chatRoom.ID.String()).Updates(&chatRoom).Error
	if err != nil {
	    return models.ChatRoom{}, err
	}
	return chatRoom, nil
}

func (r *ChatRoomRepository) GetChatRoomByID(roomID string) (models.ChatRoom, error) {
	var chatRoom models.ChatRoom
	err := r.db.Where("id = ?", roomID).First(&chatRoom).Error
	if err != nil {
	    return models.ChatRoom{}, err
	}
	return chatRoom, nil
}

func (crr *ChatRoomRepository) DeleteChatRoomRecord(roomID string) (bool, error) {
	err := crr.db.Model(&models.ChatRoom{}).Where("id = ?", roomID).Update("is_deleted", true).Error
	
	if err != nil {
	    return false, err
	}
	return true, nil
}