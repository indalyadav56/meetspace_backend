package main

import (
	"fmt"
	authRoutes "meetspace_backend/auth/routes"
	chatRoutes "meetspace_backend/chat/routes"
	websocketRoute "meetspace_backend/chat/websocket"
	clientRoutes "meetspace_backend/client/routes"
	"meetspace_backend/config"
	docs "meetspace_backend/docs"
	"meetspace_backend/middlewares"
	userRoutes "meetspace_backend/user/routes"
	"net/http"

	"github.com/gin-gonic/gin"
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
	
	r := gin.Default()

	r.StaticFS("/uploads", http.Dir("./uploads"))
	
	// middlewares
	r.Use(middlewares.LoggerMiddleware())
	r.Use(middlewares.CorsMiddleware())
	r.Use(middlewares.AuthMiddleware())
	
	// routes
	authRoutes.AuthRouter(r)
	userRoutes.UserRouter(r)
	chatRoutes.ChatRouter(r)
	websocketRoute.WebSocketRouter(r)
	clientRoutes.ClientRouter(r)
	docs.SwaggerRouter(r)

	fmt.Println("server:->", "http://localhost:8080")
	r.Run()
}