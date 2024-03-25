package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/solnsumei/simple-chat/utils"

	// using sqlite3 db
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB variable
var DB *gorm.DB

// ConnectDatabase -- connect to database
func ConnectDatabase(config *utils.Config) error {
	db, err := gorm.Open("sqlite3", config.DBName)

	if err != nil {
		return err
	}

	DB = db
	return nil
}

// RunMigration -- run database migration
func RunMigration(config *utils.Config) error {
	err := ConnectDatabase(config)
	if err != nil {
		log.Fatal(err)
		return err
	}

	DB.AutoMigrate(&User{}, &Chat{}, &Message{})
	fmt.Println("Database tables created successfully")

	return nil
}
