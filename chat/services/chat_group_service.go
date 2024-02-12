package services

import (
	"meetspace_backend/auth/constants"
	"meetspace_backend/chat/models"
	"meetspace_backend/chat/repositories"
	"meetspace_backend/chat/types"
	"meetspace_backend/config"
	userModel "meetspace_backend/user/models"
	userService "meetspace_backend/user/services"
	"meetspace_backend/utils"

	"github.com/google/uuid"
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

func (s *ChatGroupService) CreateChatGroup(user *userModel.User, reqData types.AddChatGroup) *utils.Response {
	// validate request struct data
	if err := utils.GetValidator().Struct(reqData); err != nil {
		data := utils.ParseError(err, reqData)
		return utils.ErrorResponse(constants.AUTH_REQUEST_VALIDATION_ERROR_MSG, data)
    }

	var chatRoom models.ChatRoom
	var roomUsers []*userModel.User

	chatRoom.IsGroup = true
	chatRoom.RoomName = reqData.Title
	chatRoom.RoomOwner = user

	for _, userId := range reqData.UserIds {
		user, err := s.UserService.UserRepository.GetUserByID(userId)
		if err == nil {
			roomUsers = append(roomUsers, user)
		}
	}

	roomUsers = append(roomUsers, user)
	chatRoom.RoomUsers = roomUsers

	createdChatRoom, err := s.ChatRoomRepository.CreateRecord(chatRoom)
	if err != nil {
		return utils.ErrorResponse(err.Error(), nil)
	}

	respData := map[string]interface{}{
		"id": createdChatRoom.ID.String(),
		"room_name": chatRoom.RoomName,
	}
	return utils.SuccessResponse("Successfully created chat group!", respData)

}

func (s *ChatGroupService) GetGroupMembers(roomId string) *utils.Response{
	var chatRooms []models.ChatRoom

	type GroupMemberResponse struct {
		UserId uuid.UUID `json:"id"`
		Email string `json:"email"`
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
	}
	
    
    var members []GroupMemberResponse

    config.DB.Preload("RoomUsers").Where("id = ?", roomId).Find(&chatRooms)

    for _, group := range chatRooms{
        
        for _, user := range group.RoomUsers{
            members = append(members, GroupMemberResponse{
                UserId: user.ID,
                FirstName: user.FirstName,
                LastName: user.LastName,
                Email: user.Email,
            })
        }
    }
    return utils.SuccessResponse("success", members)
}

func (s *ChatGroupService) AddGroupMembers(roomId string) *utils.Response{
	// var chatRoom models.ChatRoom
	// var members []GroupMember
    return utils.SuccessResponse("success", nil)
}