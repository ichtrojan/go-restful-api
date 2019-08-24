package main

import (
	"go-restful-api/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	controller "go-restful-api/controllers"
)

func main() {
	database.Init()

	route := mux.NewRouter()

	route.HandleFunc("/", controller.Home)
	route.HandleFunc("/add-user", controller.AddUser)

	log.Fatal(http.ListenAndServe(":9090", route))
}
