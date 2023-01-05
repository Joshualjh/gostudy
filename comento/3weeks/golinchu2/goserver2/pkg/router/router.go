package router

import (
	"github.com/gin-gonic/gin"
	"main.go/pkg/api"
)

func Router(apis *api.APIs) *gin.Engine {
	r := gin.Default()
	r.GET("/students/", apis.GetStudentlist)
	r.POST("/students/create-student", apis.CreateStudent)
	r.GET("/students/edit-student/:rollno", apis.GetStudentlistByrollID)
	r.PUT("/students/update-student/:rollno", apis.UpdateStudentList)
	r.DELETE("/students/delete-student/:rollno", apis.DeleteStudent)
	return r
}
