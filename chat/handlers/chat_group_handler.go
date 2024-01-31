package handlers

import (
	"meetspace_backend/chat/models"
	"meetspace_backend/chat/types"
	"meetspace_backend/config"
	userModel "meetspace_backend/user/models"
	"meetspace_backend/utils"

	"github.com/gofiber/fiber/v2"
)

type ChatGroupHandler struct {

}

func NewChatGroupHandler() *ChatGroupHandler {
    return &ChatGroupHandler{
        
    }
}



// AddChatGroup godoc
//	@Summary		UserLogin User account
//	@Description	UserLogin User account
//	@Tags			Chat-Group
//	@Produce		json
// @Param user body types.LoginRequest true "User login details"
//	@Router			/v1/chat/room/groups [post]
// @Security Bearer
func AddChatGroup(ctx *fiber.Ctx) error{
    // currentUser, exists := utils.GetUserFromContext(ctx)
    // if !exists{
    //     return nil
    // }
	var reqData types.AddChatGroup

	var chatRoom models.ChatRoom
	var roomUsers []*userModel.User

	chatRoom.IsGroup = true
	chatRoom.RoomName = reqData.Title

	for _, userId := range reqData.UserIds {
		user, err := config.ChatGroupService.UserService.UserRepository.GetUserByID(userId)
		if err == nil {
			roomUsers = append(roomUsers, &user)
		}
	}

	// roomUsers = append(roomUsers, currentUser)
	// chatRoom.RoomUsers = roomUsers
	// chatRoom.RoomOwner = currentUser

	chatGroup, _ := config.ChatGroupService.CreateChatGroup(chatRoom)
	ctx.JSON(utils.SuccessResponse("success", chatGroup))
	return nil
}