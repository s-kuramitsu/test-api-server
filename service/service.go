package service

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/s-kuramitsu/my-api/db"
)

type DBConnection struct {
	*gorm.DB
}

func getDBConnection() *DBConnection {
	con := &DBConnection{db.GetConnection()}
	return con
}

func printErrors(errors []error) {
	for _, err := range errors {
		log.Print(err)
	}
}
