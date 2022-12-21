package router

import "github.com/gin-gonic/gin"

func Createrouter(c *gin.Context) {
	r := gin.Default()

	r.Run()
}
