package main

import (
	"fmt"
	authHandlers "meetspace_backend/auth/handlers"
	authRepo "meetspace_backend/auth/repositories"
	authRoutes "meetspace_backend/auth/routes"
	authServices "meetspace_backend/auth/services"
	chatHandlers "meetspace_backend/chat/handlers"
	chatRepo "meetspace_backend/chat/repositories"
	chatRoutes "meetspace_backend/chat/routes"
	chatServices "meetspace_backend/chat/services"
	websocketRoute "meetspace_backend/chat/websocket"
	commonServices "meetspace_backend/common/services"
	"meetspace_backend/config"
	"meetspace_backend/middlewares"
	userHandlers "meetspace_backend/user/handlers"
	userRepo "meetspace_backend/user/repositories"
	userRoutes "meetspace_backend/user/routes"
	userServices "meetspace_backend/user/services"
	"net/http"

	docs "meetspace_backend/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           MeetSpace API
// @version         1.0
// @description     MeetSpace API documentation.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and JWT token.
func main() {
	// load environment
	config.LoadEnv()
	
	// initialize database connection
	db := config.InitDB()
	redisDB := config.InitRedis()

	// repositories
	userRepo := userRepo.NewUserRepository(db)
	verificationRepo := authRepo.NewVerificationRepository(db)
	chatMessageRepo := chatRepo.NewChatMessageRepository(db)
	chatRoomRepo := chatRepo.NewChatRoomRepository(db)

	// services
	loggerService := commonServices.NewLoggerService()
	redisService := commonServices.NewRedisService(redisDB)
	userService := userServices.NewUserService(userRepo)
	tokenService := authServices.NewTokenService()
	authService := authServices.NewAuthService(loggerService, redisService, tokenService, userService)
	verificationService := authServices.NewVerificationService(verificationRepo)
	chatRoomService := chatServices.NewChatRoomService(chatRoomRepo, userService)
	chatGroupService := chatServices.NewChatGroupService(chatRoomRepo, userService)
	chatMessageService := chatServices.NewChatMessageService(chatMessageRepo, userService, chatRoomService)

	// handlers
	authHandler := authHandlers.NewAuthHandler(authService, verificationService)
	userHandler := userHandlers.NewUserHandler(userService)
	chatRoomHandler := chatHandlers.NewChatRoomHandler()
	chatGroupHandler := chatHandlers.NewChatGroupHandler(chatGroupService)
	chatMessageHandler := chatHandlers.NewChatMessageHandler(chatMessageService)
	
	r := gin.Default()

	// static
	r.StaticFS("/uploads", http.Dir("./uploads"))
	
	// middlewares
	r.Use(middlewares.LoggerMiddleware())
	r.Use(middlewares.CorsMiddleware())
	r.Use(middlewares.AuthMiddleware(loggerService, tokenService))
	
	// routes
	authRoutes.AuthRouter(r, authHandler)
	userRoutes.UserRouter(r, userHandler)
	chatRoutes.ChatRouter(r, chatRoutes.ChatHandlers{
		ChatRoomHandler: chatRoomHandler,
		ChatGroupHandler: chatGroupHandler, 
		ChatMessageHandler: chatMessageHandler,
	})
	websocketRoute.WebSocketRouter(r)

	// swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1)),
	)
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	fmt.Println("server:->", "http://localhost:8080")
	r.Run()
}
