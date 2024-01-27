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
	roomOnwer := crs.UserService.GetUserByID(roomOwnerId)
	roomOnwerResp := roomOnwer.Data.(userModel.User)
	var roomUsers []*userModel.User
	
	for _, userId := range roomUserIds {
		user := crs.UserService.GetUserByID(userId)
		userResponse := user.Data.(userModel.User)
		roomUsers = append(roomUsers, &userResponse)
	}

	roomUsers = append(roomUsers,&roomOnwerResp)

	chatRoom := models.ChatRoom{
		RoomName: roomName,
		RoomOwner: &roomOnwerResp,
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
