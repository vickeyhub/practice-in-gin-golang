package controllers

import (
	"learn-go/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var users []models.User

func Home(c *gin.Context) {
	c.String(http.StatusOK, "Hello Gin!")
}

func About(c *gin.Context) {
	c.String(http.StatusOK, "About Page")
}

func GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func GetUser(c *gin.Context) {
	var user_id string = c.Param("id")

	for _, user := range users {
		if strconv.Itoa(user.ID) == user_id {
			c.JSON(http.StatusOK, gin.H{
				"statusCode": http.StatusOK,
				"data":       user,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "User not found",
	})
}

func Search(c *gin.Context) {
	var query string = c.Query("query")
	var optional string = c.Query("optional")
	c.String(http.StatusOK, "Search Query: %s, Optional: %s", query, optional)
}

func CreateUser(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	users = append(users, user)
	c.JSON(200, gin.H{
		"message": "User created successfully",
		"user":    user,
	})

}

func UpdateUser(c *gin.Context) {
	user_id := c.Param("id")

	var updatedUser models.User

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	for i, user := range users {
		if strconv.Itoa(user.ID) == user_id {
			users[i] = updatedUser
			c.JSON(http.StatusOK, gin.H{
				"statusCode": http.StatusOK,
				"message":    "User updated successfully",
				"data":       updatedUser,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "User not found",
	})
}

func DeleteUser(c *gin.Context) {
	user_id := c.Param("id")

	for i, user := range users {
		if strconv.Itoa(user.ID) == user_id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"statusCode": http.StatusOK,
				"message":    "User deleted successfully",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "User not found",
	})
}
