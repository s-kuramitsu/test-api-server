package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/s-kuramitsu/my-api/entity"
)

// Connect DB
func GetConnection() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./test.db")
	if err != nil {
		log.Println(err)
		return nil
	}
	return db
}

// Migrate DB
func Migrate(db *gorm.DB) {
	db.AutoMigrate(entity.User{}, entity.Area{}, entity.CheckList{})
}
