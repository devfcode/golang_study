package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang_study/myprojects/testurl/mystruct"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.OPTIONS("/getTest", Test)
	r.GET("/getTest", GetTest)
	r.OPTIONS("/getPost", Test)
	r.POST("/getPost", PostTest)
	r.OPTIONS("/testFormdata", Test)
	r.POST("/testFormdata", TestFormdata)
	r.OPTIONS("/upload", Test)
	r.POST("/upload", uploadFile)

	r.GET("/user_info", GetUser)
	r.POST("/login", PostUser)
	r.Run("103.100.211.187:8848") //103.100.211.187	127.0.0.1    www.anant.club
	//使用TLS协议
	//r.RunTLS("103.100.211.187:8848", "/etc/letsencrypt/live/www.anant.club/cert.pem", "/etc/letsencrypt/live/www.anant.club/privkey.pem")
}

//Retrofit Post
func PostUser(c *gin.Context) {
	var user mystruct.User
	user.Name = c.PostForm("name")
	user.Num = rand.Int() % 1000
	user.Password = c.PostForm("password")
	c.JSON(200, user)
}

//Retrofit Get
func GetUser(c *gin.Context) {
	var user mystruct.User
	user.Name = c.DefaultQuery("name", "no value")
	user.Password = c.PostForm("password")
	user.Num = rand.Int() % 1000
	c.JSON(200, user)
}

func Test(c *gin.Context) {
	c.JSON(200, "opption test is success!")
}

//POST 提交form表单
func TestFormdata(c *gin.Context) {
	username := c.PostForm("username")
	area := c.PostForm("area")
	age := c.PostForm("age")
	action := c.PostForm("action")

	if area == "" || username == "" || age == "" || action == "" {
		c.JSON(402, mystruct.Tip{401, "格式错误"})
		return
	}

	var content mystruct.MyContent
	content.Id = rand.Int()
	content.Username = username
	content.Area = area
	content.MyWords = "you form-data is success"
	content.Createtime = time.Now().Unix()
	c.JSON(200, content)
}

//GET 请求
func GetTest(c *gin.Context) {
	var content mystruct.MyContent
	content.Id = rand.Int()
	content.Username = "Dio Brand"
	content.Area = "Guiyang"
	content.MyWords = "You do a get request"
	content.Createtime = time.Now().Unix()
	c.JSON(200, content)
}

//POST 以JSON格式传递参数
func PostTest(c *gin.Context) {
	var userInfo mystruct.UserInfo
	err := c.BindJSON(&userInfo)
	if err != nil {
		fmt.Println(err)
		c.JSON(403, mystruct.Tip{403, "参数格式不对"})
		return
	}
	fmt.Println(userInfo)

	if userInfo.Username == "Dio" && userInfo.Argot == "You are geat!" {
		var content mystruct.MyContent
		content.Id = rand.Int()
		content.Username = "Dio Brand"
		content.Area = "Guiyang"
		content.Createtime = time.Now().Unix()
		content.MyWords = "Happy ! you do it"
		c.JSON(200, content)
		return
	}

	if userInfo.Argot != "You are geat!" {
		c.JSON(401, mystruct.Tip{401, "你的暗号不对"})
		return
	}

	c.JSON(401, mystruct.Tip{401, "你请求成功了，但参数不是我想要的"})
}

//上传文件
func uploadFile(context *gin.Context) {
	var result mystruct.Result
	file, header, err := context.Request.FormFile("file") //获取传过来的文件
	if err == nil {
		filename := header.Filename                                   //从header中取得文件名
		out, err := os.Create("../images/uploadedImages/" + filename) //服务器上存储图片的文件夹路径+图片名
		if err == nil {
			defer out.Close()
			_, err = io.Copy(out, file) //将传过来的二进制图片file 复制到刚才创建的out文件里
			if err == nil {
				log.Println("上传图片成功")
				res := map[string]interface{}{
					"filePath": "/yourimagepath/" + filename, //图片网络请求路径
					"fileName": filename,
				}
				result.Code = 0
				result.Data = res
				result.Msg = "图片上传成功"
			} else {
				result.Code = -3
				result.Msg = "复制文件出错"
			}
		} else {
			result.Code = -2
			result.Msg = "创建文件出错"
		}
	} else {
		result.Code = -1
		result.Msg = "接收文件出错"
	}
	context.JSON(http.StatusOK, result)
}
