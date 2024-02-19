package services

import (
	"fmt"
	"meetspace_backend/chat/models"
	"meetspace_backend/chat/repositories"
	"meetspace_backend/chat/types"
	"meetspace_backend/config"
	"meetspace_backend/utils"
	"time"

	userService "meetspace_backend/user/services"
)

type ChatMessageService struct {
	ChatMessageRepository *repositories.ChatMessageRepository
	UserService  *userService.UserService
	ChatRoomService  *ChatRoomService
}

func NewChatMessageService(
	repo *repositories.ChatMessageRepository,
	newUserService *userService.UserService,
	chatRoomService *ChatRoomService,
) *ChatMessageService {
	return &ChatMessageService{
		ChatMessageRepository: repo,
		UserService: newUserService,
		ChatRoomService: chatRoomService,
	}
}

func (chatMessageService *ChatMessageService) CreateChatMessage(currentUserID string, reqBody types.CreateChatRequestBody) (models.ChatMessage, error) {
	senderUser, _ := chatMessageService.UserService.UserRepository.GetUserByID(currentUserID)
	chatRoom, err := chatMessageService.ChatRoomService.GetChatRoomByID(reqBody.RoomID)

	if err != nil{
		fmt.Println("chat room give not found", reqBody)
		createdChatRoom, _ := chatMessageService.ChatRoomService.CreateChatRoom(
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
	
		return chatMessageService.ChatMessageRepository.CreateRecord(chatMessage)
	}

	fmt.Println("chat room found", reqBody)

	chatMessage := models.ChatMessage{
		Content: reqBody.Content,
		Sender: senderUser,
		ChatRoom: &chatRoom,
	}

	return chatMessageService.ChatMessageRepository.CreateRecord(chatMessage)
}


func formatDate(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)

	// Within today
	if diff.Hours() < 24 {
		if diff.Hours() < 1 {
			return "Today"
		} else {
			return "Yesterday"
		}
	}

	// Within this week
	if diff.Hours() < 168 {
		dayNum := int(diff.Hours() / 24)
		dayName := t.Weekday().String()[:3]
		return fmt.Sprintf("%s (%s)", dayName, dayNum)
	}

	// Within this year
	if now.Year() == t.Year() {
		return fmt.Sprintf("%d-%s", t.Day(), t.Month().String()[:3])
	}

	// Default format
	return fmt.Sprintf("%d-%s-%d", t.Day(), t.Month().String()[:3], t.Year())
}


func (chatMessageService *ChatMessageService) GetChatMessageByRoomId(roomID, userID string) ([]types.ChatMessageResponse, error){
    var messages []models.ChatMessage
    config.DB.Preload("Sender").Preload("ChatRoom").
    Where("chat_room_id=?", roomID).Order("created_at").
    Find(&messages)

    msgAsPerDay := make(map[string][]types.SingleChatMessageResponse)

    var respData types.SingleChatMessageResponse
    
    for _, message := range messages {
        respData.ID = message.ID.String()
        respData.Content = message.Content
        respData.ChatRoomId = message.ChatRoomID.String()
        respData.Sender = message.Sender
        msgAsPerDay[formatDate(message.CreatedAt)] = append(msgAsPerDay[formatDate(message.CreatedAt)], respData)
    }

    var resp []types.ChatMessageResponse

    for timestamp, msg := range msgAsPerDay {
        resp = append(resp, types.ChatMessageResponse{
            TimeStamp: timestamp,
            ChatMessage: msg,
        })
    }
    return resp, nil
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
