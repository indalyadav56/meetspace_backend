package handlers

import (
	"meetspace_backend/chat/services"
	"meetspace_backend/chat/types"
	"meetspace_backend/utils"

	"github.com/gin-gonic/gin"
)

type ChatGroupHandler struct {
	ChatGroupService *services.ChatGroupService
}

func NewChatGroupHandler(svc *services.ChatGroupService) *ChatGroupHandler {
    return &ChatGroupHandler{
		ChatGroupService: svc,
    }
}

// AddChatGroup godoc
//	@Summary		add-chat-group
//	@Description	Add Chat group
//	@Tags			Chat-Group
//	@Produce		json
// @Param user body types.AddChatGroup true "Add chat group details"
//	@Router			/v1/chat/groups [post]
// @Security Bearer
func (h *ChatGroupHandler) AddChatGroup(ctx *gin.Context){
	currentUser, _ := utils.GetUserFromContext(ctx)
  
	var req types.AddChatGroup
	if errResp := utils.BindJsonData(ctx, &req); errResp != nil {
		ctx.JSON(errResp.StatusCode, errResp)
        return
    }

	resp := h.ChatGroupService.CreateChatGroup(currentUser, req)
    
	ctx.JSON(resp.StatusCode, resp)
	return
}
