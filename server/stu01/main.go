package main

import (
	"net/http"
	"log"
	"fmt"
)

type Listener struct {
	Code int
	Msg string
}

func (p *Listener) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path == "/say" {
		fmt.Fprintf(w, "this is a test, you can say")
		fmt.Println("say")
		return
	}else if r.URL.Path == "/run" {
		fmt.Fprintf(w,"this is a test, you can run")
		fmt.Println("run")
		return
	}
}

func main() {
	var listener Listener
	listener.Code = 200
	listener.Msg = "listening"

	err := http.ListenAndServe("127.0.0.1:10002", &listener)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
