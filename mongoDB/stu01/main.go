package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
	"time"
	"log"
	"gopkg.in/mgo.v2"
)

type Service struct {
	Num int `bson:"num"`
	UserId int `bson:"user_id"`
	Account string `bson:"account"`
	Mark string `bson:"mark"`
	Problem string `bson:"problem"`
	Status string `bson:"status"`
	IpAddr string `bson:"ip_address"`
	CreateTime string `bson:"create_time"`
}

const (
	url  = "localhost:27017" //数据库连接配置
	mydb = "testgo"
	collection = "service"
)


func GetConnection(db string, collection string) (*mgo.Session, *mgo.Collection) {
	var session  *mgo.Session
	var err error
	session, err= mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	//defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	var c *mgo.Collection
	c = session.DB(db).C(collection)
	fmt.Println("MongoDB:connected successfully!")
	return session,c
}

func GetConnection2(db string, collection string) (s *mgo.Session, c *mgo.Collection) {
	var url2  = "61.14.253.155:27017"
	dialInfo := &mgo.DialInfo{
		Addrs:     []string{url2},
		Direct:    false,
		Timeout:   time.Second * 2,
		Database:  "qidian", //数据库名字
		//Source:    collection,
		Username:  "qidian", //数据库的用户名
		Password:  "qidian_pw_2035", //数据库的用户密码
		PoolLimit: 4096, // 连接池
	}
	var err error
	s, err= mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatal(err)
	}
	//defer session.Close()
	s.SetMode(mgo.Monotonic, true)
	c = s.DB(db).C(collection)
	fmt.Println("MongoDB service :connected successfully!")
	return s,c
}

func main() {
	session, c1 := GetConnection2("qidian", "book")
	defer session.Close()
	//查询总数据量
	countNum, err1 := c1.Find(bson.M{"name":bson.M{"$regex":"混沌幽莲空*"}}).Count()
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println(countNum)

	//Insert()
	//QueryOne()
	//QueryMany()
	//Update()
	//Count()
	//Delete()
}

//插入数据
func Insert() {
	rand.Seed(time.Now().Unix())
	rnd1 := rand.Intn(990)
	rnd2 := rand.Intn(990)

	temp := Service{
		Num:        rnd1,
		UserId:     109095,
		Account:    "游客",
		Mark:       "穷人",
		Problem:    "Nice to meet you",
		Status:     "未处理",
		IpAddr:     "192.168.10.3(黑龙江)",
		CreateTime: "2017/4/17 12:56",
	}
	temp2 := Service{
		Num:        rnd2,
		UserId:     109095,
		Account:    "游客",
		Mark:       "标记",
		Problem:    "Nice to meet you",
		Status:     "未处理",
		IpAddr:     "192.168.10.4(湖北)",
		CreateTime: "2017/4/17 23:56",
	}
	session, c1 := GetConnection(mydb, collection)
	defer session.Close()
	var err error
	err = c1.Insert(temp, temp2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("insert successed!")
}

//查询单条数据
func QueryOne() {
	session, c1 := GetConnection(mydb, collection)
	defer session.Close()

	var result Service
	err := c1.Find(bson.M{"mark": "标记"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("success:", result)
}

//查询多条数据
func QueryMany() {
	session, c1 := GetConnection(mydb, collection)
	defer session.Close()

	var results  []Service
	err := c1.Find(bson.M{"num" : bson.M{"$lt":580}}).All(&results)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)
}

//更新数据
func Update() {
	session, c1 := GetConnection(mydb, collection)
	defer session.Close()

	err := c1.Update(bson.M{"num": 223}, bson.M{"$set": bson.M{"mark": "蠢材"}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("update sucessed!")
}

//获取集合元素个数
func Count() {
	session, c1 := GetConnection(mydb, collection)
	defer session.Close()
	countNum, err := c1.Count()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("count:", countNum)
}

//删除元素
func Delete() {
	session, c1 := GetConnection(mydb, collection)
	defer session.Close()
	_, err := c1.RemoveAll(bson.M{"num": 223})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("delete successed!")
}
