package main

import (
	"ORM.testORM/product"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&product.Product{})
   
	// Create
	db.Create(&product.Product{Code: "D42", Price: 100})
   
	// Read
	var _product product.Product
	db.First(&_product, 1) // find product with integer primary key
	db.First(&_product, "code = ?", "D42") // find product with code D42
   
	// Update - update product's price to 200
	db.Model(&_product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&_product).Updates(product.Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&_product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
   
	// Delete - delete product
	db.Delete(&_product, 1)
   }