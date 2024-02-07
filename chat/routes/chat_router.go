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
	chatGroup := chatRouter.Group("/group")
	chatGroups := chatRouter.Group("/groups")

	// rooms
	chatRouter.GET("/contact", handler.ChatRoomHandler.GetChatRoomContact)

	chatRoom.GET("", handler.GetChatRooms)
	chatRoom.POST("", handler.CreateChatRoom)
	chatRoom.DELETE("", handler.DeleteChatRoom)

	// groups
	chatGroups.POST("", handler.AddChatGroup)
	chatGroup.GET("/members/:roomId", handler.GetGroupMembers)

	// messages
	// chatGroup.POST("/messages", handlers.CreateChatMessageAPI)
	// chatGroup.PATCH("/messages", handlers.UpdateChatMessage)
	chatRouter.GET("/messages", handler.GetChatMessages)
	chatRouter.GET("/messages/:chatRoomId", handler.GetChatMessageByRoomID)
}