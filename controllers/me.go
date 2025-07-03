package controllers

import (
	"gomall/models"
	"gomall/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) Me(c *gin.Context) {
	userID := c.GetUint("user_id")

	var user models.User
	if err := uc.DB.First(&user, userID).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "用户不存在")
		return
	}

	utils.Success(c, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	}, "获取用户信息成功")
}
