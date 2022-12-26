package router

import (
	"main/pkg/api"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	// r.Use(cors.New(
	// 	cors.Config{
	// 		AllowOrigins: []string{"http://localhost:8081"},
	// 		AllowMethods: []string{"GET,POST"},
	// 	}))
	r.Static("/request", "./public")
	r.POST("/images", api.Createimage)
	r.GET("/images/:id", api.ReadImage)
	r.PUT("/images/:id", api.UpdateImage)
	r.DELETE("images/:id", api.DeleteImage)
	return r
}
