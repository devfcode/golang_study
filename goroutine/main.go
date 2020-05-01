package main

import (
	"fmt"
	"time"
	"runtime"
)

func Fly()  {
	for i := 0;i < 100; i++ {
		fmt.Println("fly:\t",i)
	}
}

func Run()  {
	for i := 200;i < 300; i++ {
		fmt.Println("run:\t",i)
	}
}

func main() {
	runtime.GOMAXPROCS(1)   //指定使用cup的一个核
	go  Fly()
	go  Run()
	time.Sleep(5 * time.Second)
}
