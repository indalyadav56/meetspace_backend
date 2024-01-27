package websocket

import (
	"github.com/gin-gonic/gin"
)

var groupPools = make(map[string]*Pool)

func WebSocketRouter(r *gin.Engine) {
	socketRouter := r.Group("/v1/chat")
	socketRouter.GET("", handleWebSocketConnection("new pool"))
	socketRouter.GET("/:groupName", handleWebSocketConnectionFromParam())
}