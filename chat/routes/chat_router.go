package http

import (
	"meetspace_backend/chat/handlers"

	"github.com/gofiber/fiber/v2"
)


func ChatRouter(app *fiber.App){
	chatRouter := app.Group("/v1/chat")
	chatRoom := app.Group("/v1/chat/rooms")

	// rooms
	chatRouter.Get("/room/contact", handlers.GetChatRoomContact)
	
	chatRoom.Get("/", handlers.GetChatRooms)
	chatRoom.Post("/rooms", handlers.CreateChatRoom)
	chatRoom.Delete("/rooms", handlers.DeleteChatRoom)

	// groups
	chatRouter.Post("/groups", handlers.AddChatGroup)
	// chatGroup.GET("/group/members/:roomId", handlers.GetGroupMembers)

	// messages
	// chatGroup.POST("/messages", handlers.CreateChatMessageAPI)
	// chatGroup.PATCH("/messages", handlers.UpdateChatMessage)
	chatRouter.Get("/messages/:chatRoomId", handlers.GetChatMessageAPI)
}