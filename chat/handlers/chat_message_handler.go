package handlers

import (
	"fmt"
	"meetspace_backend/chat/services"
	"meetspace_backend/chat/types"
	"meetspace_backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChatMessageHandler struct {
	ChatMessageService *services.ChatMessageService
}

func NewChatMessageHandler(svc *services.ChatMessageService) *ChatMessageHandler {
    return &ChatMessageHandler{
        ChatMessageService: svc,
    }
}

// CreateChatMessage godoc
//	@Summary		add chat message
//	@Description	add chat message
//	@Tags			Chat-Message
//	@Produce		json
//	@Router			/v1/chat/messages [post]
//	@Security		Bearer
//	@Param			user	body	types.CreateChatRequestBody	true	"add chat message details"
//	@Success		201	"add chat message successfully"
//	@Failure		400	"Bad request"
//	@Failure		500	"Internal server error"
func (h *ChatMessageHandler) CreateChatMessage(ctx *gin.Context){
    // get user from context
	currentUser, _ := utils.GetUserFromContext(ctx)
    
	var reqBody types.CreateChatRequestBody
	
	if err := utils.BindJsonData(ctx, &reqBody); err != nil {
		resp:= utils.ErrorResponse("Invalid JSON", nil)
		ctx.JSON(resp.StatusCode, resp)
		return
	}

	msg, err := h.ChatMessageService.CreateChatMessage(currentUser.ID.String(), reqBody)
	fmt.Println("error getting", err)
    ctx.JSON(http.StatusOK, utils.SuccessResponse("success", msg))
    return
}

// GetChatMessageAPI godoc
//	@Summary		get chat messages by room id
//	@Description	Get chat messages by room id
//	@Tags			Chat-Message
//	@Produce		json
//	@Param			chat_room_id	path	string	true	"Chat Room ID"
//	@Router			/v1/chat/messages/{chat_room_id} [get]
//	@Security		Bearer
//	@Success		201	"get messages successfully"
//	@Failure		400	"Bad request"
//	@Failure		500	"Internal server error"
func (h *ChatMessageHandler) GetChatMessageByRoomID(ctx *gin.Context){
    // get user from context
	currentUser, _ := utils.GetUserFromContext(ctx)
    chatRoomID := ctx.Param("chatRoomId")
	data := h.ChatMessageService.GetChatMessageByRoomId(chatRoomID, currentUser.ID.String())
    ctx.JSON(data.StatusCode, data)
    return
}

// GetChatMessages godoc
//	@Summary		get chat messages
//	@Description	Get chat message
//	@Tags			Chat-Message
//	@Produce		json
// @Param user_id query string true "User ID"
//	@Router			/v1/chat/messages [get]
//	@Security		Bearer
//	@Success		201	"get messages successfully"
//	@Failure		400	"Bad request"
//	@Failure		500	"Internal server error"
func (h *ChatMessageHandler) GetChatMessages(ctx *gin.Context){
	currentUser, _ := utils.GetUserFromContext(ctx)
    userID := ctx.Query("user_id")
    
	resp := h.ChatMessageService.GetChatMessageByUserID(currentUser.ID.String(), userID)
    ctx.JSON(resp.StatusCode, resp)
    return
}

