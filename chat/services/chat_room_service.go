package services

import (
	"meetspace_backend/chat/models"
	"meetspace_backend/chat/repositories"
	userModel "meetspace_backend/user/models"
	userService "meetspace_backend/user/services"
)

type ChatRoomService struct {
	ChatRoomRepository *repositories.ChatRoomRepository
	UserService  *userService.UserService
}

func NewChatRoomService(repo *repositories.ChatRoomRepository, userService *userService.UserService) *ChatRoomService {
	return &ChatRoomService{
		ChatRoomRepository: repo,
		UserService: userService,
	}
}

func (crs *ChatRoomService) CreateChatRoomRecord(roomName string, roomOwnerId string, roomUserIds []string) (models.ChatRoom, error) {
	roomOnwer, _ := crs.UserService.UserRepository.GetUserByID(roomOwnerId)
	var roomUsers []*userModel.User
	
	for _, userId := range roomUserIds {
		user, _ := crs.UserService.UserRepository.GetUserByID(userId)
		roomUsers = append(roomUsers, user)
	}

	roomUsers = append(roomUsers, roomOnwer)

	chatRoom := models.ChatRoom{
		RoomName: roomName,
		RoomOwner: roomOnwer,
		RoomUsers: roomUsers,
	}
	return crs.ChatRoomRepository.CreateChatRoomRecord(chatRoom)
}

func (crs *ChatRoomService) GetChatRoomByID(roomID string) (models.ChatRoom, error) {
	return crs.ChatRoomRepository.GetChatRoomByID(roomID)
}

func (crs *ChatRoomService) DeleteChatRoomRecord() (models.ChatRoom, error) {
	return crs.ChatRoomRepository.DeleteChatRoomRecord()
}
