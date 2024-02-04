package http

import (
	"meetspace_backend/chat/handlers"

	"github.com/gin-gonic/gin"
)

type ChatHandlers struct {
	*handlers.ChatRoomHandler
	*handlers.ChatGroupHandler
	*handlers.ChatMessageHandler
}


func ChatRouter(e *gin.Engine, handler ChatHandlers){
	chatRouter := e.Group("/v1/chat")

	chatRoom := chatRouter.Group("/rooms")
	chatGroup := chatRouter.Group("/groups")

	// rooms
	chatRouter.GET("/contact", handler.ChatRoomHandler.GetChatRoomContact)

	chatRoom.GET("", handler.ChatRoomHandler.GetChatRooms)
	chatRoom.POST("", handler.ChatRoomHandler.CreateChatRoom)
	chatRoom.DELETE("", handler.ChatRoomHandler.DeleteChatRoom)

	// groups
	chatGroup.POST("", handler.ChatGroupHandler.AddChatGroup)
	// chatGroup.GET("/group/members/:roomId", services.GetGroupMembers)

	// messages
	// chatGroup.POST("/messages", handlers.CreateChatMessageAPI)
	// chatGroup.PATCH("/messages", handlers.UpdateChatMessage)
	chatRouter.GET("/messages/:chatRoomId", handler.GetChatMessageByRoomID)
}