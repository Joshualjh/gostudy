package api

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type RequestURI struct {
	ImageId string `uri:"id"`
}

type RequestFile struct {
	File    string `json:"image"`
	No      string `json:"no"`
	ImageId string `uri:"id"`
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
		c.JSON(http.StatusBadRequest, Response{Res: "file is not valid2"})
		return
	}
	if err := os.WriteFile("./images/"+req.No, data, 0755); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "file is not valid3"})
		return
	}
	c.JSON(http.StatusOK, Response{Res: "success"})

}

func ReadImage(c *gin.Context) {
	reqURI := RequestURI{}
	if err := c.ShouldBindUri(&reqURI); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "resquest is not valid"})
		return
	}

	c.File("./images/" + reqURI.ImageId)
}

func DeleteImage(c *gin.Context) {
	reqURI := RequestURI{}
	if err := c.ShouldBindUri(&reqURI); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "resquest is not valid"})
		return
	}
	if err := os.Remove("./images/" + reqURI.ImageId); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "there is no file. file name : " + reqURI.ImageId})
		return

	}
	c.JSON(http.StatusOK, Response{Res: "success"})

}

func UpdateImage(c *gin.Context) {
	req := RequestFile{}
	reqURI := RequestURI{}
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "file is not valid1"})
		return
	}
	if err := c.ShouldBindUri(&reqURI); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "resquest is not valid"})
		return
	}
	data, errBase := base64.RawStdEncoding.DecodeString(strings.Split(req.File, "base64,")[1])
	if errBase != nil {
		fmt.Println(errBase)
		c.JSON(http.StatusBadRequest, Response{Res: "file is not valid2"})
		return
	}
	if err := os.Remove("./images/" + reqURI.ImageId); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "there is no file. file name : " + reqURI.ImageId})
		return
	}
	if err := ioutil.WriteFile("./images/"+reqURI.ImageId, data, 0755); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "file is not valid3"})
		return
	}
	c.JSON(http.StatusOK, Response{Res: "success"})

}
