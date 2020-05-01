package main

import (
	"github.com/gin-gonic/gin"
	"golang_study/myprojects/diary/controllers"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/register", controllers.Register)
	r.POST("/login",controllers.Login)
	r.POST("/upload",controllers.Upload)
	r.POST("/update",controllers.UpdateDiary)
	r.POST("/delete",controllers.DeleteDiary)
	r.Run(":10000")
}
