package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtkey = []byte("secretkey")

func Authmiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Autherization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid"})
			c.Abort()
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "bearer")
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return jwtkey, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid"})
			c.Abort()
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid"})
			c.Abort()
			return
		}
		c.Set("user_id", uint(claims["user_id"].(float64)))
		c.Next()
	}
}
