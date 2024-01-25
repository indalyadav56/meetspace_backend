package http

import (
	"meetspace_backend/chat/handlers"

	"github.com/gin-gonic/gin"
)


func ChatRouter(e *gin.Engine){
	chatRouter := e.Group("/v1/chat")

	// rooms
	chatRouter.GET("/room/contact", handlers.GetChatRoomContact)
	chatRouter.GET("/rooms", handlers.GetChatRooms)
	chatRouter.POST("/rooms", handlers.CreateChatRoom)
	chatRouter.DELETE("/rooms", handlers.DeleteChatRoom)

	// groups
	chatRouter.POST("/groups", handlers.AddChatGroup)
	// chatGroup.GET("/group/members/:roomId", handlers.GetGroupMembers)

	// messages
	// chatGroup.POST("/messages", handlers.CreateChatMessageAPI)
	// chatGroup.PATCH("/messages", handlers.UpdateChatMessage)
	chatRouter.GET("/messages/:chatRoomId", handlers.GetChatMessageAPI)
}