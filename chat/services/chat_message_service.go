package services

import (
	"fmt"
	"meetspace_backend/chat/models"
	"meetspace_backend/chat/repositories"
	"meetspace_backend/chat/types"
	"meetspace_backend/config"
	userModel "meetspace_backend/user/models"
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


// GetChatMessageByRoomId godoc
//	@Summary		GetChatMessageByRoomId
//	@Description	GetChatMessageByRoomId
//	@Tags			Chat-Message
//	@Produce		json
// @Param user body types.GetChatMessageRequestBody true "User registration details"
//	@Router			/v1/chat/messages/{room_id} [get]
func (chatMessageService *ChatMessageService) GetChatMessageByRoomId(roomID, userID string) ([]types.ChatMessageResponse, error){
    var messages []models.ChatMessage
    config.DB.Preload("Sender").Preload("ChatRoom").
    Where("chat_room_id=?", roomID).Order("created_at").
    Find(&messages)

    msgAsPerDay := make(map[string][]types.SingleChatMessageResponse)

    var respData types.SingleChatMessageResponse
    
    for _, message := range messages {
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
