package main

import (
	"github.com/butga/paketin/config"
	"github.com/butga/paketin/handler"
	"github.com/butga/paketin/src/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB = config.SetupDatabaseConnection()
	userRepository          = user.NewRepository(db)
	userService             = user.NewService(userRepository)
	userHandler             = handler.NewUserHandler(userService)
)

func main() {
	defer config.CloseDatabaseConnection(db)

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

	userRoutes := router.Group(version + "/api")
	{
		userRoutes.POST("/users", userHandler.PostUserHandler)
	}

	router.Run()
}
