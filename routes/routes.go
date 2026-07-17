package routes

import (
	"learn-go/controllers"
	"learn-go/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// router.METHOD(path, handler)
	router.GET("/", controllers.Home)
	router.GET("/about",
		middleware.CheckAPIKey(),
		controllers.About,
	)

	api := router.Group("/api")
	{
		user := api.Group("/user")
		{
			user.GET("/", controllers.GetAllUsers)
			user.POST("/", controllers.CreateUser)
			user.GET("/:id", controllers.GetUser)
			user.PUT("/:id", controllers.UpdateUser)
			user.DELETE("/:id", controllers.DeleteUser)
			user.GET("/search", controllers.Search)

		}
		// Product Route
		api.GET("/product", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"data": []string{"Apple", "Mango", "Banana"},
			})
		})
	}
}
