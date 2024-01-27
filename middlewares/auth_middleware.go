package middlewares

import (
	"meetspace_backend/config"
	"meetspace_backend/user/models"
	"meetspace_backend/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func containsPath(urls []string, path string) bool {
	for _, url := range urls {
	  if strings.Contains(url, path) {
		return true
	  }
	}
	return false
}
  
func AuthMiddleware() gin.HandlerFunc  {
	return func(c *gin.Context) {
		// Allow WebSocket requests
        if upgrade := c.Request.Header.Get("Upgrade"); upgrade == "websocket" {
			websocketToken := c.Request.URL.Query().Get("token")

			if websocketToken != "" {
				userId, err := utils.VerifyToken(websocketToken)
        
				if err != nil {
					return
				}
		
				var user models.User
    			config.DB.Where("id = ?", userId).Find(&user)

				c.Set("user", user)
				c.Next() 
				return
			}
			c.AbortWithStatus(401)
			c.Next()
            return
        }

		authHeader := c.Request.Header.Get("Authorization")

		exemptedPaths := []string{
				"/v1/auth/login",
				"/v1/auth/register",
				"/v1/auth/forgot-password",
				"/v1/clients",
				"/v1/auth/send-email",
				"/v1/auth/verify-email",
				"/v1/user/check-email",
				"/",
			}
		requestPath := c.Request.URL.Path

		// Check if the request path is exempted
		if containsPath(exemptedPaths, requestPath) ||  strings.HasPrefix(requestPath, "/docs/") {
			c.Next() // Skip authentication and proceed to the next middleware
			return
		}
		
		if len(strings.TrimSpace(authHeader) ) <= 0 {
			c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Token is required", nil))
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if len(strings.TrimSpace(authHeader)) <= 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Success": false,
				"Message": "Token is required",
			})
			c.Abort()
			return
		}
		
		tokenString := extractToken(c)

		userId, err := utils.VerifyToken(tokenString)
        
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Success": false,
				"Message": "Token is invalid!",
			})
            c.AbortWithStatus(401) 
            return
        }

		var user models.User
    	config.DB.Where("id = ?", userId).Find(&user)

		c.Set("user", user)
        c.Next() 
	  }
  
  }


  func extractToken(c *gin.Context) string {
    //get token from Authorization header
    authHeader := c.GetHeader("Authorization") 
    return strings.TrimPrefix(authHeader, "Bearer ")
}