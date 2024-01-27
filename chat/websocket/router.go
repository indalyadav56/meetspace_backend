package websocket

import (
	"github.com/gofiber/fiber/v2"
)

var groupPools = make(map[string]*Pool)

func WebSocketRouter(app *fiber.App) {
	socketRouter := app.Group("/v1/chat")


	socketRouter.Get("", handleWebSocketConnection("new pool"))
	socketRouter.Get("/:groupName", handleWebSocketConnectionFromParam())
}
