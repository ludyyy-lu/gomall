package routers

import (
	"gomall/controllers"
	"gomall/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// 用户认证
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	// 商品公共接口
	r.GET("/products", controllers.GetProducts)
	r.GET("/products/:id", controllers.GetProductDetail)
	// 分类公共接口
	r.GET("/categories", controllers.GetCategories)

	// 需要登陆的接口
	auth := r.Group("/")
	auth.Use(middlewares.JWTAuthMiddleware())
	// 商品管理
	// 商品模块
	product := auth.Group("/products")
	{
		product.POST("", controllers.CreateProduct)
		product.PUT("/:id", controllers.UpdateProduct)
		product.DELETE("/:id", controllers.DeleteProduct)
		product.PATCH("/:id/status", controllers.UpdateProductStatus)
		// 分类
		product.POST("/:id/categories", controllers.SetProductCategories)
		product.GET("/:id/categories", controllers.GetProductCategories)
		product.DELETE("/:id/categories/:category_id", controllers.RemoveProductCategory)
	}

	// 分类模块
	category := auth.Group("/categories")
	{
		category.POST("", controllers.CreateCategory)
	}

	// 购物车 需要鉴权
	cart := auth.Group("/cart")
	{
		cart.POST("", controllers.AddToCart)
		cart.GET("", controllers.GetCartItems)
		cart.DELETE("/:id", controllers.DeleteCartItem)
		cart.PATCH("/:id", controllers.UpdateCartItem)

	}
	order := auth.Group("/orders")
	{
		order.POST("", controllers.CreateOrder)
		order.GET("", controllers.GetOrders)          // 查询当前用户的订单列表
		order.GET("/:id", controllers.GetOrderDetail) // 查看订单详情
		order.POST("/:id/pay", controllers.PayOrder)  // 模拟支付
	}
}
