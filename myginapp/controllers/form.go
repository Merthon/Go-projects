package controllers

import "github.com/gin-gonic/gin"

// PostFormData 接收用户名和密码
func PostFormData(c *gin.Context) {
    user := c.PostForm("user")
    pass := c.PostForm("pass")
    c.JSON(200, gin.H{"user": user, "pass": pass})
}