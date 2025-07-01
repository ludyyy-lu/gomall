package routers

import (
	"gomall/controllers"
	"gomall/middlewares"

	"github.com/gin-gonic/gin"
)

// 这样写对吗
func RegisterRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/products", controllers.GetProducts)
	r.GET("/products/:id", controllers.GetProductDetail)
	r.PUT("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)
    r.POST("/categories", controllers.CreateCategory)
    r.GET("/categories", controllers.GetCategories)
    r.POST("/products/:id/categories", controllers.SetProductCategories)
    r.GET("/products/:id/categories", controllers.GetProductCategories)
    r.DELETE("/products/:product_id/categories/:category_id", controllers.RemoveProductCategory)

	auth := r.Group("/")
	auth.Use(middlewares.JWTAuthMiddleware())
	auth.POST("/products", controllers.CreateProduct)
}
