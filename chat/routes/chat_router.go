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

	// rooms
	chatRouter.GET("/room/contact", handler.ChatRoomHandler.GetChatRoomContact)
	chatRouter.GET("/rooms", handler.ChatRoomHandler.GetChatRooms)
	chatRouter.POST("/rooms", handler.ChatRoomHandler.CreateChatRoom)
	chatRouter.DELETE("/rooms", handler.ChatRoomHandler.DeleteChatRoom)

	// groups
	// chatGroup.POST("/groups", services.AddChatGroup)
	// chatGroup.GET("/group/members/:roomId", services.GetGroupMembers)

	// messages
	// chatGroup.POST("/messages", handlers.CreateChatMessageAPI)
	// chatGroup.PATCH("/messages", handlers.UpdateChatMessage)
	chatRouter.GET("/messages/:chatRoomId", handlers.GetChatMessageAPI)
}