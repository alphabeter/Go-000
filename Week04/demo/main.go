package main

import (
	"github.com/alphabeter/Go-000/demo/model"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test/test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{})
	db.Create(&model.User{Name: "A1001", Age: 43})
	db.Create(&model.User{Name: "B1001", Age: 33})
	db.Create(&model.User{Name: "C1001", Age: 32})
	return db
}

func main() {
	db := initDatabase()

	userAPI := InitUserAPI(db)
	r := gin.Default()
	r.GET("/users", userAPI.FindAll)

	if err := r.Run(); err != nil {
		panic(err)
	}
}
