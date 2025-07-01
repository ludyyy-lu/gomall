package controllers

import (
	"gomall/config"
	"gomall/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {
	userID := c.GetUint("user_id")
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}
