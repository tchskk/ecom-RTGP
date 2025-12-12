package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tchskk/ecom-RTGP/backend/config"
	"github.com/tchskk/ecom-RTGP/backend/models"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}