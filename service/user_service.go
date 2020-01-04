package service

import (
	"log"

	"github.com/s-kuramitsu/my-api/entity"
)

// Get user by ID using DB connection receiver
func (con *DBConnection) getUserById(id uint) (*entity.User, []error) {
	user := &entity.User{}
	errors := con.First(user, id).GetErrors()
	return user, errors
}

// Create DB connection and get user by ID
func GetUserById(id uint) *entity.User {
	var con *DBConnection
	if con = getDBConnection(); con == nil {
		log.Print("fail to connect DB")
		return nil
	}
	defer con.Close()

	user, errors := con.getUserById(id)
	if len(errors) > 0 {
		printErrors(errors)
		return nil
	}
	return user
}

// Get user by email using DB connection receiver
func (con *DBConnection) getUserByEmail(email string) (*entity.User, []error) {
	var user = &entity.User{}
	errors := con.Where("email = ?", email).Take(user).GetErrors()
	return user, errors
}

// Create DB connection and get user by email
func GetUserByEmail(email string) *entity.User {
	var con *DBConnection
	if con = getDBConnection(); con == nil {
		log.Print("fail to connect DB")
		return nil
	}
	defer con.Close()

	user, errors := con.getUserByEmail(email)
	if len(errors) > 0 {
		printErrors(errors)
		return nil
	}
	return user
}

// Create DB connection and Create user
func CreateUser(user *entity.User) *entity.User {
	var con *DBConnection
	if con = getDBConnection(); con == nil {
		log.Print("fail to connect DB")
		return nil
	}
	defer con.Close()

	if errors := con.Create(user).GetErrors(); len(errors) > 0 {
		printErrors(errors)
		return nil
	}

	createdUser, errors := con.getUserByEmail(user.Email)
	if len(errors) > 0 {
		printErrors(errors)
		return nil
	}
	return createdUser
}
