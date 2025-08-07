package internal

import (
	"io"
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
			Name: "Locate Us",	
			Description: "Sends location data to user",
			Versions: []Version{
				{Number: "v1.0.0"},
				{Number: "v1.1.0"},
				{Number: "v1.2.0"},
				{Number: "v1.3.0"}
			},
		},
		{
			Name: "Order Service",
			Description: "Handles order processing",
			Versions: []Version{
				{Number: "v1.0.0"},
				{Number: "v1.1.0"},
				{Number: "v1.2.0"},
			},
		},
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
		{
			Name: "Payment Service",
			Description: "Handles payment processing",
			Versions: []Version{
				{Number: "v1.0.0"},
				{Number: "v1.2.0"},
			},
		},
		{
			Name: "Notification Service",
			Description: "Handles notifications",
			Versions: []Version{
				{Number: "v1.0.0"},
				{Number: "v1.2.0"},
			},
		},
		{
			Name: "Reporting Service",
			Description: "Handles reporting",
			Versions: []Version{
				{Number: "v1.0.0"},
				{Number: "v1.1.0"},
			},
		},
		{
			Name: "User Service",
			Description: "Handles user management",
			Versions: []Version{
				{Number: "v1.0.0"},
				{Number: "v1.1.0"},
			},
		},
		{
			Name: "Contact Service",
			Description: "Handles contact requests",
			Versions: []Version{
				{Number: "v1.0.0"},
				{Number: "v1.1.0"},
				{Number: "v1.2.0"},
			},
		},
		{
			Name: "Security Service",
			Description: "Handles security",
			Versions: []Version{
				{Number: "v1.0.0"},
				{Number: "v1.1.0"},
			},
		},
	}
	db.Create(&services)
}
