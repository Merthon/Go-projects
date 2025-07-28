package controllers

import "github.com/gin-gonic/gin"

func SearchItem(c *gin.Context)  {
	query := c.Query("q")           // 若无 ?q，则返回空字符串
    page  := c.DefaultQuery("page", "1") // 若无 ?page，则用 "1"

	c.JSON(200, gin.H{
		"query": query,
        "page":  page,
	})
}