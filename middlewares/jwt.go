package middlewares

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "未认证"})
            c.Abort()
            return
        }

        tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
        token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
			// 确保签名算法是预期的，增加安全性
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, errors.New("非预期的签名算法")
            }
            return []byte(os.Getenv("JWT_SECRET")), nil
        })

		 // 1. 首先，检查解析过程中是否发生错误
        if err != nil {
            // 根据不同的错误类型返回不同的信息
            if errors.Is(err, jwt.ErrTokenExpired) {
                c.JSON(http.StatusUnauthorized, gin.H{"error": "Token已过期"})
            } else {
                c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的Token"})
            }
            c.Abort()
            return
        }

		// 2. 然后，再检查Token的有效性和声明
        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            userID := uint(claims["user_id"].(float64))
            c.Set("user_id", userID) // 存到上下文
            c.Next()
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "认证失败"})
            c.Abort()
        }
    }
}
