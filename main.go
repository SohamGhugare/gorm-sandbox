package main

import (
	"fmt"

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

	id, err := createProduct(db, "D32", 5)
	fmt.Printf("Product #%d created successfully", id)
}

// Creating new product
func createProduct(db *gorm.DB, code string, price uint) (uint, error) {
	product := Product{
		Code:  code,
		Price: price,
	}
	result := db.Create(&product)
	if result.Error != nil {
		return 0, result.Error
	}
	return product.ID, nil

}
