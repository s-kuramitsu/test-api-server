package main

import (
	"github.com/gin-contrib/location"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/s-kuramitsu/my-api/db"
	"github.com/s-kuramitsu/my-api/web/router"
)

func main() {
}

func init() {
	initDB()
	initRouter()
}

// Initialize database
func initDB() {
	var con = db.GetConnection()
	db.Migrate(con)
	con.Close()
}

// Initialize web router
func initRouter() {
	r := gin.New()
	r.Use(location.New(initLocation()))
	r.Use(logger.SetLogger())
	router.Router(r)
	router.UserRouter(r)
	r.Run()
}

// Initialize location information
func initLocation() location.Config {
	return location.DefaultConfig()
}
