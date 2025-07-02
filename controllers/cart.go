package controllers

import (
	"gomall/config"
	"gomall/models"
	"gomall/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddCartItemInput struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  uint `json:"quantity" binding:"required"`
}

// 添加购物车
// POST /cart
func AddToCart(c *gin.Context) {
	var input AddCartItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	if input.Quantity == 0 {
		utils.Error(c, http.StatusBadRequest, "购买数量必须大于 0")
		return
	}

	userID := c.GetUint("user_id")

	// 查询商品是否存在
	var product models.Product
	if err := config.DB.First(&product, input.ProductID).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "商品不存在")
		return
	}

	if !product.OnSale || product.Stock == 0 {
		utils.Error(c, http.StatusBadRequest, "商品已下架或库存不足")
		return
	}

	// 查找是否已经存在相同商品的购物车项
	var cartItem models.CartItem
	err := config.DB.
		Where("user_id = ? AND product_id = ?", userID, input.ProductID).
		First(&cartItem).Error

	if err == nil {
		// 已存在，增加数量
		cartItem.Quantity += input.Quantity
		config.DB.Save(&cartItem)
	} else {
		// 不存在，新增购物车项
		cartItem = models.CartItem{
			UserID:    userID,
			ProductID: input.ProductID,
			Quantity:  input.Quantity,
		}
		config.DB.Create(&cartItem)
	}

	utils.Success(c, gin.H{"cart_item": cartItem}, "添加成功")
}

// GET /cart
func GetCartItems(c *gin.Context) {
	userID := c.GetUint("user_id")

	var cartItems []models.CartItem
	if err := config.DB.
		Where("user_id = ?", userID).
		Preload("Product").
		Find(&cartItems).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "获取购物车失败")
		return
	}

	utils.Success(c, gin.H{"items": cartItems}, "获取成功")
}
