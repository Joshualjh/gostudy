package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string `form:"name" json:"name"`
	Address string `form:"address" json:"address"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Println(file.Filename)
		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

	})

	r.GET("/test", func(c *gin.Context) {
		person := Person{}
		if err := c.ShouldBind(&person); err != nil {
			fmt.Println("Bind Err")
		}

		c.JSON(200, person)
	})

	r.Run()

}
