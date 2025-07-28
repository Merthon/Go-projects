package controllers

import "github.com/gin-gonic/gin"

// LoginJSON 接收 JSON 体
func LoginJSON(c *gin.Context) {
    var json struct {
        User string `json:"user" binding:"required"`
        Pass string `json:"pass" binding:"required"`
    }
    if err := c.ShouldBindJSON(&json); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"user": json.User, "pass": json.Pass})
}