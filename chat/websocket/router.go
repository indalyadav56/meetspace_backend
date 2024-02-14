package websocket

import (
	"github.com/gin-gonic/gin"
)

var groupPools = make(map[string]*Pool)

func WebSocketRouter(r *gin.Engine, handler *WebSocketHandler) {
	socketRouter := r.Group("/v1/chat")
	socketRouter.GET("", handler.handleWebSocketConnection("new pool"))
	socketRouter.GET("/:groupName", handler.handleWebSocketConnectionFromParam())
}