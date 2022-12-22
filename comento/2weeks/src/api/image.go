package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var id int = 0

func Createimages(c *gin.Context) {
	if file, err := c.FormFile("file"); err != nil {
		fmt.Println(err)
	} else {
		c.SaveUploadedFile(file, file.Filename)
	}

}
