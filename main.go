package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Product schema
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// Opening a connetion
	db, err := gorm.Open(sqlite.Open("products.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	// Migrating schema
	db.AutoMigrate(&Product{})

}
