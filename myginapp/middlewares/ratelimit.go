package middlewares

import (
    "github.com/gin-gonic/gin"
    "golang.org/x/time/rate"
    "net/http"
    "sync"

)

var visitors = make(map[string]*rate.Limiter)
var mu sync.Mutex

// getVisitor 返回对应 key（IP）的限流器
func getVisitor(key string) *rate.Limiter {
    mu.Lock()
    defer mu.Unlock()
    limiter, exists := visitors[key]
    if !exists {
        limiter = rate.NewLimiter(1, 5) // 每秒 1 个令牌，桶容量 5
        visitors[key] = limiter
    }
    return limiter
}

// RateLimitMiddleware 限流中间件
func RateLimitMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        key := c.ClientIP()
        limiter := getVisitor(key)
        // 等待最多 100ms 获取令牌
        if !limiter.Allow() {
            c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
                "error": "too many requests",
            })
            return
        }
        c.Next()
    }
}
