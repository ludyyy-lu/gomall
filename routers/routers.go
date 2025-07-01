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

	auth := r.Group("/")
	auth.Use(middlewares.JWTAuthMiddleware())
	auth.POST("/products", controllers.CreateProduct)
}
