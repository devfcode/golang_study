package main

import (
	"net/http"
	"fmt"
	"log"
)

type MyMux struct {}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/say" {
		fmt.Println("say")
		sayhelloName(w, r)
		return
	}else if r.URL.Path == "/run" {
		fmt.Println("run")
		run(w,r)
		return
	}

	http.NotFound(w, r)
	return
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("path:", r.URL.Path)
	fmt.Fprintf(w, "hello go this is a test sayhelloName")
}

func run(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	fmt.Println("path:", r.URL.Path)
	fmt.Fprintf(w,"this a go program  run")
}

func main() {
	mux := &MyMux{}
	err := http.ListenAndServe(":10001", mux)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
