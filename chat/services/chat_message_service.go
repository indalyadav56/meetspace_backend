package services

import (
	"meetspace_backend/chat/models"
	"meetspace_backend/chat/repositories"
	userModel "meetspace_backend/user/models"

	userService "meetspace_backend/user/services"
)

type ChatMessageService struct {
	ChatMessageRepository *repositories.ChatMessageRepository
	UserService  *userService.UserService
	ChatRoomService  *ChatRoomService
}

func NewChatMessageService() *ChatMessageService {
	repo := repositories.NewChatMessageRepository()
	newUserService := userService.NewUserService()
	chatRoomService := NewChatRoomService()
	
	return &ChatMessageService{
		ChatMessageRepository: repo,
		UserService: newUserService,
		ChatRoomService: chatRoomService,
	}
}

func (chatMessageService *ChatMessageService) CreateChatMessage(content string, senderId string, chatRoomId string) (models.ChatMessage, error) {
	sender := chatMessageService.UserService.GetUserByID(senderId)
	chatRoom, _ := chatMessageService.ChatRoomService.GetChatRoomByID(chatRoomId)

	userResponse := sender.Data.(userModel.User)
	
	chatMessage := models.ChatMessage{
		Content: content,
		Sender: userResponse,
		ChatRoom: chatRoom,
	}

	return chatMessageService.ChatMessageRepository.CreateRecord(chatMessage)
}