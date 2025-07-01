package controllers

import (
	"gomall/config"
	"gomall/models"
	"gomall/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 创建分类接口
type CategoryInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

func CreateCategory(c *gin.Context) {
	var input CategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	category := models.Category{
		Name:        input.Name,
		Description: input.Description,
	}

	if err := config.DB.Create(&category).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "创建失败")
		return
	}

	utils.Success(c, gin.H{"category": category}, "创建成功")
}

// 获取分类列表接口
func GetCategories(c *gin.Context) {
	var categories []models.Category
	if err := config.DB.Find(&categories).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "查询失败")
		return
	}
	utils.Success(c, gin.H{"categories": categories}, "查询成功")
}

// 查询分类下的商品
func GetProductsByCategory(c *gin.Context) {
	categoryID := c.Param("id")
	pageStr := c.DefaultQuery("page", "1")
	sizeStr := c.DefaultQuery("size", "10")

	page, _ := strconv.Atoi(pageStr)
	size, _ := strconv.Atoi(sizeStr)
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 10
	}
	offset := (page - 1) * size

	var category models.Category
	if err := config.DB.First(&category, categoryID).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "分类不存在")
		return
	}

	var products []models.Product
	if err := config.DB.Model(&category).
		Offset(offset).
		Limit(size).
		Association("Products").
		Find(&products); err != nil {
		utils.Error(c, http.StatusInternalServerError, "查询失败")
		return
	}

	utils.Success(c, gin.H{"products": products}, "查询成功")
}
