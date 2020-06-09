package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	//预请求测试
	r.OPTIONS("/count", Test)
	r.GET("/count", Mytest)

	r.Run("103.100.211.187:10004")//103.100.211.187	127.0.0.1
}

func Test(c *gin.Context)  {
	c.JSON(200, "opption test is success!")
}

func Mytest(c *gin.Context)  {
	c.JSON(201, "this is the value!")
}