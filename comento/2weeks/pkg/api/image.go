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
	File string `json:"image"`
	No   string `json:"no"`
}

type Response struct {
	Res string `json:"res"`
}

func Createimage(c *gin.Context) {
	req := RequestFile{}

	if err := c.ShouldBindJSON(&req); err != nil { //먼저 json으로 바인딩 type = string
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "binding error"})
		return
	}
	if err := os.WriteFile("./images/"+req.No, []byte(req.File), 0755); err != nil { //파일생성
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "Do not create file"})
		return
	}
	data, errBase := base64.RawStdEncoding.DecodeString(strings.Split(req.File, "base64,")[1]) // :"base64, 부분까지 제거를 한 후 []byte로 변환"
	if errBase != nil {
		fmt.Println(errBase)
		c.JSON(http.StatusBadRequest, Response{Res: "byte conversion error"})
		return
	}
	if err := os.WriteFile("./images/"+req.No, data, 0755); err != nil { //파일생성
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "Do not create file"})
		return
	}
	c.JSON(http.StatusOK, Response{Res: "success to upload images"})

}

func ReadImage(c *gin.Context) {
	reqURI := RequestURI{}
	if err := c.ShouldBindUri(&reqURI); err != nil { //:id 바인딩
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "binding error"})
		return
	}

	c.File("./images/" + reqURI.ImageId) //파일 읽기
}

func DeleteImage(c *gin.Context) {
	reqURI := RequestURI{}
	if err := c.ShouldBindUri(&reqURI); err != nil { //:id 바인딩
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "binding error"})
		return
	}
	if err := os.Remove("./images/" + reqURI.ImageId); err != nil { //파일 삭제 (해당 path에 존재하지 않을경우 에러)
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "there is exist file. file name : " + reqURI.ImageId})
		return

	}
	c.JSON(http.StatusOK, Response{Res: "success to delete images"})

}

func UpdateImage(c *gin.Context) {
	req := RequestFile{}
	reqURI := RequestURI{}
	if err := c.ShouldBindJSON(&req); err != nil { //파일에 대한 바인딩
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "file binding error"})
		return
	}
	if err := c.ShouldBindUri(&reqURI); err != nil { //:id에 대한 바인딩
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "id binding error"})
		return
	}
	data, errBase := base64.RawStdEncoding.DecodeString(strings.Split(req.File, "base64,")[1]) //[]byte로 변환
	if errBase != nil {
		fmt.Println(errBase)
		c.JSON(http.StatusBadRequest, Response{Res: "byte conversion error"})
		return
	}
	if err := os.Remove("./images/" + reqURI.ImageId); err != nil { //update기 때문에 만약 파일이 존재하지 않는다면 업데이트를 할 수 없으므로 미리 삭제 명령을 사용하여 파일의 존재 여부를 파악하고 파일이 존재하면 삭제
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "there is exist file. file name : " + reqURI.ImageId})
		return
	}
	if err := ioutil.WriteFile("./images/"+reqURI.ImageId, data, 0755); err != nil { //파일 생성
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "Do not create file"})
		return
	}
	c.JSON(http.StatusOK, Response{Res: "success to update images"})

}
