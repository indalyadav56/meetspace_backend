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

func NewChatGroupService() *ChatGroupService {
	repo := repositories.NewChatRoomRepository()
	newUserService := userService.NewUserService()
	
	return &ChatGroupService{
		ChatRoomRepository: repo,
		UserService: newUserService,
	}
}

func (crs *ChatGroupService) CreateChatGroup(chat models.ChatRoom) (models.ChatRoom, error) {
	return crs.ChatRoomRepository.CreateChatRoomRecord(chat)
}

