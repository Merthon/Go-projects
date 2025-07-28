package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/Merthon/myginapp/models"
)

// ListUsers GET /api/users
func ListUsers(c *gin.Context) {
    users := models.GetAllUsers()
    c.JSON(http.StatusOK, users)
}

// GetUser GET /api/users/:id
func GetUser(c *gin.Context) {
    id := c.Param("id")
    user, err := models.GetUserByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, user)
}

// CreateUser POST /api/users
func CreateUser(c *gin.Context) {
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    user := models.CreateUser(input)
    c.JSON(http.StatusCreated, user)
}

// UpdateUser PUT /api/users/:id
func UpdateUser(c *gin.Context) {
    id := c.Param("id")
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    user, err := models.UpdateUser(id, input)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, user)
}

// DeleteUser DELETE /api/users/:id
func DeleteUser(c *gin.Context) {
    id := c.Param("id")
    if err := models.DeleteUser(id); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}
