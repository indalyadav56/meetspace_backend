package main

import (
	"fmt"

	"meetspace_backend/config"
	"meetspace_backend/middlewares"

	authRoute "meetspace_backend/auth/routes"
	userRoute "meetspace_backend/user/routes"

	// chatRoutes "meetspace_backend/chat/routes"
	// websocketRoute "meetspace_backend/chat/websocket"
	// clientRoutes "meetspace_backend/client/routes"

	_ "meetspace_backend/docs"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	// chatRoutes.ChatRouter(r)
	// websocketRoute.WebSocketRouter(r)
	// clientRoutes.ClientRouter(r)
	// docs.SwaggerRouter(r)

	fmt.Println("server:->", "http://localhost:8080")
    app.Listen(":8080")
}
