package middlewares

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func LoggerMiddleware() fiber.Handler {
    log.SetFormatter(&log.JSONFormatter{})

    return func(c *fiber.Ctx) error {
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
            "request_id": c.Get("X-Request-ID"),
            "method":     c.Method(),
            "path":       c.Path(),
            "client_ip":  c.IP(),
            "user_agent": c.Get("User-Agent"),
            "user_id":    c.Locals("user_id"), // Assuming user_id is set in the middleware
        }).Info("Incoming request")

        err = c.Next()

        // Log response details
        log.WithFields(log.Fields{
            "request_id": c.Get("X-Request-ID"),
            "status":     c.Response().StatusCode(),
            "duration":   time.Since(start).Seconds(),
        }).Info("Request processed")

        return err
    }
}
