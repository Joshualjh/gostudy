package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/pkg/model"
)

type Res struct {
	Res string `json:"res"`
}

func (apis *APIs) CreateStudent(c *gin.Context) {
	req := &model.Students{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, &Res{Res: "not vaild"})
		return
	}

	result, err := apis.db.CreateStudent(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Res{Res: "not vaild"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (apis *APIs) GetStudentlist(c *gin.Context) {
	res, err := apis.db.GetStudentlist()
	if err != nil {
		c.JSON(http.StatusBadRequest, &Res{Res: "not vaild"})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (apis *APIs) GetStudentlistByrollID(c *gin.Context) {
	rollno := c.Param("rollno")

	res, err := apis.db.GetStudentlistByrollID(rollno)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Res{Res: "not vaild"})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (apis *APIs) UpdateStudentList(c *gin.Context) {
	rollno := c.Param("rollno")
	req := &model.Students{}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, &Res{Res: "not vaild"})
		return
	}
	result, err := apis.db.UpdateStudentList(rollno, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Res{Res: "not vaild"})
		return
	}

	c.JSON(http.StatusOK, result)

}
func (apis *APIs) DeleteStudent(c *gin.Context) {
	rollno := c.Param("rollno")
	if err := apis.db.DeleteStudent(rollno); err != nil {
		c.JSON(http.StatusBadRequest, &Res{Res: "not vaild"})
		return
	}
	c.JSON(http.StatusOK, 200)
}
