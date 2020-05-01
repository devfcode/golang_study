package main

import (
	"fmt"
)

type Phone interface {
	call()
}

type IPhone struct {
	Company string
	Price float64
}

func (iPhone IPhone) call() {
	fmt.Println("Company:",iPhone.Company,"\tPrice:",iPhone.Price)
	fmt.Println("I am iPhone, I can call you!")
}

type NokiaPhone struct {
	Company string
	Price float64
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("Company:",nokiaPhone.Company,"\tPrice:", nokiaPhone.Price)
	fmt.Println("I am Nokia, I can call you!")
}

func show()  {
	fmt.Println("hello go")
}

func main() {
	var noki NokiaPhone
	noki.Company = "Nokia"
	noki.Price = 50.0
	noki.call()

	var iph IPhone
	iph.Company = "Apple"
	iph.Price = 599.9
	iph.call()

	fmt.Println("***********************************************		分割线   	***************************************************")
	var phone Phone

	phone = noki
	phone.call()

	phone = iph
	phone.call()

	show()
}
