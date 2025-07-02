package controllers

import (
	"gomall/config"
	"gomall/models"
	"gomall/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 创建商品
func CreateProduct(c *gin.Context) {
	var input struct {
		Name        string  `json:"name" binding:"required"`
		Description string  `json:"description"`
		Price       float64 `json:"price" binding:"required"`
		Stock       uint    `json:"stock"`
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

// 商品列表
func GetProducts(c *gin.Context) {
	// 分页参数
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

	// 搜索关键词
	keyword := c.Query("keyword")
	sort := c.DefaultQuery("sort", "created_at_desc")

	orderStr := "created_at DESC"
	switch sort {
	case "price_asc":
		orderStr = "price ASC"
	case "price_desc":
		orderStr = "price DESC"
	case "created_at_asc":
		orderStr = "created_at ASC"
	case "created_at_desc":
		orderStr = "created_at DESC"
	}

	var products []models.Product
	query := config.DB.Model(&models.Product{}).Where("on_sale = ?", true)
	if keyword != "" {
		query = query.Where("name LIKE ? OR description LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}

	var total int64
	query.Count(&total)

	if err := query.Order(orderStr).Limit(size).Offset(offset).Find(&products).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "获取商品列表失败")
		return
	}

	utils.Success(c, gin.H{
		"page":     page,
		"size":     size,
		"total":    total,
		"products": products,
	}, "获取成功")
}

// 商品详情
func GetProductDetail(c *gin.Context) {
	id := c.Param("id")

	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "商品不存在")
		return
	}

	utils.Success(c, gin.H{
		"product": product,
	}, "查询成功")
}

// 商品更新
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "商品不存在")
		return
	}

	var input struct {
		Name        *string  `json:"name"`
		Description *string  `json:"description"`
		Price       *float64 `json:"price"`
		Stock       *uint    `json:"stock"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, "请求参数错误")
		return
	}

	if input.Name != nil {
		product.Name = *input.Name
	}
	if input.Description != nil {
		product.Description = *input.Description
	}
	if input.Price != nil {
		product.Price = *input.Price
	}
	if input.Stock != nil {
		product.Stock = *input.Stock
	}

	if err := config.DB.Save(&product).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "更新失败")
		return
	}

	utils.Success(c, gin.H{"product": product}, "更新成功")
}

// 软删除商品
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "商品不存在")
		return
	}

	// 删除前解除关联
	config.DB.Model(&product).Association("Categories").Clear()

	if err := config.DB.Delete(&product).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "删除失败")
		return
	}

	utils.Success(c, nil, "删除成功")
}

// 给商品创建分类
type ProductCategoryInput struct {
	CategoryIDs []uint `json:"category_ids" binding:"required"`
}

// POST /products/:id/categories
func SetProductCategories(c *gin.Context) {
	var input ProductCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	productID := c.Param("id")

	var product models.Product
	if err := config.DB.Preload("Categories").First(&product, productID).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "商品不存在")
		return
	}

	var categories []models.Category
	if err := config.DB.Where("id IN ?", input.CategoryIDs).Find(&categories).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "查询分类失败")
		return
	}

	if err := config.DB.Model(&product).Association("Categories").Replace(&categories); err != nil {
		utils.Error(c, http.StatusInternalServerError, "关联分类失败")
		return
	}

	utils.Success(c, gin.H{"categories": categories}, "设置成功")
}

// 查询商品分类接口
// GET /products/:id/categories
func GetProductCategories(c *gin.Context) {
	productID := c.Param("id")

	var product models.Product
	if err := config.DB.Preload("Categories").First(&product, productID).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "商品不存在")
		return
	}

	utils.Success(c, gin.H{"categories": product.Categories}, "查询成功")
}

// 删除商品和某个分类的关联接口
func RemoveProductCategory(c *gin.Context) {
	productID := c.Param("product_id")
	categoryID := c.Param("category_id")

	var product models.Product
	if err := config.DB.Preload("Categories").First(&product, productID).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "商品不存在")
		return
	}

	var category models.Category
	if err := config.DB.First(&category, categoryID).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "分类不存在")
		return
	}

	if err := config.DB.Model(&product).Association("Categories").Delete(&category); err != nil {
		utils.Error(c, http.StatusInternalServerError, "解绑失败")
		return
	}

	utils.Success(c, nil, "解绑成功")
}

// PATCH /products/:id/status
func UpdateProductStatus(c *gin.Context) {
	productID := c.Param("id")
	var input struct {
		OnSale bool `json:"on_sale"` // 是否上架
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	var product models.Product
	if err := config.DB.First(&product, productID).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "商品不存在")
		return
	}

	product.OnSale = input.OnSale
	if err := config.DB.Save(&product).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "更新失败")
		return
	}

	utils.Success(c, gin.H{
		"product": product,
	}, "商品状态已更新")
}
