package controllers

import (
	"net/http"
	"gomall/config"
	"gomall/models"
	"gomall/utils"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	userID := c.GetUint("user_id")

	var cartItems []models.CartItem
	// 查询用户购物车中所有项（你也可以只下单选中的）
	if err := config.DB.Where("user_id = ?", userID).Preload("Product").Find(&cartItems).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "获取购物车失败")
		return
	}

	if len(cartItems) == 0 {
		utils.Error(c, http.StatusBadRequest, "购物车为空")
		return
	}

	var orderItems []models.OrderItem
	var totalPrice float64

	// 遍历购物车项，构造订单项
	for _, item := range cartItems {
		product := item.Product

		if !product.OnSale || product.Stock < item.Quantity {
			utils.Error(c, http.StatusBadRequest, "商品已下架或库存不足："+product.Name)
			return
		}

		// 更新库存
		product.Stock -= item.Quantity
		if err := config.DB.Save(&product).Error; err != nil {
			utils.Error(c, http.StatusInternalServerError, "更新库存失败")
			return
		}

		subTotal := float64(item.Quantity) * product.Price
		totalPrice += subTotal

		orderItems = append(orderItems, models.OrderItem{
			ProductID:  product.ID,
			Quantity:   item.Quantity,
			UnitPrice:  product.Price,
			TotalPrice: subTotal,
		})
	}

	// 创建订单
	order := models.Order{
		UserID:     userID,
		TotalPrice: totalPrice,
		OrderItems: orderItems,
	}

	if err := config.DB.Create(&order).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "创建订单失败")
		return
	}

	// 清空购物车
	config.DB.Where("user_id = ?", userID).Delete(&models.CartItem{})

	utils.Success(c, gin.H{
		"order": order,
	}, "订单创建成功")
}
