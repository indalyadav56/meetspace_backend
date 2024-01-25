package routes

import (
	"meetspace_backend/client/handlers"

	"github.com/gin-gonic/gin"
)


func ClientRouter(e *gin.Engine){
	v1 := e.Group("/v1")
	
	clientRouter := v1.Group("/clients")
	clientUserRouter := v1.Group("/client/users")

	clientRouter.POST("", handlers.RegisterClientHandler)
	clientRouter.GET("/:clientId", handlers.GetClientById)
	clientRouter.GET("", handlers.GetAllClients)
	
	
	clientUserRouter.POST("", handlers.ClientAddUser)
	clientUserRouter.GET("", handlers.GetClientUsers)
}