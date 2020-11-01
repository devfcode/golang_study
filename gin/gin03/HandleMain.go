package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	engine := gin.Default()

	//http://localhost:8080/hello?name=Tom
	engine.Handle("GET", "/hello", func(context *gin.Context) {
		//取到全路径
		path := context.FullPath()
		fmt.Println(path)

		//取到路径上name参数的值
		name := context.DefaultQuery("name", "no value")
		fmt.Println(name)

		//输出
		context.Writer.Write([]byte("Hello, I am Server!  name is " + name))
	})

	//http://localhost:8080/login
	engine.Handle("POST", "login", func(context *gin.Context) {
		//取到全路径
		path := context.FullPath()
		fmt.Println(path)

		//Form 里的参数值
		username := context.PostForm("username")
		password := context.PostForm("password")
		fmt.Println(username)
		fmt.Println(password)

		//输出
		context.Writer.Write([]byte("Hello, I am Server!  username is " + username))
	})

	engine.Run("127.0.0.1:8080")
}
