package utils

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
)

const (
	host     = "47.244.33.7"
	port     = 5432
	user     = "dbuser"
	password = "132132"
	dbname   = "mydb"
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
