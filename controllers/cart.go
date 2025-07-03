package controllers

import (
	"gomall/models"
	"gomall/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CartController struct {
	DB *gorm.DB
	// RDB *redis.Client
}

func NewCartController(db *gorm.DB) *CartController {
	return &CartController{DB: db}
}

type AddCartItemInput struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  uint `json:"quantity" binding:"required"`
}

// 添加商品到购物车
// POST /cart
func (cc *CartController) AddToCart(c *gin.Context) {
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
	if err := cc.DB.First(&product, input.ProductID).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "商品不存在")
		return
	}

	if !product.OnSale || product.Stock == 0 {
		utils.Error(c, http.StatusBadRequest, "商品已下架或库存不足")
		return
	}

	// 查找是否已经存在相同商品的购物车项
	var cartItem models.CartItem
	err := cc.DB.
		Where("user_id = ? AND product_id = ?", userID, input.ProductID).
		First(&cartItem).Error

	if err == nil {
		// 已存在，增加数量
		cartItem.Quantity += input.Quantity
		cc.DB.Save(&cartItem)
	} else {
		// 不存在，新增购物车项
		cartItem = models.CartItem{
			UserID:    userID,
			ProductID: input.ProductID,
			Quantity:  input.Quantity,
		}
		cc.DB.Create(&cartItem)
	}

	utils.Success(c, gin.H{"cart_item": cartItem}, "商品成功添加至购物车")
}

// GET /cart
func (cc *CartController) GetCartItems(c *gin.Context) {
	userID := c.GetUint("user_id")

	var cartItems []models.CartItem
	if err := cc.DB.
		Where("user_id = ?", userID).
		Preload("Product").
		Find(&cartItems).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "获取购物车列表失败")
		return
	}

	utils.Success(c, gin.H{"items": cartItems}, "获取购物车列表成功")
}

// 删除购物车项
// DELETE /cart/:id
func (cc *CartController) DeleteCartItem(c *gin.Context) {
	userID := c.GetUint("user_id")
	cartItemID := c.Param("id")

	var cartItem models.CartItem
	if err := cc.DB.First(&cartItem, cartItemID).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "购物项不存在")
		return
	}

	if cartItem.UserID != userID {
		utils.Error(c, http.StatusForbidden, "无权限操作")
		return
	}

	if err := cc.DB.Delete(&cartItem).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "删除购物车中选中物品失败")
		return
	}

	utils.Success(c, nil, "删除购物车中选中物品成功")
}

type UpdateCartItemInput struct {
	Quantity uint `json:"quantity" binding:"required"`
}

// 修改数量
// PATCH /cart/:id
func (cc *CartController) UpdateCartItem(c *gin.Context) {
	userID := c.GetUint("user_id")
	cartItemID := c.Param("id")

	var input UpdateCartItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	if input.Quantity == 0 {
		utils.Error(c, http.StatusBadRequest, "数量必须大于 0")
		return
	}

	var cartItem models.CartItem
	if err := cc.DB.First(&cartItem, cartItemID).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "购物项不存在")
		return
	}

	if cartItem.UserID != userID {
		utils.Error(c, http.StatusForbidden, "无权限操作")
		return
	}

	cartItem.Quantity = input.Quantity
	if err := cc.DB.Save(&cartItem).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "更新购物车内选中商品数量失败")
		return
	}

	utils.Success(c, gin.H{"cart_item": cartItem}, "更新购物车内选中商品数量成功")
}
