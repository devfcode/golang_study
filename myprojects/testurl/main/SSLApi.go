package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang_study/myprojects/testurl/mystruct"
	"time"
	"math/rand"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/getssl", GetSSLTest)
	r.POST("/postssl", PostSSLTest)
	r.RunTLS("103.100.211.187:8081", "/etc/letsencrypt/live/www.anant.club/cert.pem", "/etc/letsencrypt/live/www.anant.club/privkey.pem")
}

//GET 请求
func GetSSLTest(c *gin.Context) {
	var content mystruct.MyContent
	content.Id = rand.Int()
	content.Username = "Dio Brand"
	content.Area = "Guiyang"
	content.MyWords = "You do a get request"
	content.Createtime = time.Now().Unix()
	c.JSON(200, content)
}

//POST 以JSON格式传递参数
func PostSSLTest(c *gin.Context) {
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
