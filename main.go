package main

import (
	"github.com/butga/paketin/handler"
	"github.com/gin-gonic/gin"
)

func main()  {
	// Routing
	router := gin.Default()

	// Versioning
	version := "/v1"

	// API Information Routes
	router.GET("", handler.Welcome)

	// Grouping Routes
	rootRoutes := router.Group(version+"/api")
	{
		rootRoutes.GET("/health", handler.Health)
	}

	router.Run()
}