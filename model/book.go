package model

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title string `gorm:"type:varchar(255);" json:"title"`
	Author string `gorm:"type: varchar(100);" json:"author"`
	Description string `gorm:"type: text" json:"description"`
}

func GetBook() Book {
	var book Book
	return book
}

func GetBooks() []Book {
	var books []Book
	return books
}