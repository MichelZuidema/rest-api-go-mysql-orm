package service

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
	"rest-api-go-mysql/model"
)

var dbconn *gorm.DB

type Response struct {
	Data []model.Book `json:"data"`
	Message string `json:"message"`
}

func GetAllBooks(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	books := model.GetBooks()
	var resp Response

	err := dbconn.Find(&books).Error
	if err == nil {
		resp.Data = books
		resp.Message = "SUCCESS"

		json.NewEncoder(rw).Encode(&resp)
	} else {
		log.Println(err)
		http.Error(rw, err.Error(), 400)
	}
}

func GetBook(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	id, _ := mux.Vars(r)["id"]

	var resp Response
	var book = model.GetBook()
	err := dbconn.Where("id = ?", id).Find(&book).Error

	if err == nil {
		resp.Data = append(resp.Data, book)
		resp.Message = "SUCCESS"

		json.NewEncoder(rw).Encode(&resp)
	} else {
		http.Error(rw, err.Error(), 400)
	}
}

func CreateBook(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	var resp Response
	var book = model.GetBook()
	_ = json.NewDecoder(r.Body).Decode(&book)

	err := dbconn.Create(&book).Error
	if err != nil {
		http.Error(rw, "Something went wrong...", 400)
		return
	}

	resp.Message = "CREATED"
	json.NewEncoder(rw).Encode(resp)
}

func DeleteBook(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var resp Response
	var book = model.GetBook()

	err := dbconn.Delete(&book, id).Error
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	resp.Message = "DELETED"
	json.NewEncoder(rw).Encode(resp)
}

func UpdateBook(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var resp Response
	var book = model.GetBook()

	err := dbconn.Delete(&book, id).Error
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	resp.Message = "UPDATED"
	resp.Data = append(resp.Data, book)

	json.NewEncoder(rw).Encode(resp)
}

func Index(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello, World!")
}

func SetDB(db *gorm.DB) {
	dbconn = db

	var book = model.GetBook()
	dbconn.AutoMigrate(book)
}