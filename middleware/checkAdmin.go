package middleware

import "github.com/gin-gonic/gin"

func CheckAdmin() gin.HandlerFunc {

	return func(c *gin.Context) {
		if c.GetHeader("role") != "admin" {
			c.JSON(401, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
