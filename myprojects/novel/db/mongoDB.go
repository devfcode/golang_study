package db

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"time"
	"log"
)

//数据库连接配置
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


const (
	url_real  = "61.14.253.155:27017"
	username = "qidian"
	password = "qidian_pw_2035"
)
func GetConnection2(db string, collection string) (s *mgo.Session, c *mgo.Collection) {
	var dialInfo *mgo.DialInfo = &mgo.DialInfo{
		Addrs:     []string{url_real},
		Direct:    false,
		Timeout:   time.Second * 5,
		Database:  db, //数据库名字
		//Source:    collection,
		Username:  username, //数据库的用户名
		Password:  password, //数据库的用户密码
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
