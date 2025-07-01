package main

import (
    "gomall/config"
    "gomall/routers"
    "github.com/gin-gonic/gin"
)

func main() {
    config.InitDB()

    r := gin.Default()
    routers.RegisterRoutes(r)

    r.Run(":8080") // 监听端口
}
