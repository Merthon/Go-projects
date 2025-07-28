package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ShowIndex(c *gin.Context)  {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
    "Title": "我的 Gin 示例首页",
    "Now":   time.Now().Format("2006-01-02 15:04:05"),
    "User":  "Merthon Chen",
})
}