package controllers

import (
	"gomall/config"
	"gomall/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register 注册用户
func Register(c *gin.Context) {
    var input struct {
        Username string `json:"username" binding:"required,min=3,max=20"`
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required,min=6"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
        return
    }

    // 加密密码
    hashedPwd, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "密码加密失败"})
        return
    }

    user := models.User{
        Username: input.Username,
        Email:    input.Email,
        Password: string(hashedPwd),
    }

    if err := config.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "用户创建失败（可能用户名/邮箱已存在）"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}
