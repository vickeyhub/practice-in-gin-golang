package controllers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadPhoto(c *gin.Context) {
	// create the uploads directory if it doesn't exist
	os.MkdirAll("./uploads", os.ModePerm)

	file, err := c.FormFile("photo")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "No file uploaded",
		})
		return
	}

	dst, err := os.Create(filepath.Join("./uploads", file.Filename))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   "Unable to create file",
		})
		return
	}

	defer dst.Close()

	// save the file to the server
	if err := c.SaveUploadedFile(file, dst.Name()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   "Unable to save file",
		})
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "File uploaded successfully",
	// })
}
