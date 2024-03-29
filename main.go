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
	websocket "meetspace_backend/chat/websocket"
	commonServices "meetspace_backend/common/services"
	"meetspace_backend/config"
	"meetspace_backend/middlewares"
	userHandlers "meetspace_backend/user/handlers"
	userRepo "meetspace_backend/user/repositories"
	userRoutes "meetspace_backend/user/routes"
	userServices "meetspace_backend/user/services"
	"net/http"
	"os"

	docs "meetspace_backend/docs"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/prometheus/client_golang/prometheus/promhttp"
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

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and JWT token.
func main() {
	// load environment
	config.LoadEnv()
	
	// sso providers
	const BaseURL = "http://localhost:8080"

	goth.UseProviders(
		github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), BaseURL+"/v1/auth/github/callback"),
		google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), BaseURL+"/v1/auth/google/callback"),
	)

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
	loggerService.Infof("Hello There")
	redisService := commonServices.NewRedisService(redisDB)
	userService := userServices.NewUserService(userRepo)
	tokenService := authServices.NewTokenService()
	authService := authServices.NewAuthService(loggerService, redisService, tokenService, userService)
	verificationService := authServices.NewVerificationService(verificationRepo)
	chatRoomService := chatServices.NewChatRoomService(chatRoomRepo, userService, redisService)
	chatGroupService := chatServices.NewChatGroupService(chatRoomRepo, userService)
	chatMessageService := chatServices.NewChatMessageService(chatMessageRepo, userService, chatRoomService, redisService)

	// handlers
	authHandler := authHandlers.NewAuthHandler(authService, verificationService)
	userHandler := userHandlers.NewUserHandler(userService)
	chatRoomHandler := chatHandlers.NewChatRoomHandler(chatRoomService)
	chatGroupHandler := chatHandlers.NewChatGroupHandler(chatGroupService)
	chatMessageHandler := chatHandlers.NewChatMessageHandler(chatMessageService)
	wsHandler := websocket.NewWebSocketHandler(redisService)
	
	r := gin.Default()

	// static
	r.StaticFS("/uploads", http.Dir("./uploads"))
	
	// middlewares
	r.Use(middlewares.LoggerMiddleware(loggerService))
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
	websocket.WebSocketRouter(r, wsHandler)

	// swagger
	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/docs/index.html") 
	})
	r.GET("/docs/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1)),
	)
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Prometheus metrics
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	fmt.Println("server running at:- ", BaseURL)
	r.Run()
}
