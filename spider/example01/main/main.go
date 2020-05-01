package main

import (
	"golang_study/spider/example01/action"
	"log"
)

//https://www.cnblogs.com/majianguo/p/8186429.html
func main() {
	var err error
	err = action.GetStockListA("spider/resourse/sseA.csv")
	if err != nil {
		log.Fatal(err)
	}
	err = action.GetStockListB("spider/resourse/sseB.csv")
	if err != nil {
		log.Fatal(err)
	}
}
