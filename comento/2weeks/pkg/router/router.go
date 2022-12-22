package router

import (
	"main/pkg/api"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.POST("/images", api.Createimage)
	r.GET("/images/:id", api.ReadImage)
	return r
}
