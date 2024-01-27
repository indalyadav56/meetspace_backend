package main

import (
	"fmt"

	"meetspace_backend/config"
	"meetspace_backend/middlewares"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	authRoute "meetspace_backend/auth/routes"
	chatRoute "meetspace_backend/chat/routes"
	websocketRoute "meetspace_backend/chat/websocket"
	clientRoute "meetspace_backend/client/routes"
	_ "meetspace_backend/docs"
	userRoute "meetspace_backend/user/routes"
)

// @securityDefinitions.basic BasicAuth
// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
func main() {
	// load environment
	config.LoadEnv()
	
	// initialize database connection
	config.InitDB()
	
	app := fiber.New()

	app.Static("/uploads","./uploads")
	
	// middlewares
    app.Use(logger.New())
	app.Use(middlewares.LoggerMiddleware())
	app.Use(middlewares.CorsMiddleware())
	// r.Use(middlewares.AuthMiddleware())
	
	// routes
	app.Get("/*", swagger.HandlerDefault)
	
	authRoute.AuthRouter(app)
	userRoute.UserRouter(app)
	chatRoute.ChatRouter(app)
	clientRoute.ClientRouter(app)
	websocketRoute.WebSocketRouter(app)

	fmt.Println("server:->", "http://localhost:8080")
    app.Listen(":8080")
}
