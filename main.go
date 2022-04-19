package main

import (
	"github.com/butga/paketin/config"
	"github.com/butga/paketin/router"
	"gorm.io/gorm"
)

var db *gorm.DB = config.SetupDatabaseConnection()

func main() {
	// Sandbox 19-04-2022

	defer config.CloseDatabaseConnection(db)

	router.Controller(db)
}
