package main

import (
	"gopkg.in/mgo.v2"
	"time"
	"log"
	"fmt"
	"gopkg.in/mgo.v2/bson"
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


const url  = "localhost:27017"

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


func main() {
	session, c1 := GetConnection("qidian", "book")
	defer session.Close()
	err := c1.Update(bson.M{"_id":bson.ObjectIdHex("5b07db39e1382377e755bfa1")}, bson.M{"$set":bson.M{"chapter_list.1.name":"病假条"}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("success!")
}
