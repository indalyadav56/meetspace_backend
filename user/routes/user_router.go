package routes

import (
	"meetspace_backend/user/handlers"

	"github.com/gofiber/fiber/v2"
)


func UserRouter(app *fiber.App){
	userRouter := app.Group("/v1/user")
	usersRouter := app.Group("/v1/users")

	usersRouter.Post("", handlers.CreateUserHandler)
	usersRouter.Get("", handlers.GetAllUsers)
	usersRouter.Get("/:userId", handlers.GetUserByID)
	usersRouter.Patch("", handlers.UpdateUser)

	userRouter.Get("/check-email", handlers.CheckUserEmail)
}