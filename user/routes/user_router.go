package routes

import (
	"meetspace_backend/user/handlers"

	"github.com/gin-gonic/gin"
)


func UserRouter(e *gin.Engine, handler *handlers.UserHandler){
	userRouter := e.Group("/v1/user")
	usersRouter := e.Group("/v1/users")

	usersRouter.POST("", handler.CreateUserHandler)
	usersRouter.GET("", handler.GetAllUsers)
	usersRouter.GET("/:userId", handler.GetUserByID)
	usersRouter.PATCH("", handler.UpdateUser)
	
	userRouter.GET("/search", handler.SearchUser)
	userRouter.GET("/profile", handler.GetUserProfile)
	userRouter.GET("/check-email", handler.CheckUserEmail)
}