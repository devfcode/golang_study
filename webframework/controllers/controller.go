package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang_study/webframework/db"
	"golang_study/webframework/structs"
)

var mydb *gorm.DB

func GetProjects(c *gin.Context) {
	para := c.Query("name")
	fmt.Printf("name = %s\n",para)
	mydb = db.Init()
	var stu []structs.Student
	if err := mydb.Table("student").Find(&stu).Error; err != nil {
		c.AbortWithStatus(404)
		panic(err)
	} else {
		fmt.Println(stu)
		c.JSON(200, stu)
	}
}
