package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	controller "go-restful-api/controllers"
)

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/", controller.Home)
	route.HandleFunc("/add-user", controller.AddUser)

	log.Fatal(http.ListenAndServe(":9090", route))
}
