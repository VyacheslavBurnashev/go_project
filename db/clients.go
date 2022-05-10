package db

import (
	"log"

	"github.com/VyacheslavBurnashev/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB

var dbError error

func Connect(connection string) {
	Instance, dbError = gorm.Open(mysql.Open(connection), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Failed connect to db")
	}
	log.Println("Connected to db")

}

func Migrated() {
	Instance.AutoMigrate(&model.User{})
	log.Println("Database Migration Completed")
}
