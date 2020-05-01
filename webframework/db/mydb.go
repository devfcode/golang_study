package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "elm"
	password = "123456"
	dbname   = "postgres"
)


func Init() *gorm.DB {
	var db *gorm.DB
	var err error

	var psqlInfo string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("postgresql was successfully connected!")
	return db
}
