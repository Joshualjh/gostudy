package api

import (
	"main/pkg/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Res struct {
	Res string `json:"res"`
}

func (apis *APIs) CreateBoard(c *gin.Context) {
	req := &model.Board{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, &Res{Res: "not vaild"})
		return
	}
	res, err := apis.db.CreateBoard(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Res{Res: "not vaild1"})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (apis *APIs) GetBoradList(c *gin.Context) {
	res, err := apis.db.GetBoradList()
	if err != nil {
		c.JSON(http.StatusBadRequest, &Res{Res: "not vaild2"})
	}
	c.JSON(http.StatusOK, res)
}

func (apis *APIs) DeleteBoard(c *gin.Context) {
	id := c.Param("id")
	if err := apis.db.DeleteBoard(id); err != nil {
		c.JSON(http.StatusBadRequest, &Res{Res: "not vaild3"})
	}
}

func (apis *APIs) UpdateBoard(c *gin.Context) {
	ids := c.Param("id")
	req := &model.Board{}

	id, err := strconv.ParseUint(ids, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Res{Res: "not vaild4"})
	}

	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, Res{Res: "not vaild5"})
	}
	res, err := apis.db.UpdateBoard(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Res{Res: "not vaild6"})
	}
	c.JSON(http.StatusOK, res)

}
