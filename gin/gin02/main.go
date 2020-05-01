package main

import (
	//"gopkg.in/gin-gonic/gin.v1"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"os"
	"log"
	"io"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.LoadHTMLGlob("gin/gin02/templates/*")
	router.GET("/", all)//请求重定向,把空请求定向到 /home 地址
	router.GET("/home", home)//把上传文件上页面传给浏览器
	router.POST("/upload", upload)//上传文件接口,图片以二进制的形式传过来

	router.POST("/form", getForm)//post body form-data 中的参数
	router.POST("/gepostbody", getPost)//post body-raw 中的 json参数

	router.Run("127.0.0.1:10004")
}

func all(c *gin.Context)  {
	fmt.Println("这里是重定向!")
	//gin对于重定向的请求
	//c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:10004/home")
}

func home(c *gin.Context)  {
	fmt.Println("this is home")
	c.HTML(http.StatusOK, "upload.html", gin.H{})
}

func upload(c *gin.Context)  {
	name := c.PostForm("mytext")
	fmt.Println("name:",name)
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	filename := header.Filename

	fmt.Println(file, err, filename)

	var filepath string = "gin/gin02/files/" + filename
	out, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	c.HTML(http.StatusOK, "success.html", gin.H{})
}

//获取  form-data中的参数
func getForm(c *gin.Context)  {
	name := c.PostForm("name")
	var email string = c.PostForm("email")
	fmt.Println(name,"   ++++    ",email)
	c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:10004/home")
}


type Para struct {
	Name string `json:"name"`
	Email  string    `json:"email"`
}
func getPost(c *gin.Context)  {//获取post body中的参数
	var reqInfo Para
	var err = c.BindJSON(&reqInfo)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("name:",reqInfo.Name)
		fmt.Println("age:",reqInfo.Email)
	}
}