package router

import (
	"github.com/butga/paketin/handler"
	"github.com/butga/paketin/src/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Controller(db *gorm.DB) {

	// Routing
	router := gin.Default()

	// API Information Routes
	router.GET("", handler.Welcome)

	// Grouping Routes
	version := "/v1"

	// Root Routes
	rootRoutes := router.Group(version + "/api")
	{
		rootRoutes.GET("/health", handler.Health)
	}

	// User Routes
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	userRoutes := router.Group(version + "/api")
	{
		userRoutes.POST("/users", userHandler.PostUserHandler)
		userRoutes.GET("/users", userHandler.GetUsersHandler)
		userRoutes.DELETE("/users/:id", userHandler.DeleteUsersHandler)
	}

	router.Run()
}
