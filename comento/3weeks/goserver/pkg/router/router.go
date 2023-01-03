package router

import (
	"main/pkg/api"

	"github.com/gin-gonic/gin"
)

func Router(apis *api.APIs) *gin.Engine {
	r := gin.Default()

	r.GET("/crud", apis.GetBoradList)
	r.POST("/crud", apis.CreateBoard)
	r.DELETE("/crud/:id", apis.DeleteBoard)
	r.PUT("/crud", apis.UpdateBoard)

	return r
}
