package services

import (
	"meetspace_backend/chat/models"
	"meetspace_backend/chat/repositories"
	"meetspace_backend/config"
	userModel "meetspace_backend/user/models"
	userService "meetspace_backend/user/services"
	"meetspace_backend/utils"

	"github.com/google/uuid"
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

func (crs *ChatRoomService) CreateChatRoom(roomName string, roomOwnerId string, roomUserIds []string) (models.ChatRoom, error) {
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
	return crs.ChatRoomRepository.CreateRecord(chatRoom)
}

func (crs *ChatRoomService) GetChatRoomByID(roomID string) (models.ChatRoom, error) {
	return crs.ChatRoomRepository.GetChatRoomByID(roomID)
}

func (r *ChatRoomService) GetChatRooms(currentUserID, roomUserId, roomId string) *utils.Response {
	
	if roomUserId != ""{
		type ChatRoomData struct{
			ChatRoomID string `gorm:"column:chat_room_id" json:"chat_room_id"`
		}
        var result []ChatRoomData
        
        config.DB.Table("room_users").
            Select("chat_room_id").
            Where("user_id IN (?,?)", currentUserID, roomUserId).
            Group("chat_room_id").
            Having("COUNT(DISTINCT user_id) = ?", 2).
            Find(&result)

		if len(result) < 1{
			uuid, _ := uuid.NewUUID() 
			result = append(result, ChatRoomData{ ChatRoomID: uuid.String()})
		}
		
		return utils.SuccessResponse("error", result)
        
    }else if roomId != ""{
        var room models.ChatRoom
		err := config.DB.Preload("RoomUsers", "id != ?", currentUserID).Where("id = ?", roomId).First(&room).Error
		if err != nil{
			return utils.ErrorResponse("Error", nil)
		}
		mapData, _ := utils.StructToMap(room)
		if room.IsGroup{
			delete(mapData, "room_users")
		}
        return utils.SuccessResponse("success", mapData)
    }else{
        var rooms []models.ChatRoom

        config.DB.Model(&models.ChatRoom{}).Preload("RoomUsers").Preload("RoomOwner").Where("id IN (?)", config.DB.Table("room_users").Select("chat_room_id").Where("user_id = ? AND is_group = ?", currentUserID, false)).Find(&rooms).Order("CreatedAt DESC")
        
        return utils.SuccessResponse("success", rooms)
    }
}

func (crs *ChatRoomService) DeleteChatRoomRecord() (models.ChatRoom, error) {
	return crs.ChatRoomRepository.DeleteChatRoomRecord()
}
