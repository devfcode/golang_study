package controllers

import (
	"github.com/gin-gonic/gin"
	"golang_study/myprojects/novel/db"
	"golang_study/myprojects/novel/structs"
	"log"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

var err error


//书籍管理 	修改数据
func BookUpdate(c *gin.Context)  {
	var temp structs.BookInfo
	err = c.BindJSON(&temp)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	fmt.Println(temp)

	var condition map[string]interface{}  //查询条件
	condition = make(map[string]interface{})
	if temp.Category != ""{
		condition["category"] = temp.Category
	}
	if temp.SubCategory != "" {
		condition["sub_category"] = temp.SubCategory
	}
	if temp.ChapterPrice != 0 {
		condition["chapter_price"] = temp.ChapterPrice
	}
	//if temp.FreeChapterCount != 0 {
	//	condition["free_chapter_count"] = temp.FreeChapterCount
	//}
	condition["free"] = temp.Free
	condition["hide"] = temp.Hide

	session, c1 := db.GetConnection("qidian", "book")
	defer session.Close()
	err = c1.Update(bson.M{"_id":bson.ObjectIdHex(temp.Id)},bson.M{"$set":condition})
	if err != nil {
		c.JSON(501, "更新出错!")
		log.Fatal(err)
	}else {
		c.JSON(200, structs.Tip{Code:200, Msg:"更新成功!"})
	}

	var i int
	for i = 0; i < temp.FreeChapterCount ; i++  { //把前 n 章设成免费的
		var param string  =  "chapter_list."+strconv.Itoa(i)+".free"
		err := c1.Update(bson.M{"_id":bson.ObjectIdHex(temp.Id)}, bson.M{"$set":bson.M{param:true}})
		if err != nil {
			log.Fatal(err)
		}
	}
}

//书籍管理  显示表的数据
var bookPageIndex int64 = 1
func BookManageTable(c *gin.Context)  {
	name := c.Query("name")
	sub_category := c.Query("subCategory")
	x := c.Query("pageIndex")
	if x != "" {//判断传过来的pageIndex不为空
		pageIndex2, err3 := strconv.ParseInt(x, 10, 64)
		if err3 != nil {
			log.Fatal(err3)
		}else {
			bookPageIndex = pageIndex2
		}
	}
	//fmt.Println(name,"\t",sub_category)
	session, c1 := db.GetConnection("qidian", "book")
	defer session.Close()
	var condition map[string]interface{}  //查询条件
	condition = make(map[string]interface{})
	if name != "" {
		var param string = name+"*" //关键字的正则匹配
		condition["name"] = bson.M{"$regex": param}  //书名的正则匹配
	}
	if sub_category != "" {
		condition["sub_category"] = sub_category
	}
	//fmt.Println(condition)
	//查询总数据量   做分页
	start1 := time.Now()
	countNum, err1 := c1.Find(condition).Count()
	end1 := time.Now()
	fmt.Println("1111\t",end1.Sub(start1))
	if err1 != nil {
		c.JSON(301, err1)
		log.Fatal(err1)
	}
	//fmt.Println("aaaaaa\t",countNum)
	var content structs.Content
	var data []structs.BookList
	start2 := time.Now()
	err = c1.Find(condition).Select(bson.M{"_id":1, "category":1, "sub_category":1,  "name":1, "author":1, "free":1, "create_at":1, "hide":1, "chapter_list":1}).Skip(int(bookPageIndex - 1) * 15).Limit(15).Sort("update_at").All(&data)
	if err != nil {
		c.JSON(300, err)
		log.Fatal(err)
	}else {
		var i int
		for i  = 0; i < len(data); i++ { //计算前多少章免费、总字数
			var num1 int = 0
			var num2 int64 = 0
			var num3 int = 0
			for _, v2 := range data[i].ChapterList {
				if v2.Crawled == false { //如果这一章未爬取 则不做统计
					break
				}
				if v2.Free {
					num1 ++
				}
				num3 ++
				//fmt.Println("wordcount:\t",v2.Name)
				//fmt.Println("wordcount:\t",v2.Words)
				tempNum, err3 := strconv.ParseInt(v2.Words, 10, 64)
				if err3 != nil {
					log.Fatal(err3)
				}
				num2 += tempNum
				//fmt.Println(num2)
			}
			data[i].FreeChapterCount = num1 //前多少章免费
			data[i].WordsCount = num2 //总字数
			data[i].ChapterCount = num3 //小说已经爬下来的章节数， 有些付费小说可能不会全爬完
			content.Data = append(content.Data, data[i])
		}
			end2 := time.Now()
			fmt.Println("2222\t",end2.Sub(start2))
			//头
			var header structs.Header
			header.Title = "test"
			header.Token = "TODKFAJKLFDKSFAJKFASDFKSAFEWLO"
			//其他
			var other structs.Other
			other.Total = countNum
			other.PageIndex = bookPageIndex
			other.PageSize = 15
			//数据体
			//var content structs.Content
			content.Header = header
			content.Other = other
			//content.Data = data
			//最终返回给前端的数据
			var result structs.ResultList
			result.Code = 200
			result.Msg = "数据请求成功!"
			result.Content = content

			c.JSON(200, result)
	}

	//if err != nil {
	//	c.JSON(300, err)
	//	log.Fatal(err)
	//}else {
	//	//for _, v := range content.Data{
	//	//	fmt.Println(v)
	//	//}
	//	//头
	//	var header structs.Header
	//	header.Title = "test"
	//	header.Token = "TODKFAJKLFDKSFAJKFASDFKSAFEWLO"
	//	//其他
	//	var other structs.Other
	//	other.Total = countNum
	//	other.PageIndex = bookPageIndex
	//	other.PageSize = 15
	//	//数据体
	//	//var content structs.Content
	//	content.Header = header
	//	content.Other = other
	//	//content.Data = backAccount
	//	//最终返回给前端的数据
	//	var result structs.ResultList
	//	result.Code = 200
	//	result.Msg = "数据请求成功!"
	//	result.Content = content
	//
	//	c.JSON(200, result)
	//}
}

//用户反馈	客服管理 --- 玩家反馈  回复界面
/*func GetHistoryRecord(c *gin01.Context)  {
	user_id := c.Query("user_id")
	session, c1 := utils.GetConnection("qidian", "chatrecord")
	defer session.Close()
	var content structs.Content
	err = c1.Find(bson.M{"user_id":user_id}).Sort("reply_time").All(&content.Data)
	if err != nil {
		c.JSON(300, err)
		log.Fatal(err)
	}else {
		//for _, v := range temps{
		//	fmt.Println(v.User)
		//}
		//头
		var header structs.Header
		header.Title = "test"
		header.Token = "TODKFAJKLFDKSFAJKFASDFKSAFEWLO"
		//其他
		var other structs.Other
		other.Total = 1
		other.PageIndex = 1
		other.PageSize = 15
		//数据体
		//var content structs.Content
		content.Header = header
		content.Other = other
		//content.Data = backAccount
		//最终返回给前端的数据
		var result structs.ResultList
		result.Code = 200
		result.Msg = "数据请求成功!"
		result.Content = content

		c.JSON(200, result)
	}
}*/
func Reply(c *gin.Context)  {
	var temp structs.ChatRecord
	err = c.BindJSON(&temp)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	session, c1 := db.GetConnection("qidian", "userfeedback")
	defer session.Close()
	fmt.Println("temp:",temp)
	err = c1.Update(bson.M{"num":temp.Num,"user_id":temp.UserId},bson.M{"$set":bson.M{"reply_time":temp.ReplyTime, "text":temp.Text, "status":"已处理"}})
	if err != nil {
		c.JSON(501, "更新出错!")
		log.Fatal(err)
	}else {
		c.JSON(200, structs.Tip{Code:200, Msg:"更新成功!"})
	}
}

//用户反馈	客服管理 --- 玩家反馈  点击标记
func ReMark(c *gin.Context)  {
	id := c.Query("id")
	mark := c.Query("mark")
	session, c1 := db.GetConnection("qidian", "userfeedback")
	defer session.Close()
	//fmt.Println(mark)
	err = c1.Update(bson.M{"_id":bson.ObjectIdHex(id)},bson.M{"$set":bson.M{"mark":mark}})
	if err != nil {
		c.JSON(501, "更新出错!")
		log.Fatal(err)
	}else {
		c.JSON(200, structs.Tip{Code:200, Msg:"更新成功!"})
	}
}

//用户反馈	客服管理 --- 玩家反馈  忽略反馈
func IgnoreFeedback(c *gin.Context)  {
	id := c.Query("id")
	session, c1 := db.GetConnection("qidian", "userfeedback")
	defer session.Close()
	//res, err1 := hex.DecodeString(id)
	//if err1 != nil {
	//	log.Fatal(err1)
	//}
	err = c1.Update(bson.M{"_id":bson.ObjectIdHex(id)},bson.M{"$set":bson.M{"status":"忽略"}})
	if err != nil {
		c.JSON(501, "插入出错!")
		log.Fatal(err)
	}else {
		c.JSON(200, structs.Tip{Code:200, Msg:"更新成功!"})
	}
}

//用户反馈	客服管理 --- 玩家反馈 删除反馈
func DeleteFeedback(c *gin.Context)  {
	id := c.Query("id")
	session, c1 := db.GetConnection("qidian", "userfeedback")
	defer session.Close()
	//res, err1 := hex.DecodeString(id)
	//if err1 != nil {
	//	log.Fatal(err1)
	//}
	_, err := c1.RemoveAll(bson.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		c.JSON(501,err)
		log.Fatal(err)
	} else {
		c.JSON(200, structs.Tip{Code:200, Msg:"删除成功!"})
	}
}

var servicePageIndex int64 = 1
//用户反馈	客服管理 --- 玩家反馈  获取表格数据
func UserFeedbackList(c *gin.Context)  {
	status := c.Query("status")
	account := c.Query("account")
	user_id := c.Query("user_id")
	problem := c.Query("problem")
	startTime := c.Query("start_time")
	endTime := c.Query("end_time")
	x := c.Query("pageIndex")
	if x != "" {//判断传过来的pageIndex不为空
		pageIndex2, err3 := strconv.ParseInt(x, 10, 64)
		if err3 != nil {
			log.Fatal(err3)
		}else {
			servicePageIndex = pageIndex2
		}
	}
	//fmt.Print(status, account, user_id, problem, startTime, endTime, servicePageIndex)
	var condition map[string]interface{}  //查询条件
	condition = make(map[string]interface{})
	if status != "" {
		condition["status"] = status
	}
	if account != "" {
		condition["account"] = account
 	}
	if user_id != "" {
		condition["user_id"] = user_id
	}
	if problem != "" {
		var param string = problem+"*" //关键字的正则匹配
		condition["problem"] = bson.M{"$regex": param}
	}
	if startTime != "" {
		fmt.Println(startTime,"\t",endTime)
		condition["create_time"] = bson.M{"$gte":startTime, "$lte":endTime}
	}

	session, c1 := db.GetConnection("qidian", "userfeedback")
	defer session.Close()
	//查询总数据量
	countNum, err1 := c1.Find(condition).Count()
	if err1 != nil {
		c.JSON(301, err1)
		log.Fatal(err)
	}
	//数据体
	var content structs.Content
	err = c1.Find(condition).Skip(int(servicePageIndex - 1) * 15).Limit(15).Sort("-create_time").All(&content.Data)
	if err != nil {
		c.JSON(300, err)
		log.Fatal(err)
	}else {
		//for _, v := range temps{
		//	fmt.Println(v.User)
		//}
		//头
		var header structs.Header
		header.Title = "test"
		header.Token = "TODKFAJKLFDKSFAJKFASDFKSAFEWLO"
		//其他
		var other structs.Other
		other.Total = countNum
		other.PageIndex = servicePageIndex
		other.PageSize = 15
		//数据体
		//var content structs.Content
		content.Header = header
		content.Other = other
		//content.Data = backAccount
		//最终返回给前端的数据
		var result structs.ResultList
		result.Code = 200
		result.Msg = "数据请求成功!"
		result.Content = content

		c.JSON(200, result)
	}
}


//公告管理  删除公告
func DeleteNotice(c *gin.Context)  {
	id := c.Query("id")
	session, c1 := db.GetConnection("qidian", "noticemanage")
	defer session.Close()
	//res, err1 := hex.DecodeString(id) //取到页面的字符串id后要经过 16进制解码 然后再转成ObjectId ,系统的ObjectId()转换只能处理12字节的，而这的字符串长度是24字节的，所以要自己先处理一次
	//if err1 != nil {
	//	log.Fatal(err1)
	//}
	//fmt.Println("res1:\t",res)
	//fmt.Println("res2:", bson.ObjectId(res))
	//fmt.Println("res3:\t",string(res))
	_, err := c1.RemoveAll(bson.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		c.JSON(501,err)
		log.Fatal(err)
	} else {
		c.JSON(200, structs.Tip{Code:200, Msg:"删除成功!"})
	}
}
//公告管理  关闭 和 编辑
func Edit(c *gin.Context)  {
	var temp structs.NoticeEdit
	err = c.BindJSON(&temp)
	if err != nil {
		log.Fatal(err)
	}
	//else {
	//	fmt.Println("temp:\n",temp)
	//}

	var condition map[string]interface{}  //查询条件
	condition = make(map[string]interface{})
	if temp.Type != "" {
		condition["type"] = temp.Type
	}
	if temp.Priority != ""{
		condition["priority"] = temp.Priority
	}
	if temp.Status != "" {
		condition["status"] = temp.Status
	}
	if temp.Platform != "" {
		condition["platform"] = temp.Platform
	}
	if temp.StartTime != "" {
		condition["start_time"] = temp.StartTime
	}
	if temp.EndTime != "" {
		condition["end_time"] = temp.EndTime
	}
	if temp.Substance != "" {
		condition["substance"] = temp.Substance
	}

	session, c1 := db.GetConnection("qidian", "noticemanage")
	defer session.Close()
	//res, err1 := hex.DecodeString(temp.Id)
	//if err1 != nil {
	//	fmt.Println(err1)
	//	//log.Fatal(err1)
	//}
	err = c1.Update(bson.M{"_id":bson.ObjectIdHex(temp.Id)},bson.M{"$set":condition})
	if err != nil {
		c.JSON(501, "插入出错!")
		log.Fatal(err)
	}else {
		c.JSON(200, structs.Tip{Code:200, Msg:"更新成功!"})
	}
}

var noticePageIndex int64 = 0
//公告管理 公告列表
func NoticeList(c *gin.Context)  {
	typeN := c.Query("type")
	priority := c.Query("priority")
	status := c.Query("status")
	platform := c.Query("platform")
	x := c.Query("pageIndex")
	if x != "" {//判断传过来的pageIndex不为空
		pageIndex2, err3 := strconv.ParseInt(x, 10, 64)
		if err3 != nil {
			log.Fatal(err3)
		}else {
			noticePageIndex = pageIndex2
		}
	}
	fmt.Print(typeN, priority, noticePageIndex, status, platform)
	var condition map[string]interface{}  //查询条件
	condition = make(map[string]interface{})
	if typeN != "" {
		condition["type"] = typeN
	}
	if priority != ""{
		condition["priority"] = priority
	}
	if status != "" {
		condition["status"] = status
	}
	if platform != "" {
		condition["platform"] = platform
	}
	session, c1 := db.GetConnection("qidian", "noticemanage")
	defer session.Close()
	countNum, err1 := c1.Find(condition).Count()
	if err1 != nil {
		c.JSON(301, err1)
		panic(err1)
	}
	//数据体
	var content structs.Content
	err = c1.Find(condition).Skip(int(noticePageIndex) * 15).Limit(15).Sort("-start_time").All(&content.Data)
	if err != nil {
		c.JSON(300, err)
		log.Fatal(err)
	}else {
		//for _, v := range temps{
		//	fmt.Println(v.User)
		//}
		//头
		var header structs.Header
		header.Title = "test"
		header.Token = "TODKFAJKLFDKSFAJKFASDFKSAFEWLO"
		//其他
		var other structs.Other
		other.Total = countNum
		other.PageIndex = noticePageIndex
		other.PageSize = 15
		//数据体
		//var content structs.Content
		content.Header = header
		content.Other = other
		//content.Data = backAccount
		//最终返回给前端的数据
		var result structs.ResultList
		result.Code = 200
		result.Msg = "数据请求成功!"
		result.Content = content

		c.JSON(200, result)
	}
}

//公告管理  新增公告
func AddNotice(c *gin.Context) {
	var temp structs.Notice
	err = c.BindJSON(&temp)
	if err != nil {
		log.Fatal(err)
		panic(err)
	} else {
		fmt.Println("temp:\n",temp)
	}

	session, c1 := db.GetConnection("qidian", "noticemanage")
	defer session.Close()
	err = c1.Insert(temp)
	if err != nil {
		c.AbortWithStatus(404)
	}else {
		c.JSON(200, structs.Tip{Code:200, Msg:"添加成功!"})
	}
}


var vistPageIndex int64 = 0
//后台访问记录 查看记录
func ShowVistRecord(c *gin.Context)  {
	time := c.Query("time")
	user := c.Query("user")
	x := c.Query("pageIndex")
	if x != "" {//判断传过来的pageIndex不为空
		pageIndex2, err3 := strconv.ParseInt(x, 10, 64)
		if err3 != nil {
			log.Fatal(err3)
		}else {
			vistPageIndex = pageIndex2
		}
	}
	//fmt.Print(time, user, vistPageIndex)
	var condition map[string]interface{}  //查询条件
	condition = make(map[string]interface{})
	if time != "" {
		var param string = "^"+time+"*" //时间的正则匹配
		condition["time"] = bson.M{"$regex": param}
	}
	if user != "" {
		condition["user"] = user
	}
	session, c1 := db.GetConnection("qidian", "backvisit")
	defer session.Close()
	countNum, err1 := c1.Find(condition).Count()
	if err1 != nil {
		c.JSON(301, err1)
		panic(err1)
	}
	//数据体
	var content structs.Content
	err = c1.Find(condition).Skip(int(vistPageIndex)).Limit(15).Sort("-create_time").All(&content.Data)
	if err != nil {
		c.JSON(300, err)
		log.Fatal(err)
	}else {
		//for _, v := range temps{
		//	fmt.Println(v.User)
		//}
		//头
		var header structs.Header
		header.Title = "test"
		header.Token = "TODKFAJKLFDKSFAJKFASDFKSAFEWLO"
		//其他
		var other structs.Other
		other.Total = countNum
		other.PageIndex = pageIndex
		other.PageSize = 15
		//数据体
		//var content structs.Content
		content.Header = header
		content.Other = other
		//content.Data = backAccount
		//最终返回给前端的数据
		var result structs.ResultList
		result.Code = 200
		result.Msg = "数据请求成功!"
		result.Content = content

		c.JSON(200, result)
	}

}

//后台访问记录 添加记录
func BackVistAdd(c *gin.Context)  {
	var temp structs.VistRecord
	err = c.BindJSON(&temp)
	if err != nil {
		log.Fatal(err)
		panic(err)
	} else {
		fmt.Println("temp:\n",temp)
	}

	session, c1 := db.GetConnection("qidian", "backvisit")
	defer session.Close()
	err = c1.Insert(temp)
	if err != nil {
		c.AbortWithStatus(404)
	}else {
		c.JSON(200, structs.Tip{Code:200, Msg:"添加成功!"})
	}
}


//后台权限管理 	删除账号
func DeleteAccount(c *gin.Context)  {
	account := c.Query("account")
	session, c1 := db.GetConnection("qidian", "backaccount")
	defer session.Close()
	_, err := c1.RemoveAll(bson.M{"account": account})
	if err != nil {
		c.AbortWithStatus(404)
		log.Fatal(err)
	} else {
		c.JSON(200, structs.Tip{Code:200, Msg:"删除成功!"})
	}
}


var pageIndex int64 = 0
//后台权限管理  	查询数据
func QueryBackground(c *gin.Context)  {
	account := c.Query("account")
	remarks := c.Query("remarks")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	x := c.Query("pageIndex")
	if x != "" {//判断传过来的pageIndex不为空
		pageIndex2, err3 := strconv.ParseInt(x, 10, 64)
		if err3 != nil {
			log.Fatal(err3)
		}else {
			pageIndex = pageIndex2
		}
	}
	fmt.Printf("account = %s\t remarks = %s\t startTime = %s\t endTime = %s \n", account, remarks, startTime, endTime)
	var conditionN map[string]interface{}   //查询条件
	conditionN = make(map[string]interface{})
	if account != "" {
		conditionN["account"] = account
	}
	if remarks != "" {
		conditionN["remarks"] = remarks
	}
	if startTime != "" {
		conditionN["create_time"] = bson.M{"$gte":startTime}
	}
	if endTime != "" {
		conditionN["create_time"] = bson.M{"$lte":endTime}
	}

	session, c1 := db.GetConnection("qidian", "backaccount")
	defer session.Close()
	//condition := bson.M{"account" : account, "remarks" : remarks}//, "create_time":bson.M{"$in":[]string{startTime, endTime}}}       //   , "remarks" : remarks, "create_time":bson.M{"$in":[]string{startTime, endTime}}
	countNum, err1 := c1.Find(conditionN).Count()  //bson.M{"account" : account}
	if err1 != nil {
		c.JSON(301, err1)
		panic(err1)
	}
	//var backAccount []structs.BackAccount
	//数据体
	var content structs.Content
	err = c1.Find(conditionN).Skip(int(pageIndex)).Limit(15).Sort("-create_time").All(&content.Data) //bson.M{"account" : account}
	if err != nil {
		c.JSON(300, err)
		log.Fatal(err)
	}else {
		//for _, v := range backAccount{
		//	fmt.Println(v.Account)
		//}
		//头
		var header structs.Header
		header.Title = "test"
		header.Token = "TODKFAJKLFDKSFAJKFASDFKSAFEWLO"
		//其他
		var other structs.Other
		other.Total = countNum
		other.PageIndex = pageIndex
		other.PageSize = 15
		//数据体
		content.Header = header
		content.Other = other
		//content.Data = backAccount
		//最终返回给前端的数据
		var result structs.ResultList
		result.Code = 200
		result.Msg = "数据请求成功!"
		result.Content = content

		c.JSON(200, result)
	}
}

//后台权限管理   插入数据
func AddAccount(c *gin.Context)  {
	var temp structs.BackAccount
	err = c.BindJSON(&temp)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	//else {
	//	fmt.Println("name:",temp.Account)
	//}

	session, c1 := db.GetConnection("qidian", "backaccount")
	defer session.Close()
	err = c1.Insert(temp)
	if err != nil {
		c.AbortWithStatus(404)
	}else {
		c.JSON(200, structs.Tip{Code:200, Msg:"添加成功!"})
	}
}
