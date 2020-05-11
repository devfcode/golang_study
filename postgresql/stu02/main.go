package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"math/rand"
	"time"
	"strconv"
	"log"
)

const (
	host     = "103.100.211.187"
	port     = 5432
	user     = "postgres"
	password = "1qaz2wsx"
	dbname   = "postgres"
)

var err error

func InitDB() (db *sql.DB)  {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//defer utils.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return db
}

type Student struct {
	Id int
	Name string
	Gender int
	Age int
	City string
	School string
}

//直接使用sql操作数据库,在项目中不建议使用
func main() {
	//QueryOne()
	QueryMany()
	//Insert()
	//Update()
	//Delete()
}

func QueryOne()  {
	var db  *sql.DB
	db = InitDB()
	defer db.Close()

	var sqlStatement string
	sqlStatement = `SELECT * FROM student where id=$1;`
	row := db.QueryRow(sqlStatement, 9)
	var stu Student
	err = row.Scan(&stu.Id, &stu.Name,&stu.Gender,&stu.Age,&stu.City,&stu.School)

	if err != nil {//异常处理
		if err == sql.ErrNoRows {//在查一行时，如果没数据也被视为异常
			fmt.Println("no value in database!")
			return
		}else {
			log.Fatal(err)
		}
	}
	fmt.Println(stu)
}

func QueryMany()  {
	var db  *sql.DB
	db = InitDB()
	defer db.Close()

	var sqlStr string
	sqlStr =  `SELECT * FROM student  where age>$1;`
	var rows *sql.Rows
	rows, err = db.Query(sqlStr, 19)
	if err != nil {
		panic(err)
	}

	var stus []Student
	var stu1 Student
	for rows.Next() {
		err = rows.Scan(&stu1.Id, &stu1.Name, &stu1.Gender, &stu1.Age, &stu1.City, &stu1.School)
		if err != nil {
			panic(err)
		}
		stus = append(stus, stu1)
		//fmt.Println(stu1)
	}
	if rows.Err() != nil {
		panic(err)
	}

	for _, stu := range stus {
		fmt.Println(stu)
	}
}

func Insert()  {
	var db  *sql.DB
	db = InitDB()
	defer db.Close()

	stmt, err1 := db.Prepare(`INSERT INTO student(name, gender,age,city,school) VALUES ($1, $2, $3, $4, $5);`)
	if err1 != nil {
		panic(err1)
	}

	rand.Seed(time.Now().Unix())
	rnd1 := rand.Intn(990)
	var nameN string = "liying" + strconv.Itoa(rnd1)
	_, err2 := stmt.Exec(nameN,1,19,"jianpuzhai","New York")
	if err2 != nil {
		panic(err2)
	}
	fmt.Println("Insert successed!")
}

func Update()  {
	var db  *sql.DB
	db = InitDB()
	defer db.Close()

	stmt, err1 := db.Prepare(`UPDATE student set name = $1, age = $2 where id = $3;`)
	if err1 != nil {
		panic(err1)
	}
	_, err2 := stmt.Exec("elm",999,7)
	if err2 != nil {
		 panic(err2)
	}
	fmt.Println("Update successed!")
}

func Delete()  {
	var db  *sql.DB
	db = InitDB()
	defer db.Close()

	stmt, err1 := db.Prepare(`delete from student where id = $1;`)
	if err1 != nil {
		panic(err1)
	}
	_, err2 := stmt.Exec(3)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println("Delete successed!")
}
