package routers

import (
	"gomall/controllers"
	"gomall/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB, rdb *redis.Client) {
	// 用户认证
	userController := controllers.NewUserController(db)
	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)
	auth := r.Group("/")
	// 商品公共接口
	productController := controllers.NewProductController(db)
	product := auth.Group("/products")
	r.GET("/products", productController.GetProducts)
	r.GET("/products/:id", productController.GetProductDetail)

	// 需要登陆的接口
	auth.Use(middlewares.JWTAuthMiddleware())
	// 商品管理
	// 商品模块
	
	{
		product.POST("", productController.CreateProduct)
		product.PUT("/:id", productController.UpdateProduct)
		product.DELETE("/:id", productController.DeleteProduct)
		product.PATCH("/:id/status", productController.UpdateProductStatus)
		// 分类
		product.POST("/:id/categories", productController.SetProductCategories)
		product.GET("/:id/categories", productController.GetProductCategories)
		product.DELETE("/:id/categories/:category_id", productController.RemoveProductCategory)
	}

	// 分类模块
	categoryController := controllers.NewCategoryController(db)
	// 分类公共接口
	r.GET("/categories", categoryController.GetCategories)
	category := auth.Group("/categories")
	{
		category.POST("", categoryController.CreateCategory)
	}

	// 购物车 需要鉴权
	cartController := controllers.NewCartController(db)
	cart := auth.Group("/cart")
	{
		cart.POST("", cartController.AddToCart)
		cart.GET("", cartController.GetCartItems)
		cart.DELETE("/:id", cartController.DeleteCartItem)
		cart.PATCH("/:id", cartController.UpdateCartItem)

	}
	orderController := controllers.NewOrderController(db, rdb)
	order := auth.Group("/orders")
	{
		order.POST("", orderController.CreateOrder)
		order.GET("", orderController.GetOrders) // 查询当前用户的订单列表
		order.GET("/stats", orderController.GetOrderStats)
		order.GET("/:id", orderController.GetOrderDetail)        // 查看订单详情
		order.POST("/:id/pay", orderController.PayOrder)         // 模拟支付
		order.POST("/:id/cancel", orderController.CancelOrder)   //取消订单
		order.POST("/:id/ship", orderController.ShipOrder)       // 发货
		order.POST("/:id/confirm", orderController.ConfirmOrder) //确认收货
		order.GET("/auto-cancel", orderController.AutoCancelOrders)
	}
}
