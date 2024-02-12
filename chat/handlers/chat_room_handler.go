package handlers

import (
	"meetspace_backend/chat/models"
	"meetspace_backend/chat/services"
	"meetspace_backend/chat/types"
	"meetspace_backend/config"
	userModel "meetspace_backend/user/models"
	"meetspace_backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ChatRoomHandler struct {
    ChatRoomService *services.ChatRoomService

}

func NewChatRoomHandler(svc *services.ChatRoomService) *ChatRoomHandler {
    return &ChatRoomHandler{
        ChatRoomService: svc,
    }
}


type CreateChatRoomBody struct {
    RoomId uuid.UUID `json:"room_id"`
    RoomUsers []string `json:"room_users"`
}

// GetChatRoomContact godoc
//	@Summary		UserLogin User account
//	@Description	UserLogin User account
//	@Tags			Chat-Room
//	@Produce		json
//	@Param			user	body	types.LoginRequest	true	"User login details"
//	@Router			/v1/chat/room/contact [get]
//	@Security		Bearer
func (h *ChatRoomHandler) GetChatRoomContact(ctx *gin.Context){
    currentUser, exists := utils.GetUserFromContext(ctx)
    if !exists{
        return 
    }
    currentUserID := currentUser.ID

	var rooms []models.ChatRoom

	config.DB.Model(&models.ChatRoom{}).
	Select("id", "room_name", "is_group", "CreatedAt", "UpdatedAt").
	Preload("RoomUsers").
	Where("id IN (?)", 
	config.DB.Table("room_users").Select("chat_room_id").Where("user_id = ?", currentUserID)).
	Order("chat_rooms.updated_at DESC").Find(&rooms)

	var respData []types.ChatContactResponse
	var chatMessage models.ChatMessage

	for _, room := range rooms{
		config.DB.Where("chat_room_id = ?", room.ID).Find(&chatMessage).Order("updated_at DESC").Limit(1)
		if !room.IsGroup{
			for _, user := range room.RoomUsers{
				if user.ID != currentUserID{
					respData = append(respData, types.ChatContactResponse{
						RoomId: room.ID,
						IsGroup: room.IsGroup,
						UserId: &user.ID,
						Email: user.Email,
						FirstName: user.FirstName,
						LastName: user.LastName,
						IsActive: user.IsActive,
						LastMessage: chatMessage.Content,
						MessageUnSeenCount: 0,
					})
				}
			}
			}else{
				respData = append(respData, types.ChatContactResponse{
					RoomId: room.ID,
					IsGroup: room.IsGroup,
					RoomName: room.RoomName,
					LastMessage: chatMessage.Content,
					MessageUnSeenCount: 0,
			})
		}
		
	}
	
	ctx.JSON(http.StatusOK, utils.SuccessResponse("aeraeraewr", respData))
	return
}

// CreateChatRoom godoc
//	@Summary		CreateChatRoom 
//	@Description	CreateChatRoom
//	@Tags			Chat-Room
//	@Produce		json
//	@Param			user	body	types.LoginRequest	true	"User login details"
//	@Router			/v1/chat/rooms [post]
//	@Security		Bearer
func (h *ChatRoomHandler) CreateChatRoom (ctx *gin.Context){
    currentUserID := ctx.MustGet("userId")

    var reqBody CreateChatRoomBody
    
    if err := ctx.ShouldBindJSON(&reqBody); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
   
    var user1 userModel.User
    var user2 userModel.User

    
    config.DB.Where("id = ?", currentUserID).Find(&user1)
    config.DB.Where("id = ?", reqBody.RoomUsers[0]).Find(&user2)

    var result []struct {
        ChatRoomID string `gorm:"column:chat_room_id" json:"chat_room_id"`
    }

    config.DB.Table("room_users").
            Select("chat_room_id").
            Where("user_id IN (?,?)", currentUserID, user2.ID).
            Group("chat_room_id").
            Having("COUNT(DISTINCT user_id) = ?", 2).
            Find(&result)
    
    if(len(result) <= 0) {

        chatRoomData := models.ChatRoom{
            ID: reqBody.RoomId,
            RoomName: "NewChatRoom",
            RoomOwner: &user1,
            RoomUsers: []*userModel.User{
                &user1,
                &user2,
            },
        }
        config.DB.Create(&chatRoomData)
    
        ctx.JSON(http.StatusOK, utils.SuccessResponse("aeraer",chatRoomData))
        
        return
    }else{
        ctx.JSON(http.StatusOK, utils.SuccessResponse(
			"Room Already Created!!",
            result,
		))
        
        return

    }

}

// DeleteChatRoom godoc
//	@Summary		DeleteChatRoom 
//	@Description	DeleteChatRoom
//	@Tags			Chat-Room
//	@Produce		json
//	@Param			user	body	types.LoginRequest	true	"User login details"
//	@Router			/v1/chat/rooms [delete]
//	@Security		Bearer
func (h *ChatRoomHandler) DeleteChatRoom (ctx *gin.Context){
    return
}

// GetChatRooms godoc
//	@Summary		GetChatRooms
//	@Description	GetChatRooms
//	@Tags			Chat-Room
//	@Produce		json
//	@Router			/v1/chat/rooms [get]
//	@Security		Bearer
// @Param user_id query string false "User ID"
// @Param room_id query string false "Chat Room ID"
func (h *ChatRoomHandler) GetChatRooms(ctx *gin.Context){
    currentUser, _ := utils.GetUserFromContext(ctx)
    roomUserId := ctx.Query("user_id")
    roomId := ctx.Query("room_id")

    resp := h.ChatRoomService.GetChatRooms(currentUser.ID.String(), roomUserId, roomId)
    ctx.JSON(resp.StatusCode, resp)
}