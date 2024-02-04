package handlers

import (
	"meetspace_backend/chat/services"
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


// GetChatMessageAPI godoc
//	@Summary		Register User account
//	@Description	Register User account
//	@Tags			Chat-Message
//	@Produce		json
//	@Router			/v1/chat/messages [get]
// @Security Bearer
func (h *ChatMessageHandler) GetChatMessageByRoomID(ctx *gin.Context){
    // get user from context
	currentUser, _ := utils.GetUserFromContext(ctx)
    chatRoomID := ctx.Param("chatRoomId")
    
	msg, _ := h.ChatMessageService.GetChatMessageByRoomId(chatRoomID, currentUser.ID.String())
    ctx.JSON(http.StatusOK, utils.SuccessResponse("success", msg))
    return
}


