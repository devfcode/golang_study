package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang_study/myprojects/diary/structs"
	"golang_study/myprojects/diary/utils"
	"net/http"
)

var mydb *gorm.DB

var err error

//注册账户
func Register(c *gin.Context)  {
	//buf := make([]byte, 1024)
	//n, _ := c.Request.Body.Read(buf)
	//fmt.Println(string(buf[0:n]))
	var para structs.Account
	err = c.BindJSON(&para)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("username:",para.Username)
		fmt.Println("passwords:",para.Passwords)
		fmt.Println("createtime:",para.Createtime)
	}
	mydb = utils.Init()
	var count int
	mydb.Table("account").Where("username = ?", para.Username).Count(&count)
	if count != 0 {
		c.JSON(http.StatusOK, structs.Tip{304, "该账号已注册过了"})
		fmt.Println("该账号已注册过了")
		return
	}

	err = mydb.Table("account").Create(&para).Error
	defer mydb.Close()
	if err != nil {
		c.JSON(http.StatusSeeOther, structs.Tip{http.StatusSeeOther, "注册失败"})
		panic(err)
	}
	c.JSON(http.StatusOK,structs.Tip{http.StatusOK, "success"})
	fmt.Println("insert successed!")
}

//登陆
func Login(c *gin.Context)  {
	var para structs.Account
	err = c.BindJSON(&para)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("username:",para.Username)
		fmt.Println("passwords:",para.Passwords)
	}
	mydb = utils.Init()
	var count int
	mydb.Table("account").Where("username = ?", para.Username).Count(&count)
	if count == 0 {
		c.JSON(http.StatusOK, structs.Tip{304, "该账号尚未注册过"})
		fmt.Println("该账号尚未注册过")
		return
	}

	var result structs.Account
	err = mydb.Table("account").Select("passwords").Where("username = ?", para.Username).First(&result).Error
	if err != nil {
		panic(err)
	}
	if result.Passwords == para.Passwords {
		c.JSON(http.StatusOK, structs.Tip{200, "账号验证正确"})
	}else {
		c.JSON(http.StatusOK, structs.Tip{401, "账号或密码错误"})
	}
	defer mydb.Close()
}

//上传笔记
func Upload(c *gin.Context)  {
	var para structs.Diary
	err = c.BindJSON(&para)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("username:",para.Username)
	}
	mydb = utils.Init()
	err = mydb.Table("diary").Create(&para).Error
	if err != nil {
		c.JSON(http.StatusSeeOther, structs.Tip{http.StatusSeeOther, "上传失败"})
		panic(err)
		return
	}
	c.JSON(http.StatusOK,structs.Tip{http.StatusOK, "上传成功!"})
	fmt.Println("upload successed!")
	defer mydb.Close()
}

//更改日志
func UpdateDiary(c *gin.Context)  {
	var para structs.Diary
	err = c.BindJSON(&para)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("username:",para.Username)
	}
	mydb = utils.Init()
	err = mydb.Table("diary").Where("username = ? and diaryid = ?", para.Username,para.Diaryid).
		Updates(structs.Diary{Title:para.Title,Content:para.Content,Lastupdatetime:para.Lastupdatetime}).Error
	if err != nil {
		c.JSON(http.StatusOK, structs.Tip{http.StatusSeeOther, "更新失败!"})
		panic(err)
		return
	}
	c.JSON(http.StatusOK,structs.Tip{http.StatusOK, "更新成功!"})
	fmt.Println("update successed!")
	defer mydb.Close()
}

//删除日志
func DeleteDiary(c *gin.Context)  {
	var para structs.Diary
	err = c.BindJSON(&para)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("username:",para.Username)
	}
	mydb = utils.Init()
	err = mydb.Table("diary").Where("username = ? and diaryid = ?", para.Username,para.Diaryid).
		Delete(structs.Diary{}).Error
	if err != nil {
		c.JSON(http.StatusOK, structs.Tip{http.StatusSeeOther, "删除失败!"})
		panic(err)
		return
	}
	c.JSON(http.StatusOK,structs.Tip{http.StatusOK, "删除成功!"})
	fmt.Println("delete successed!")
	defer mydb.Close()
}
