package main

import (
	"log"
	"net/http"

	"rest-api-go-mysql/db"
	"rest-api-go-mysql/route"
	"rest-api-go-mysql/service"
)

func main() {
	var database = db.GetConnection()
	service.SetDB(database)

	var router = route.CreateRouter()

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
