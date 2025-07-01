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
    auth := r.Group("/")
    auth.Use(middlewares.JWTAuthMiddleware())
    auth.POST("/products", controllers.CreateProduct)
}
