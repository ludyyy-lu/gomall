package routers
import (
    "github.com/gin-gonic/gin"
    "gomall/controllers"
)

func RegisterRoutes(r *gin.Engine) {
    r.POST("/register", controllers.Register)
}