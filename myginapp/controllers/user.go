package controllers

import "github.com/gin-gonic/gin"

// GetUser 根据路径参数 id 返回用户信息
func GetUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"user_id": id})
}