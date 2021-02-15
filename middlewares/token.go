package middlewares

import (
	"net/http"
	"strings"

	"github.com/QMCHE/diary-server/utils"
	"github.com/gin-gonic/gin"
)

// VerifyToken is middleware for token validation
func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Request.Cookie("access-token")
		if err != nil || strings.Trim(tokenString.Value, " ") == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Token is required",
			})
			c.Abort()
			return
		}

		// if utils.IsExpired(tokenString.Value) {
		// 	c.JSON(http.StatusUnauthorized, gin.H{
		// 		"error": "Token is expired",
		// 	})
		// 	c.Abort()
		// 	return
		// }

		claims, err := utils.VerifyToken(tokenString.Value)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
