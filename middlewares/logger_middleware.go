package middlewares

import (
	commonServices "meetspace_backend/common/services"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func LoggerMiddleware(loggerService *commonServices.LoggerService) gin.HandlerFunc {
    log.SetFormatter(&log.JSONFormatter{}) // Use JSON format for logs

    return func(c *gin.Context) {
        start := time.Now()

        fileName := "logs/app-" + time.Now().Format("2006-01-02") + ".log"
        file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
        if err == nil {
            log.SetOutput(file)
        } else {
            log.Info("Failed to open log file, using default stderr")
        }

        // Log request details
        log.WithFields(log.Fields{
            "request_id": c.GetHeader("X-Request-ID"),
            "method":     c.Request.Method,
            "path":       c.Request.URL.Path,
            "client_ip":  c.ClientIP(),
            "user_agent": c.Request.UserAgent(),
            "user_id":    c.GetString("user_id"), // Assuming user_id is set in the middleware
        }).Info("Incoming request")

        c.Next()

        // Log response details
        log.WithFields(log.Fields{
            "request_id": c.GetHeader("X-Request-ID"),
            "status":     c.Writer.Status(),
            "duration":   time.Since(start).Seconds(),
        }).Info("Request processed")
    }
   
}
