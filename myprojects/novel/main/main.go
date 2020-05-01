package main

import (
	"github.com/gin-gonic/gin"
	"golang_study/myprojects/novel/controllers"
)


func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	//后台权限管理
	r.OPTIONS("/addaccount", Test)//添加账号
	r.POST("/addaccount", controllers.AddAccount)//添加账号
	r.OPTIONS("/showtable", Test)//显示数据
	r.GET("/showtable", controllers.QueryBackground)//显示数据
	r.OPTIONS("/deleteaccount", Test)//删除账号
	r.GET("/deleteaccount", controllers.DeleteAccount)//删除账号
	//后台访问记录
	r.POST("/backvistadd", controllers.BackVistAdd)
	r.GET("/showvistrecord", controllers.ShowVistRecord)
	//公告管理
	r.POST("/noticeadd", controllers.AddNotice)
	r.GET("/noticelist", controllers.NoticeList)
	r.POST("/edit", controllers.Edit)
	r.GET("/deletenotice", controllers.DeleteNotice)
	//用户反馈
	r.GET("/userfeedbacklist",controllers.UserFeedbackList)
	r.GET("/deletefeedback", controllers.DeleteFeedback)
	r.GET("/ignorefeedback", controllers.IgnoreFeedback)
	r.GET("/remark", controllers.ReMark)
	r.POST("/reply", controllers.Reply)
	//r.GET("/gethistoryrecord", controllers.GetHistoryRecord)
	//书籍管理
	r.GET("/booklist", controllers.BookManageTable)
	r.POST("/bookupdate", controllers.BookUpdate)


	//测试
	r.GET("/showOne", controllers.QueryOne)
	r.GET("/insert", controllers.Insert)
	r.POST("/showMany", controllers.QueryMany)
	r.GET("/update", controllers.Update)
	r.GET("/delete", controllers.Delete)
	r.GET("/count", controllers.Count)
	//预请求测试
	r.OPTIONS("/test", Test)

	r.Run(":8088")
}


func Test(c *gin.Context)  {
	c.JSON(200, "opption test is success!")
}
