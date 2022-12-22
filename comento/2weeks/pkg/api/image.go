package api

import (
	"fmt"
	"main/internal/global"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RequestURI struct {
	ImageId string `uri:"id"`
}

type RequestFile struct {
	File *multipart.FileHeader `form:"file"`
}

type Response struct {
	Res string `json:"res"`
}

func Createimage(c *gin.Context) {
	req := RequestFile{}
	if err := c.ShouldBind(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "file is not valid"})
		return
	}

	if err := c.SaveUploadedFile(req.File, "images/"+strconv.FormatInt(global.MaxImageNumber, 10)); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "file is not valid"})
		return
	}
	global.MaxImageNumber++
	c.JSON(http.StatusOK, Response{Res: "success"})

}

func ReadImage(c *gin.Context) {
	reqURI := RequestURI{}
	if err := c.ShouldBindUri(&reqURI); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "resquest is not valid"})
		return
	}
	c.File("images/" + reqURI.ImageId)
}
