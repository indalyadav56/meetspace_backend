package middlewares

import (
	"errors"
	authServices "meetspace_backend/auth/services"
	commonServices "meetspace_backend/common/services"
	"meetspace_backend/config"
	"meetspace_backend/user/models"
	"meetspace_backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	exemptedPaths = []string{
		"/v1/auth/login",
		"/v1/auth/register",
		"/v1/auth/forgot-password",
		"/v1/clients",
		"/v1/auth/send-email",
		"/v1/auth/verify-email",
		"/v1/user/check-email",
	}
)


func containsPath(urls []string, path string) bool {
	for _, url := range urls {
	  if strings.Contains(url, path) {
		return true
	  }
	}
	return false
}
  
func AuthMiddleware(loggerService *commonServices.LoggerService, tokenService *authServices.TokenService) gin.HandlerFunc  {
	return func(c *gin.Context) {
		// Allow WebSocket requests
        if upgrade := c.Request.Header.Get("Upgrade"); upgrade == "websocket" {
			websocketToken := c.Request.URL.Query().Get("token")
			if websocketToken != "" {
				claims, err := tokenService.VerifyToken(websocketToken, "access")
				userId := claims["user_id"]
				if err != nil {
					c.AbortWithStatus(401)
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

		exemptPathFromAuth(c)

		requestPath := c.Request.URL.Path
		// Check if the request path is exempted
		if containsPath(exemptedPaths, requestPath) ||  strings.HasPrefix(requestPath, "/docs/") {
			c.Next()
			return
		}

		tokenString, err := getTokenFromContext(c)
		if err != nil {
			resp := utils.UnauthorizedResponse("Token is required")
			c.JSON(resp.StatusCode, resp)
			c.Abort()
			return
		}

		claims, err := tokenService.VerifyToken(tokenString, "access")
		if err != nil {
			resp := utils.UnauthorizedResponse(err.Error())
            c.JSON(resp.StatusCode, resp) 
			c.Abort()
            return
        }

		var user models.User
		userId := claims["user_id"]
    	config.DB.Where("id = ?", userId).Find(&user)

		c.Set("user", user)
        c.Next() 
	}
  
}

func upgradeWebSocketConnection(c *gin.Context, tokenService *authServices.TokenService) {
	if upgrade := c.Request.Header.Get("Upgrade"); upgrade == "websocket" {
		websocketToken := c.Request.URL.Query().Get("token")

		if websocketToken != "" {
			claims, err := tokenService.VerifyToken(websocketToken, "access")
			userId := claims["user_id"]
			if err != nil {
				c.AbortWithStatus(401)
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
	c.Next()
}

func exemptPathFromAuth(c *gin.Context) {
	requestPath := c.Request.URL.Path
	// Check if the request path is exempted
	if containsPath(exemptedPaths, requestPath) ||  strings.HasPrefix(requestPath, "/docs/") {
		return
	}
}

func getTokenFromContext(c *gin.Context) (string, error) {
	err := errors.New("Token is required")
	authHeader := c.Request.Header.Get("Authorization")

	if len(strings.TrimSpace(authHeader) ) <= 0 {
		return "", err
	}

	if len(strings.TrimSpace(authHeader)) <= 0 {
		return "", err
	}
    return strings.TrimPrefix(authHeader, "Bearer "), nil
}