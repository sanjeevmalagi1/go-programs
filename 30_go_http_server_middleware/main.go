package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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

func uploadFile(c *gin.Context) {
	fmt.Println(c)

	c.IndentedJSON(http.StatusOK, []byte(""))
}

func downloadFile(c *gin.Context) {
	fmt.Println(c.Params)
	ps := c.Params

	filename, found := ps.Get("file_name")

	if found == false {
		fmt.Println("file_name not found")
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
		fmt.Println(filepath)
		c.FileAttachment(filepath, filename)
		return
	}

	c.IndentedJSON(http.StatusOK, []byte(""))
}

func main() {

	router := gin.Default()

	router.GET("api/v1/download/:file_name", downloadFile)
	router.POST("api/v1/upload", uploadFile)

	router.Run("localhost:3000")
}
