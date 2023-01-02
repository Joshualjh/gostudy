package router

import (
	"main/pkg/api"

	"github.com/gin-gonic/gin"
)

func Router(apis *api.APIs) *gin.Engine {
	r := gin.Default()
<<<<<<< HEAD
	r.GET(":3000/crud", apis.GetBoradList)
	r.POST(":3000/crud/:id", apis.CreateBoard)
	r.DELETE(":3000/crud/:id", apis.DeleteBoard)
	r.PUT(":3000/crud/:id", apis.UpdateBoard)
=======
	r.GET(":3000/crud")
>>>>>>> f0db2b31fd49a8f8497f1585cd02ade6e469717b

	return r
}
