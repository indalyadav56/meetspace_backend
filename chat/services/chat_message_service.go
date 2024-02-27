package services

import (
	"fmt"
	"meetspace_backend/chat/constants"
	"meetspace_backend/chat/models"
	"meetspace_backend/chat/repositories"
	"meetspace_backend/chat/types"
	"meetspace_backend/config"
	userModels "meetspace_backend/user/models"
	"meetspace_backend/utils"

	commonServices "meetspace_backend/common/services"
	userService "meetspace_backend/user/services"
)

type ChatMessageService struct {
	ChatMessageRepository *repositories.ChatMessageRepository
	UserService  *userService.UserService
	ChatRoomService  *ChatRoomService
	RedisService *commonServices.RedisService
}

func NewChatMessageService(
	repo *repositories.ChatMessageRepository,
	newUserService *userService.UserService,
	chatRoomService *ChatRoomService,
	redisService *commonServices.RedisService,
) *ChatMessageService {
	return &ChatMessageService{
		ChatMessageRepository: repo,
		UserService: newUserService,
		ChatRoomService: chatRoomService,
		RedisService: redisService,
	}
}

func (c *ChatMessageService) CreateChatMessage(currentUserID string, reqBody types.CreateChatRequestBody) (models.ChatMessage, error) {
	senderUser, _ := c.UserService.UserRepository.GetUserByID(currentUserID)
	chatRoom, err := c.ChatRoomService.GetChatRoomByID(reqBody.RoomID)

	if err != nil{
		createdChatRoom, _ := c.ChatRoomService.CreateChatRoom(
			reqBody.RoomID,
			"NewChatRoom",
			currentUserID,
			[]string{reqBody.RecieverUserID},
		)
		chatMessage := models.ChatMessage{
			Content: reqBody.Content,
			Sender: senderUser,
			ChatRoom: &createdChatRoom,
		}
	
		return c.ChatMessageRepository.CreateRecord(chatMessage)
	}

	chatMessage := models.ChatMessage{
		Content: reqBody.Content,
		Sender: senderUser,
		ChatRoom: &chatRoom,
	}

	c.CheckNotification(senderUser, reqBody.RoomID, reqBody.Content)
	
	return c.ChatMessageRepository.CreateRecord(chatMessage)
}

// GroupMessagesByDate groups chat messages by date and formats them into the desired response structure
func GroupMessagesByDate(messages []models.ChatMessage) []types.ChatMessageResponse {
	groupedMessages := make(map[string][]models.ChatMessage)
	for _, message := range messages {
		date := message.CreatedAt.Format("2006-01-02")
		groupedMessages[date] = append(groupedMessages[date], message)
	}

	var chatMessages []types.ChatMessageResponse
	for date, msgs := range groupedMessages {
		var singleChatMessages []types.SingleChatMessageResponse
		for _, msg := range msgs {
			// Convert models.ChatMessage to SingleChatMessageResponse
			singleChatMessage := types.SingleChatMessageResponse{
				ID:      msg.ID.String(),
				Content: msg.Content,
				Sender: msg.Sender,
				ChatRoomId: msg.ChatRoomID.String(),
				CreatedAt : msg.CreatedAt.String(),
			}
			singleChatMessages = append(singleChatMessages, singleChatMessage)
		}

		// Create ChatMessageResponse for the current date
		chatMessageResponse := types.ChatMessageResponse{
			TimeStamp:   date,
			ChatMessage: singleChatMessages,
		}
		chatMessages = append(chatMessages, chatMessageResponse)
	}
	return chatMessages
}

func (chatMessageService *ChatMessageService) GetChatMessageByRoomId(roomID, userID string) *utils.Response{
    var messages []models.ChatMessage
    
	config.DB.Preload("Sender").Preload("ChatRoom").
    Where("chat_room_id=?", roomID).Order("created_at").
    Find(&messages)

	groupMessage := GroupMessagesByDate(messages)

    return utils.SuccessResponse("successfully fetched!", groupMessage)
}

func (s *ChatMessageService) GetChatMessageByUserID(currentUserID, otherUserID string) *utils.Response{
	users :=  []string{currentUserID, otherUserID}
    // var messages []models.ChatMessage
	// var chatRoom models.ChatRoom
	
	// chatRoom2, _ :=  s.ChatRoomService.CreateChatRoomRecord("room-2", currentUserID, users)
	// s.CreateChatMessage("test-users msg -create", currentUserID, chatRoom2.ID.String())
	var rooms []models.ChatRoom

	config.DB.Model(&models.ChatRoom{}).
	Select("id", "room_name", "is_group", "CreatedAt", "UpdatedAt").
	Preload("RoomUsers").
	Where("id IN (?)", 
	config.DB.Table("room_users").Select("chat_room_id").Where("user_id IN (?)", users)).
	Order("chat_rooms.updated_at DESC").Find(&rooms)

	// if len(rooms) >= 2 {
	// 	// If both users are in the room, fetch all messages in the room
	// 	config.DB.Preload("ChatRoom").
	// 		Where("chat_room_id = ?", chatRoom.ID).
	// 		Order("created_at").
	// 		Find(&messages)
	
	// 	fmt.Println("messages:->", messages)
	// } else {
	// 	fmt.Println("Chat room not found or doesn't contain both users")
	// }

	// fmt.Println("messages:->", messages)
    return utils.SuccessResponse("sucess", nil)
}


// call this func when adding msg in chat-room
func (c *ChatMessageService) CheckNotification(senderUser *userModels.User,roomID, lastMessage string){
	// Get group members from redis
	currentGroup := fmt.Sprintf("client:group:%v", roomID)
	members, err := c.RedisService.SMembers(currentGroup).Result()
	if err != nil {
		panic(err)
	}

	var chatRoomObj models.ChatRoom

	config.DB.Preload("RoomUsers").Where("id=?", roomID).Find(&chatRoomObj)

	for _, userObj := range chatRoomObj.RoomUsers {
		roomUserId := userObj.ID.String()
		
		exists := containsString(members, roomUserId)
		if !exists {
			payload := types.Payload{
				Event: constants.CHAT_NOTIFICATION_SENT,
				Data: map[string]interface{}{
					"room_id" : roomID,
					"receiver_user" : userObj,
					"sender_user" : senderUser,
					"content": lastMessage,
					"is_group": chatRoomObj.IsGroup,
					"room_name": chatRoomObj.RoomName,
				},
			}
			strData, _ := utils.StructToString(payload)
			c.RedisService.Publish("client", strData)
		}
	}


}

func containsString(list []string, target string) bool {
	for _, item := range list {
	  if item == target {
		return true
	  }
	}
	return false
  }