package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CorsMiddleware() fiber.Handler {

    config := cors.Config{
        AllowOrigins: "http://localhost:3000, *",
        AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
        AllowHeaders: "Origin, Content-Type, Accept",
        AllowCredentials: true,
    }

    return cors.New(config)
}