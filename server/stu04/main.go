package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"log"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}
func TodoShow(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	todoId := vars["todoId"]
	fmt.Fprintln(writer, "Todo show:", todoId)
}
func TodoIndex(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Todo Index!",request.URL.Path)
}
func Index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Welcome!")
}