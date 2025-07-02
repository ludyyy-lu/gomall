package controllers

import (
	"gomall/config"
	"gomall/models"
	"gomall/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 创建订单
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

// 查询订单列表
func GetOrders(c *gin.Context) {
	userID := c.GetUint("user_id")

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

	var orders []models.Order
	// ！！！
	if err := config.DB.
		Where("user_id = ?", userID).
		Preload("OrderItems.Product").
		Order("created_at DESC").
		Limit(size).
		Offset(offset).
		Find(&orders).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "订单查询失败")
		return
	}

	var total int64
	config.DB.Model(&models.Order{}).Where("user_id = ?", userID).Count(&total)

	utils.Success(c, gin.H{
		"page":  page,
		"size":  size,
		"data":  orders,
		"total": total,
	}, "订单获取成功")
}

// 获取订单详情
func GetOrderDetail(c *gin.Context) {
	orderID := c.Param("id")
	userID := c.GetUint("user_id")

	var order models.Order
	if err := config.DB.
		Where("id = ? AND user_id = ?", orderID, userID).
		Preload("OrderItems.Product").
		First(&order).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "订单不存在")
		return
	}

	utils.Success(c, gin.H{"order": order}, "订单详情获取成功")
}

// 模拟支付
func PayOrder(c *gin.Context) {
	orderID := c.Param("id")
	userID := c.GetUint("user_id")

	var order models.Order
	if err := config.DB.
		Preload("OrderItems.Product").
		Where("id = ? AND user_id = ?", orderID, userID).
		First(&order).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "订单不存在")
		return
	}

	if order.Status != "pending" {
		utils.Error(c, http.StatusBadRequest, "订单无法支付")
		return
	}

	// 开启事务
	// ！！！
	tx := config.DB.Begin()

	for _, item := range order.OrderItems {
		if item.Product.Stock < item.Quantity {
			tx.Rollback()
			utils.Error(c, http.StatusBadRequest, "商品库存不足："+item.Product.Name)
			return
		}
		// 扣库存
		if err := tx.Model(&item.Product).Where("id = ?", item.Product.ID).
			Update("stock", gorm.Expr("stock - ?", item.Quantity)).Error; err != nil {
			tx.Rollback()
			utils.Error(c, http.StatusInternalServerError, "库存扣除失败")
			return
		}
	}

	// 修改订单状态
	if err := tx.Model(&order).Update("status", "paid").Error; err != nil {
		tx.Rollback()
		utils.Error(c, http.StatusInternalServerError, "更新订单状态失败")
		return
	}

	tx.Commit()

	utils.Success(c, gin.H{
		"order_id": order.ID,
		"status":   "paid",
	}, "订单支付成功")
}

// 购物车状态管理
// 库存扣减和回滚 ？
// 订单取消、超时关闭（可选）
