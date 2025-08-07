package internal

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("services.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	db.AutoMigrate(&Service{}, &Version{})
	return db
}

func SeedData(db *gorm.DB) {
	var count int64
	db.Model(&Service{}).Count(&count)
	if count > 0 {
		return
	}

	services := []Service{
		{
			Name:        "Auth Service",
			Description: "Handles user authentication",
			Versions: []Version{
				{Number: "v1.0.0"},
				{Number: "v1.1.0"},
			},
		},
		{
			Name:        "Billing Service",
			Description: "Handles invoices and payments",
			Versions: []Version{
				{Number: "v2.0.0"},
			},
		},
	}
	db.Create(&services)
}
