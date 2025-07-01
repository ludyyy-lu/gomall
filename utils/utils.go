package utils

import "github.com/gin-gonic/gin"

// 成功响应
func Success(c *gin.Context, data gin.H, message string) {
	c.JSON(200, gin.H{
		"status":  "success",
		"message": message,
		"data":    data,
	})
}

// 错误响应
func Error(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"status":  "error",
		"message": message,
	})
}
