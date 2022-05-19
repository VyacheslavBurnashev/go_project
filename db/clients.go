package db

import (
	"fmt"
	"log"

	"github.com/VyacheslavBurnashev/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func (connection *Database) Connect(config *config.Config, logger *log.Logger) {
	db, err := gorm.Open(mysql.Open("mysql", fmt.Sprintf("root:rootroot@tcp(localhost:3306/users@parseTime=true")), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Failed connect to db")
	}
	log.Println("Connected to db")

}
