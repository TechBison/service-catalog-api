package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shreeshg/service-catalog-api/internal"
)

func main() {
	db := internal.InitDB()
	internal.SeedData(db)
	internal.LoadCache(db)

	router := gin.Default()

	router.GET("/services", internal.GetServices)
	router.GET("/services/:id", internal.GetServiceByID)
	router.GET("/services/:id/versions", internal.GetServiceVersions)

	log.Println("Server running on :8080")
	router.Run(":8080")
}
