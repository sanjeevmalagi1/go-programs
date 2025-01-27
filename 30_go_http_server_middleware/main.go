package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func isFileExists(filepath string) (bool, error) {
	if _, err := os.Stat(filepath); err != nil {
		if os.IsNotExist(err) {
			return false, errors.New("File does not exist: " + filepath)
		}
		return false, errors.New("Access denied :" + filepath)
	}

	return true, nil
}

func downloadFile(c *gin.Context) {
	ps := c.Params

	filename, found := ps.Get("file_name")

	if found == false {
		c.IndentedJSON(http.StatusNoContent, []byte(""))
		return
	}

	filepath := "./" + filename
	fileExists, err := isFileExists(filename)
	if err != nil {
		c.IndentedJSON(http.StatusNoContent, []byte(""))
		return
	}

	if fileExists {
		c.FileAttachment(filepath, filename)
		return
	}

	c.IndentedJSON(http.StatusOK, []byte(""))
}

func uploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")

	file_name := uuid.Must(uuid.NewRandom()).String()

	if err := c.SaveUploadedFile(file, file_name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "File uploaded successfully",
		"file_name": file_name,
		"file_size": file.Size,
	})
}

func requestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Printf("%s\t %s\t %d\n", c.Request.Method, c.FullPath(), time.Unix)

	}
}

func main() {

	router := gin.Default()

	router.Use(requestLogger())

	router.GET("api/v1/download/:file_name", downloadFile)
	router.POST("api/v1/upload", uploadFile)

	router.Run("localhost:3000")
}
