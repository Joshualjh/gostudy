package api

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"

	"strings"

	"github.com/gin-gonic/gin"
)

type RequestURI struct {
	ImageId string `image:"base64"`
	No      string `json:"no"`
}

type RequestFile struct {
	File string `json:"image"`
	No   string `json:"no"`
}

type Response struct {
	Res string `json:"res"`
}

func Createimage(c *gin.Context) {
	req := RequestFile{}

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "file is not valid1"})
		return
	}
	data, errBase := base64.RawStdEncoding.DecodeString(strings.Split(req.File, "base64,")[1])
	if errBase != nil {
		fmt.Println(errBase)
		c.JSON(http.StatusBadRequest, Response{Res: "file is not valid1"})
		return
	}

	file, errfile := c.FormFile(bytes.NewBuffer(data).String())
	if errfile != nil {
		fmt.Println(errfile)
		c.JSON(http.StatusBadRequest, Response{Res: "file is not valid2"})
		return
	}

	if err := c.SaveUploadedFile((file), "images/"+req.No); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "file is not valid3"})
		return
	}

	// global.MaxImageNumber++
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

func DeleteImage(c *gin.Context) {

}

func UpdateImage(c *gin.Context) {

}
