package services

import (
	"meetspace_backend/chat/models"
	"meetspace_backend/chat/repositories"
	userService "meetspace_backend/user/services"
)

type ChatGroupService struct {
	ChatRoomRepository *repositories.ChatRoomRepository
	UserService  *userService.UserService
}

func NewChatGroupService(repo *repositories.ChatRoomRepository, userService *userService.UserService) *ChatGroupService {
	return &ChatGroupService{
		ChatRoomRepository: repo,
		UserService: userService,
	}
}

func (crs *ChatGroupService) CreateChatGroup(chat models.ChatRoom) (models.ChatRoom, error) {
	return crs.ChatRoomRepository.CreateChatRoomRecord(chat)
}

