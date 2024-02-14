package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)


type ipLimiter struct {
    limiter *rate.Limiter
    lastSeen time.Time
}

var ipLimiters = make(map[string]*ipLimiter)

func ipLimiterMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        ip := c.ClientIP()
        if _, found := ipLimiters[ip]; !found {
            ipLimiters[ip] = &ipLimiter{
                limiter: rate.NewLimiter(2, 5), 
            }
        }

        limiter := ipLimiters[ip] 
        if limiter.limiter.Allow() == false {
            c.AbortWithStatus(429) // Too many requests
            return
        } 

        limiter.lastSeen = time.Now()  
        c.Next() 
    }
}