package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		"Hi, This is an example of https service in golang!")
}

func main() {
	http.HandleFunc("/ssltest", handler)
	/*
	 addr : 	服务器 ip ,port
	 certFile :	服务器端的证书(公钥)
	 keyFile :  服务器端的私钥
	*/
	err := http.ListenAndServeTLS("103.100.211.187:8081", "/etc/letsencrypt/live/www.anant.club/cert.pem",
		"/etc/letsencrypt/live/www.anant.club/privkey.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
