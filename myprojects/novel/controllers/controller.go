package controllers

import (
	"fmt"
	"golang_study/myprojects/novel/db"
	"golang_study/myprojects/novel/structs"
	"gopkg.in/mgo.v2/bson"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
	"log"
)


//插入数据
func Insert(c *gin.Context) {
	rand.Seed(time.Now().Unix())
	rnd1 := rand.Intn(990)
	rnd2 := rand.Intn(990)

	temp := structs.Service{
		Num:        rnd1,
		UserId:     109095,
		Account:    "游客",
		Mark:       "穷人",
		Problem:    "Nice to meet you",
		Status:     "未处理",
		IpAddr:     "192.168.10.3(黑龙江)",
		CreateTime: "2017/4/17 12:56",
	}
	temp2 := structs.Service{
		Num:        rnd2,
		UserId:     109095,
		Account:    "游客",
		Mark:       "骗子",
		Problem:    "Nice to meet you",
		Status:     "未处理",
		IpAddr:     "192.168.10.4(湖北)",
		CreateTime: "2017/4/17 23:56",
	}
	session, c1 := db.GetConnection("qidian", "service")
	defer session.Close()
	var err error
	err = c1.Insert(temp, temp2)
	if err != nil {
		c.AbortWithStatus(404)
		panic(err)
	} else {
		c.JSON(200, "insert successed!")
	}
	fmt.Println("insert successed!")
}

//查询单条数据
func QueryOne(c *gin.Context) {
	para := c.Query("name") //获取url上的参数
	fmt.Printf("name = %s\n", para)

	session, c1 := db.GetConnection("qidian", "service")
	defer session.Close()

	var result structs.Service
	err := c1.Find(bson.M{"mark": "标记"}).One(&result)
	if err != nil {
		c.AbortWithStatus(404)
		panic(err)
	} else {
		fmt.Println("success:", result)
		c.JSON(200, result)
	}
	//fmt.Println("num:", result.Num)
	//fmt.Println("user_id:", result.UserId)
	//fmt.Println("account:", result.Account)
	//fmt.Println("mark:", result.Mark)
	//fmt.Println("problem:", result.Problem)
	//fmt.Println("status:", result.Status)
	//fmt.Println("ip_address:", result.IpAddr)
	//fmt.Println("create_time:",result.CreateTime)
	//fmt.Println("This is All!!")
}

//type Para structs {//post中参数属性
//	Name string `schema:"name"`
//	Age  int    `schema:"age"`
//}
//var decoder = schema.NewDecoder()
type Para struct {//post中参数属性
	Name string `json:"name"`
	Age  int    `json:"age"`
}
//查询多条数据
func QueryMany(c *gin.Context) {
	//name=elm&age=18   样式
	//var para Para
	//buf := make([]byte, 1024)
	//n, _ := c.server.Body.Read(buf)  //获取post中body里的参数
	//////fmt.Println("buf:", string(buf[0:n]), "\n", "n:", n)
	//qs, _ := url.ParseQuery(string(buf[0:n]))
	////fmt.Println(string(buf[0:n]))
	//err :=decoder.Decode(&para, qs)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(qs)
	//fmt.Println("query:")
	//fmt.Println(para.Name)
	//fmt.Println(para.Age)

	/*{  样式 json
		"name":"zhanshan",
		"age":18
	}*/
	var reqInfo Para
	err = c.BindJSON(&reqInfo)
	if err != nil {
		log.Fatal(err)
		return
	} else {
		fmt.Println("name:",reqInfo.Name)
	}
	

	session, c1 := db.GetConnection("qidian", "service")
	defer session.Close()

	var service  structs.Service
	var serviceAll structs.ServiceList
	serviceAll.Code = 200;
	serviceAll.Msg = "数据请求成功!"
	//var serArr  []structs.Service
	iter := c1.Find(bson.M{"num" : bson.M{"$lt":580}}).Iter()
	for iter.Next(&service) {
		fmt.Printf("Num: %v\n", service.Num)
		serviceAll.Content = append(serviceAll.Content, service)
		//serArr = append(serArr, service)
	}
	//c.JSON(200, serArr)
	c.JSON(200, serviceAll)
}

//更新数据
func Update(c *gin.Context) {
	session, c1 := db.GetConnection("qidian", "service")
	defer session.Close()

	err := c1.Update(bson.M{"num": 563}, bson.M{"$set": bson.M{"mark": "蠢材"}})
	if err != nil {
		c.AbortWithStatus(404)
		panic(err)
	} else {
		c.JSON(200, "update sucessed!")
	}
	fmt.Println("update sucessed!")
}

//获取集合元素个数
func Count(c *gin.Context) {
	token := c.Request.Header.Get("token") //获取head中的参数
	logintype := c.Request.Header.Get("logintype")
	fmt.Println("token:",token,"\t\tlogintype:",logintype)
	session, c1 := db.GetConnection("qidian", "service")
	defer session.Close()
	countNum, err := c1.Count()
	if err != nil {
		c.AbortWithStatus(404)
		panic(err)
	} else {
		c.JSON(200, countNum)
	}
	fmt.Println("count:", countNum)
}

//删除元素
func Delete(c *gin.Context) {
	session, c1 := db.GetConnection("qidian", "service")
	defer session.Close()
	_, err := c1.RemoveAll(bson.M{"num": 563})
	if err != nil {
		c.AbortWithStatus(404)
		panic(err)
	} else {
		c.JSON(200, "delete successed!")
	}
	fmt.Println("delete successed!")
}
