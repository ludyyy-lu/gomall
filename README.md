# gomall
电商后端系统

# 疑问
.env 是做什么用的

routers写成这样对的嘛
func RegisterRoutes(r *gin.Engine) {
    r.POST("/register", controllers.Register)
}

func LoginRoutes(r *gin.Engine) {
    r.POST("/login", controllers.Login)
}
