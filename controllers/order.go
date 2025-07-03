package controllers

import (
	"context"
	"gomall/models"
	"gomall/utils"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type OrderController struct {
	DB  *gorm.DB
	RDB *redis.Client
}

func NewOrderController(db *gorm.DB, rdb *redis.Client) *OrderController {
	return &OrderController{DB: db, RDB: rdb}
}

// 创建订单
func (oc *OrderController) CreateOrder(c *gin.Context) {
	userID := c.GetUint("user_id")

	var cartItems []models.CartItem
	// 查询用户购物车中所有项（你也可以只下单选中的）
	if err := oc.DB.Where("user_id = ?", userID).Preload("Product").Find(&cartItems).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "获取购物车失败")
		return
	}

	if len(cartItems) == 0 {
		utils.Error(c, http.StatusBadRequest, "购物车为空")
		return
	}

	var orderItems []models.OrderItem
	var totalPrice float64

	// 开启事务
	tx := oc.DB.Begin()

	// 遍历购物车项，构造订单项
	for _, item := range cartItems {
		product := item.Product

		if !product.OnSale || product.Stock < item.Quantity {
			utils.Error(c, http.StatusBadRequest, "商品已下架或库存不足："+product.Name)
			return
		}

		// 乐观锁方式更新库存
		result := tx.Model(&models.Product{}).
			Where("id = ? AND version = ? AND stock >= ?", product.ID, product.Version, item.Quantity).
			Updates(map[string]interface{}{
				"stock":   gorm.Expr("stock - ?", item.Quantity),
				"version": gorm.Expr("version + 1"),
			})

		if result.RowsAffected == 0 {
			tx.Rollback()
			utils.Error(c, http.StatusConflict, "库存不足，或发生并发冲突："+product.Name)
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

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		utils.Error(c, http.StatusInternalServerError, "创建订单失败")
		return
	}
	//清空购物车项失败回滚
	if err := tx.Where("user_id = ?", userID).Delete(&models.CartItem{}).Error; err != nil {
		tx.Rollback()
		utils.Error(c, http.StatusInternalServerError, "清空购物车失败")
		return
	}

	// 提交事务
	tx.Commit()

	// 事务提交之后才设置超时
	// 这是为了保证「只有数据库里真实存在的订单」才会进入 Redis 延迟队列

	// 设置订单超时
	expireAt := time.Now().Add(1 * time.Minute).Unix()
	redisKey := "order:timeout"

	if err := oc.RDB.ZAdd(context.Background(), redisKey, redis.Z{
		Score:  float64(expireAt),
		Member: order.ID,
	}).Err(); err != nil {
		log.Printf("设置订单超时失败：%v", err)
		// 不中断整个流程，但是要记录下错误
	}

	utils.Success(c, gin.H{
		"order": order,
	}, "订单创建成功")

}

// 查询订单列表
func (oc *OrderController) GetOrders(c *gin.Context) {
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
	if err := oc.DB.
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
	oc.DB.Model(&models.Order{}).Where("user_id = ?", userID).Count(&total)

	utils.Success(c, gin.H{
		"page":  page,
		"size":  size,
		"data":  orders,
		"total": total,
	}, "订单获取成功")
}

// 获取订单详情
func (oc *OrderController) GetOrderDetail(c *gin.Context) {
	orderID := c.Param("id")
	userID := c.GetUint("user_id")

	var order models.Order
	if err := oc.DB.
		Where("id = ? AND user_id = ?", orderID, userID).
		Preload("OrderItems.Product").
		First(&order).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "订单不存在")
		return
	}

	utils.Success(c, gin.H{"order": order}, "订单详情获取成功")
}

// 模拟支付
func (oc *OrderController) PayOrder(c *gin.Context) {
	orderID := c.Param("id")
	userID := c.GetUint("user_id")

	var order models.Order
	if err := oc.DB.
		Preload("OrderItems.Product").
		Where("id = ? AND user_id = ?", orderID, userID).
		First(&order).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "订单不存在")
		return
	}

	if order.Status != "pending" {
		utils.Error(c, http.StatusBadRequest, "订单已失效，无法支付")
		return
	}

	// 开启事务
	// ！！！
	tx := oc.DB.Begin()

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

// 订单状态管理
// GetOrderStats 获取订单状态统计
func (oc *OrderController) GetOrderStats(c *gin.Context) {
	userID := c.GetUint("user_id")

	var total, pending, paid, cancelled, today int64
	// 全部订单数量
	// 待支付
	// 已支付
	// 已取消
	// 今日创建的订单数量

	// 总订单数
	oc.DB.Model(&models.Order{}).
		Where("user_id = ?", userID).
		Count(&total)

	// 待支付
	oc.DB.Model(&models.Order{}).
		Where("user_id = ? AND status = ?", userID, "pending").
		Count(&pending)

	// 已支付
	oc.DB.Model(&models.Order{}).
		Where("user_id = ? AND status = ?", userID, "paid").
		Count(&paid)

	// 已取消
	oc.DB.Model(&models.Order{}).
		Where("user_id = ? AND status = ?", userID, "cancelled").
		Count(&cancelled)

	// 今日订单（按创建时间统计）
	todayStart := time.Now().Truncate(24 * time.Hour)
	// time.Now()是获取当前时间 比如现在2025-07-02 15:23:45
	// time.Now().Truncate(24 * time.Hour)是获取今天的开始时间，即2025-07-02 00:00:00
	oc.DB.Model(&models.Order{}).
		Where("user_id = ? AND created_at >= ?", userID, todayStart).
		Count(&today)
	// 获取今天创建的订单

	utils.Success(c, gin.H{
		"total":     total,
		"pending":   pending,
		"paid":      paid,
		"cancelled": cancelled,
		"today":     today,
	}, "订单状态统计成功")
}

// 订单取消
// 只能取消状态是 pending（待支付）的订单
// 取消后订单状态更新为 canceled
// 恢复对应商品的库存（注意并发问题，后面可以用事务）
// 返回取消成功或失败信息
func (oc *OrderController) CancelOrder(c *gin.Context) {
	userID := c.GetUint("user_id")
	orderID := c.Param("id")

	var order models.Order
	if err := oc.DB.Where("id = ? AND user_id = ?", orderID, userID).First(&order).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "订单不存在")
		return
	}

	if order.Status != "pending" {
		utils.Error(c, http.StatusBadRequest, "订单无法取消，状态非待支付")
		return
	}

	// 启动事务
	err := oc.DB.Transaction(func(tx *gorm.DB) error {
		// 更新订单状态为取消
		if err := tx.Model(&order).Update("status", "canceled").Error; err != nil {
			return err
		}

		// 恢复库存
		var orderItems []models.OrderItem
		if err := tx.Where("order_id = ?", order.ID).Find(&orderItems).Error; err != nil {
			return err
		}
		for _, item := range orderItems {
			if err := tx.Model(&models.Product{}).
				Where("id = ?", item.ProductID).
				Update("stock", gorm.Expr("stock + ?", item.Quantity)).Error; err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "取消订单失败")
		return
	}

	utils.Success(c, nil, "订单取消成功")
}

// 订单状态流转
// 商家发货
// 发货：POST /orders/:id/ship
func (oc *OrderController) ShipOrder(c *gin.Context) {
	orderID := c.Param("id")
	userID := c.GetUint("user_id")

	var order models.Order
	if err := oc.DB.Where("id = ? AND user_id = ?", orderID, userID).First(&order).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "订单不存在")
		return
	}

	if order.Status != "paid" {
		utils.Error(c, http.StatusBadRequest, "订单未付款，不能发货")
		return
	}

	order.Status = "shipped"
	if err := oc.DB.Save(&order).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "发货失败")
		return
	}

	utils.Success(c, gin.H{"order": order}, "发货成功")
}

// 确认收货
// 确认收货：POST /orders/:id/confirm
func (oc *OrderController) ConfirmOrder(c *gin.Context) {
	orderID := c.Param("id")
	userID := c.GetUint("user_id")

	var order models.Order
	if err := oc.DB.Where("id = ? AND user_id = ?", orderID, userID).First(&order).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "订单不存在")
		return
	}

	if order.Status != "shipped" {
		utils.Error(c, http.StatusBadRequest, "订单未发货，无法确认收货")
		return
	}

	order.Status = "delivered"
	if err := oc.DB.Save(&order).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "确认收货失败")
		return
	}

	utils.Success(c, gin.H{"order": order}, "确认收货成功")
}

// 超时自动取消
// 超时关闭未支付订单：GET /orders/auto-cancel
func (oc *OrderController) AutoCancelOrders(c *gin.Context) {
	now := time.Now()
	tenMinutesAgo := now.Add(-10 * time.Minute)

	var orders []models.Order
	err := oc.DB.
		Where("status = ? AND created_at <= ?", "pending", tenMinutesAgo).
		Find(&orders).Error

	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "查询失败")
		return
	}

	// 取消这些订单
	for _, order := range orders {
		order.Status = "timeout"
		oc.DB.Save(&order)
	}

	utils.Success(c, gin.H{
		"cancelled_count": len(orders),
	}, "已自动取消超时订单")
}

// 订单取消、超时关闭（可选）
