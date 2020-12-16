package main

import (
	"github.com/alphabeter/Go-000/demo/model"
	"os"

	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	db, err := gorm.Open("mysql", os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{})
	return db
}

func main() {
	db := initDB()
}
