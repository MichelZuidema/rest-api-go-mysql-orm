package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	host = "127.0.0.1"
	port = 3306
	user = "root"
	password = "Kaas@360!420"
	dbname = "goapi"
)

func GetConnection() *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the database")
	}

	fmt.Println("DB Connection established...")

	return db
}