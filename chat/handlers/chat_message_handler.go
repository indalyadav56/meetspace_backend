package handlers

import (
	"fmt"
	"meetspace_backend/chat/models"
	"meetspace_backend/chat/types"
	"meetspace_backend/config"
	"meetspace_backend/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ChatMessageHandler struct {

}

func NewChatMessageHandler() *ChatMessageHandler {
    return &ChatMessageHandler{
        
    }
}


// GetChatMessageAPI godoc
//	@Summary		Register User account
//	@Description	Register User account
//	@Tags			Chat-Message
//	@Produce		json
// @Param user body types.GetChatMessageRequestBody true "User registration details"
//	@Router			/v1/chat/messages [get]
func GetChatMessageAPI(ctx *gin.Context){
    chatRoomID := ctx.Param("chatRoomId")
    
    var reqBody types.GetChatMessageRequestBody

    reqBody.ChatRoomId = chatRoomID
   
	msg, _ := GetChatMessageByRoomId(reqBody)
    
    ctx.JSON(http.StatusOK, utils.SuccessResponse("success", msg))
    return
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
func GetChatMessageByRoomId(inputData types.GetChatMessageRequestBody) ([]types.ChatMessageResponse, error){
    var messages []models.ChatMessage
    config.DB.Preload("Sender").Preload("ChatRoom").
    Where("chat_room_id=?", inputData.ChatRoomId).Order("created_at").
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
