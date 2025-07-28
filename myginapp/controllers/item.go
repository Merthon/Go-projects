package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// QueryInput 仅用于 GET：URI + Query 参数
type QueryInput struct {
    ID    string `uri:"id" binding:"required,uuid4"`
    Query string `form:"q" binding:"required,min=3,max=50"`
    Page  int    `form:"page" binding:"omitempty,min=1,max=100"`
}

// PostInput 仅用于 POST：JSON 体
type PostInput struct {
    Price    float64 `json:"price" binding:"required,gt=0"`
    Category string  `json:"category" binding:"required,oneof=A B C D"`
}

// GetItem 处理 GET /item/:id?q=...&page=...
func GetItem(c *gin.Context) {
    var q QueryInput
    if err := c.ShouldBindUri(&q); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := c.ShouldBindQuery(&q); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "id":    q.ID,
        "query": q.Query,
        "page":  q.Page,
    })
}

// PostItem 处理 POST /item/:id?q=...&page=...
func PostItem(c *gin.Context) {
    // 1. 绑定 URI + Query
    var q QueryInput
    if err := c.ShouldBindUri(&q); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := c.ShouldBindQuery(&q); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // 2. 绑定 JSON
    var p PostInput
    if err := c.ShouldBindJSON(&p); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // 3. 返回合并结果
    c.JSON(http.StatusOK, gin.H{
        "id":       q.ID,
        "query":    q.Query,
        "page":     q.Page,
        "price":    p.Price,
        "category": p.Category,
    })
}
