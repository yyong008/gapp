package middleware

import (
	"gapp1/pkg/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头提取 token
		token := jwt.ExtractToken(c.GetHeader("Authorization"))

		// 验证 token
		valid, err := jwt.VerifyToken(token)
		if err != nil || !valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// 设置用户信息
		claims, err := jwt.ParseJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Set("userID", claims.ID)
		c.Next()
	}
}
