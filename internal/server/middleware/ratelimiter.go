package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// RateLimiter struct
type RateLimiter struct {
	clients map[string]*rate.Limiter
	mu      sync.Mutex
	r       rate.Limit
	b       int
}

// NewRateLimiter - Creates a new rate limiter
func NewRateLimiter(r rate.Limit, b int) *RateLimiter {
	return &RateLimiter{
		clients: make(map[string]*rate.Limiter),
		r:       r,
		b:       b,
	}
}

// getLimiter - Returns a rate limiter for an IP
func (rl *RateLimiter) getLimiter(ip string) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if limiter, exists := rl.clients[ip]; exists {
		return limiter
	}

	limiter := rate.NewLimiter(rl.r, rl.b)
	rl.clients[ip] = limiter

	// Cleanup expired limiters
	go func() {
		time.Sleep(time.Minute)
		rl.mu.Lock()
		delete(rl.clients, ip)
		rl.mu.Unlock()
	}()

	return limiter
}

// RateLimitMiddleware - Gin Middleware for rate limiting
func RateLimitMiddleware(rl *RateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter := rl.getLimiter(ip)

		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests. Please try again later."})
			c.Abort()
			return
		}

		c.Next()
	}
}
