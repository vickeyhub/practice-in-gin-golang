package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before Controller")

		c.Next()

		fmt.Println("After Controller")
	}
}
