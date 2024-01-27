package docs

import (
	"github.com/gofiber/fiber/v2"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)


func SwaggerRouter(app *fiber.App){
	// protectedGroup := app.Group("/docs")

	
	SwaggerInfo.Title = "MeetSpace API"
	SwaggerInfo.Description = "This is a sample server Petstore server."
	SwaggerInfo.Version = "1.0"
	SwaggerInfo.Host = "localhost:8080"
	SwaggerInfo.Schemes = []string{"http", "https"}

	// Add Bearer Token authentication parameters to Swagger definition
	// SwaggerInfo.SecurityDefinitions = map[string]interface{}{
	// 	"Bearer": map[string]interface{}{
	// 		"type": "apiKey",
	// 		"name": "Authorization",
	// 		"in":   "header",
	// 	},
	// }

	
	// protectedGroup.Use(gin.BasicAuth(gin.Accounts{
	// 	"foo": "bar",
	// }))

	// protectedGroup.Get("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, 
	// ginSwagger.DefaultModelsExpandDepth(-1)))
}