package routers
import (
    "github.com/gin-gonic/gin"
    "gomall/controllers"
)
// 这样写对吗
func RegisterRoutes(r *gin.Engine) {
    r.POST("/register", controllers.Register)
}

func LoginRoutes(r *gin.Engine) {
    r.POST("/login", controllers.Login)
}

