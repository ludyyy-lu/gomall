package controllers

import (
    "net/http"
    "gomall/config"
    "gomall/models"
    "github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
    var input struct {
        Name        string  `json:"name" binding:"required"`
        Description string  `json:"description"`
        Price       float64 `json:"price" binding:"required"`
        Stock       int     `json:"stock"`
        ImageURL    string  `json:"image_url"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
        return
    }

    userID := c.GetUint("user_id") // 从JWT中拿用户ID

    product := models.Product{
        Name:        input.Name,
        Description: input.Description,
        Price:       input.Price,
        Stock:       input.Stock,
        ImageURL:    input.ImageURL,
        UserID:      userID,
    }

    if err := config.DB.Create(&product).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "创建失败"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "创建成功", "product": product})
}
