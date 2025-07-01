# gomall
电商后端系统

# 疑问
.env 是做什么用的

routers写成这样对的嘛?
func RegisterRoutes(r *gin.Engine) {
    r.POST("/register", controllers.Register)
}

func LoginRoutes(r *gin.Engine) {
    r.POST("/login", controllers.Login)
}

不对。RegisterRoutes里的Register不是“用户注册”的意思，而是“注册所有路由到GIn的Engine上”的意思。
* 就像注册事件、注册插件一样，这里的“register”指的是“注册（挂载）路由”。