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

	// Grouping
	rootRoutes := router.Group(version+"/api")
	{
		rootRoutes.GET("/health", handler.Health)
	}

	router.Run()
}