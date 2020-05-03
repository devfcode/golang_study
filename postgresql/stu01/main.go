package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"math/rand"
	"time"
	"strconv"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "elm"
	password = "123456"
	dbname   = "postgres"
)

type Student struct {
	Id int `gorm:"primary_key"`
	Name string `gorm:"type:varchar(20);column:name"`
	Gender int `gorm:"type:int;column:gender"`
	Age int `gorm:"type:int;column:age"`
	City string `gorm:"varchar(20);column:city"`
	School string `gorm:"varchar(20);column:school"`
}

func InitDB() (db *gorm.DB) {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = gorm.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	//Query()
	//Insert()
	//Delete()
	Update()
}

func Query()  {
	var err error
	var db *gorm.DB
	db = InitDB()
	defer db.Close()

	var stus []Student
	//err = utils.Last(&stu).Error
	//err = utils.Table("student").Select("id, name, gender, age, city, school").Where("age = ?", 18).Scan(&stu).Error
	err = db.Table("student").Find(&stus).Error
	if err != nil {
		panic(err)
	}
	for _,stu := range stus{
		fmt.Println(stu.Id,stu.Name,stu.Gender,stu.Age,stu.City,stu.School)
	}
}

func Insert()  {
	var err error
	var db *gorm.DB
	db = InitDB()
	defer db.Close()

	rand.Seed(time.Now().Unix())
	rnd1 := rand.Intn(990)
	var nameN string = "Tom" + strconv.Itoa(rnd1)
	var stu Student = Student{
		Id : rnd1,
		Name : nameN,
		Gender : 1,
		Age : 56,
		City : "shanghai",
		School : "qinghua",
	}

	err = db.Table("student").Create(&stu).Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("insert successed !")
}

func Delete()  {
	var err error
	var db *gorm.DB
	db = InitDB()
	defer db.Close()

	err = db.Table("student").Where("id = ?",2).Delete(Student{}).Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("delete successed !")
}

func Update()  {
	var err error
	var db *gorm.DB
	db = InitDB()
	defer db.Close()

	var stu Student
	stu.Id = 6
	var stu1 Student = Student{
		Name : "a new name",
		Gender : 1,
		Age : 56,
		City : "shanghai",
		School : "qinghua",
	}
	err = db.Table("student").Model(&stu).Updates(&stu1).Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("update successed !")
}
