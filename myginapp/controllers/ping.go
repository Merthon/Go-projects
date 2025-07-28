package controllers

import "github.com/gin-gonic/gin"

// Ping 处理 /ping 路由
func Ping(c *gin.Context) {
    c.JSON(200, gin.H{"message": "pong"})
}