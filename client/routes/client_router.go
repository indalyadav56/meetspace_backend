package routes

import (
	"meetspace_backend/client/handlers"

	"github.com/gofiber/fiber/v2"
)


func ClientRouter(e *fiber.App){
	v1 := e.Group("/v1")
	
	clientRouter := v1.Group("/clients")
	clientUserRouter := v1.Group("/client/users")

	clientRouter.Post("", handlers.RegisterClientHandler)
	clientRouter.Get("/:clientId", handlers.GetClientById)
	clientRouter.Get("", handlers.GetAllClients)
	
	clientUserRouter.Post("", handlers.ClientAddUser)
	clientUserRouter.Get("", handlers.GetClientUsers)
}