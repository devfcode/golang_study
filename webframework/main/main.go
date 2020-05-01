package main

import (
	"github.com/gin-gonic/gin"
	"golang_study/webframework/controllers"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/webframework", controllers.GetProjects)
	//r.GET("/webframework/:id", controllers.GetProjects)

	r.Run(":10000")
}
