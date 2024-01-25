package handlers

import (
	"fmt"
	"meetspace_backend/chat/models"
	"meetspace_backend/chat/types"
	"meetspace_backend/config"
	userModel "meetspace_backend/user/models"
	"meetspace_backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateChatRoomBody struct {
    RoomId uuid.UUID `json:"room_id"`
    RoomUsers []string `json:"room_users"`
}

// GetChatRoomContact godoc
//	@Summary		UserLogin User account
//	@Description	UserLogin User account
//	@Tags			Chat-Room
//	@Produce		json
// @Param user body types.LoginRequest true "User login details"
//	@Router			/v1/chat/room/contact [get]
func GetChatRoomContact(ctx *gin.Context){
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

    fmt.Println("currentUserID", currentUserID)
    fmt.Println("chatrooms", rooms)
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
// @Param user body CreateChatRoomBody true "User login details"
//	@Router			/v1/chat/rooms [post]
func CreateChatRoom (ctx *gin.Context){
    user, _ := utils.GetUserFromContext(ctx)
    currentUserID := user.ID
    var reqBody CreateChatRoomBody
    
    if err := ctx.ShouldBindJSON(&reqBody); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
   
    var user1 userModel.User
    var user2 userModel.User

    fmt.Println("indal", reqBody)
    
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
    fmt.Println("config", result)
    if(len(result) <= 0) {
        chatRoomData := models.ChatRoom{
            ID: reqBody.RoomId,
            RoomName: "NewChatRoom",
            RoomOwnerID: user1.ID,
            RoomOwner: &user1,
            RoomUsers: []*userModel.User{
                &user1,
                &user2,
            },
        }
        if err := config.DB.Create(&chatRoomData).Error; err != nil {
            fmt.Println("create room error", err)
        }
        fmt.Println("chatRoomData", chatRoomData)
        ctx.JSON(http.StatusOK, utils.SuccessResponse("aeraer",chatRoomData))
        
        return
    }else{
        ctx.JSON(http.StatusOK, utils.ErrorResponse(
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
// @Param user body types.LoginRequest true "User login details"
//	@Router			/v1/chat/rooms [delete]
func DeleteChatRoom (ctx *gin.Context){

}


// GetChatRooms godoc
//	@Summary		GetChatRooms
//	@Description	GetChatRooms
//	@Tags			Chat-Room
//	@Produce		json
// @Param user body types.LoginRequest true "User login details"
//	@Router			/v1/chat/rooms [get]
func GetChatRooms(ctx *gin.Context){
    currentUser, exists := utils.GetUserFromContext(ctx)
    if !exists{
        return 
    }
    currentUserID := currentUser.ID

    roomUserId := ctx.Query("user_id")

    if roomUserId != ""{
        
        var result []struct {
            ChatRoomID string `gorm:"column:chat_room_id" json:"chat_room_id"`
        }
        
        config.DB.Table("room_users").
            Select("chat_room_id").
            Where("user_id IN (?,?)", currentUserID, roomUserId).
            Group("chat_room_id").
            Having("COUNT(DISTINCT user_id) = ?", 2).
            Find(&result)

		ctx.JSON(http.StatusOK, utils.SuccessResponse(
			"success",
            result,
		))
        return
        
    }else{
        var rooms []models.ChatRoom

        config.DB.Model(&models.ChatRoom{}).Preload("RoomUsers").Preload("RoomOwner").Where("id IN (?)", config.DB.Table("room_users").Select("chat_room_id").Where("user_id = ?", currentUserID)).Find(&rooms).Order("CreatedAt DESC")
        
        ctx.JSON(http.StatusOK, utils.SuccessResponse(
			"aeraer",
            rooms,
		))
        
        return
    }
}