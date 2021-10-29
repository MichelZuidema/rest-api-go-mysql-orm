package route

import (
	"github.com/gorilla/mux"
	"rest-api-go-mysql/service"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", service.Index).Methods("GET")
	router.HandleFunc("/books", service.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/", service.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", service.GetBook).Methods("GET")
	router.HandleFunc("/books/{id}", service.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", service.DeleteBook).Methods("DELETE")
	router.HandleFunc("/books/add", service.CreateBook).Methods("POST")

	return router
}