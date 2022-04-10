package main

import (
	"github.com/butga/paketin/handler"
	"github.com/gin-gonic/gin"
)

var (
	//db *gorm.DB = config.SetupDatabaseConnection()
)

func main() {
	//defer config.CloseDatabaseConnection(db)

	// Routing
	router := gin.Default()

	// API Information Routes
	router.GET("", handler.Welcome)

	// Grouping Routes
	version := "/v1"
	rootRoutes := router.Group(version + "/api")
	{
		rootRoutes.GET("/health", handler.Health)
	}

	router.Run()
}
