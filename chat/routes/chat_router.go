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
	chatMessages := chatRouter.Group("/messages")

	// rooms
	chatRouter.GET("/contact", handler.ChatRoomHandler.GetChatRoomContact)

	chatRoom.GET("", handler.GetChatRooms)
	chatRoom.POST("", handler.CreateChatRoom)
	chatRoom.DELETE("/:charRoomID", handler.DeleteChatRoom)

	// groups
	chatGroups.POST("", handler.AddChatGroup)
	chatGroups.PATCH("", handler.UpdateChatGroup)

	chatGroup.GET("/members/:roomId", handler.GetGroupMembers)
	chatGroup.POST("/members", handler.GetGroupMembers)

	// messages
	chatMessages.POST("", handler.CreateChatMessage)
	chatMessages.GET("", handler.GetChatMessages)
	chatMessages.GET("/:chatRoomId", handler.GetChatMessageByRoomID)
	// chatGroup.PATCH("/", handlers.UpdateChatMessage)

	// call
	chatRouter.POST("/call", handler.HandleAudioVideoCall)
}