package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	"log"
	"fmt"
	"strconv"
)

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

//书籍管理
type BookList struct {
	Id bson.ObjectId `bson:"_id" json:"id"`
	Category string `bson:"category" json:"category"`
	SubCategory string `bson:"sub_category" json:"sub_category"`
	Name string	`bson:"name" json:"name"`
	Author string `bson:"author" json:"author"`
	WordsCount int64 `bson:"words_count" json:"words_count"`
	ChapterCount int `bson:"chapter_count" json:"chapter_count"`
	Free bool `bson:"free" json:"free"`
	ChapterPrice float64 `bson:"chapter_price" json:"chapter_price"`
	FreeChapterCount int `bson:"free_chapter_count" json:"free_chapter_count"`
	CreateAt string `bson:"create_at" json:"create_at"`
	ChapterList []ChapterList `bson:"chapter_list" json:"chapter_list"`
}
type ChapterList struct {
	Free bool `bson:"free" json:"free"`
	Words string `bson:"words" json:"words"`
}

func main() {
	session, c1 := GetConnection2("qidian", "book")
	defer session.Close()
	//var data []interface{}
	var d1 []BookList
	err := c1.Find(bson.M{"name": "混沌幽莲空间"}).
		Select(bson.M{"_id": 1, "category": 1, "sub_category": 1, "name": 1, "author": 1, "free": 1, "create_at": 1, "chapter_list": 1}).
		All(&d1)
	if err != nil {
		panic(err)
	}
	var i int
	for i = 0; i < len(d1); i++ {
		var num1 int = 0
		var num2 int64 = 0
		for _, v2 := range d1[i].ChapterList {
			if v2.Free {
				num1 ++
			}
			tempNum, err3 := strconv.ParseInt(v2.Words, 10, 64)
			if err3 != nil {
				log.Fatal(err3)
			}else {
				num2 += tempNum
			}
			fmt.Println(num2)
		}
		d1[i].FreeChapterCount = num1
		d1[i].WordsCount = num2

		//fmt.Println(v.ChapterList)
		//fmt.Println(reflect.TypeOf(v))

		//json str 转struct
		//var config BookList
		//if err := json.Unmarshal(v, &config); err == nil {
		//	fmt.Println("================json str 转struct==")
		//	fmt.Println(config)
		//	fmt.Println(config.Name)
		//}
	}

	fmt.Println(d1[0])


	//var d2 []ChapterList
	//err1 := c1.Find(bson.M{"name": "混沌幽莲空间"}).
	//	Select(bson.M{"chapter_list": 1}).
	//	All(&d2)
	//if err1 != nil {
	//	panic(err1)
	//}
	//for _, v := range d2 {
	//	fmt.Println(v.Name,"\t", v)
	//}

	//var d3 []ChapterList
	//err2 := c1.Find(bson.M{"name": "混沌幽莲空间","chapter_list":bson.M{"$elemMatch":bson.M{"free":true}}}).All(&d3)
	//if err2 != nil {
	//	panic(err2)
	//}
	//for _, v := range d3 {
	//	fmt.Println(v)
	//}
}
