package router

import (
	"main/pkg/api"

	"github.com/gin-gonic/gin"
)

func Router(apis *api.APIs) *gin.Engine {
	r := gin.Default()
	r.GET(":3000/crud")

	return r
}
