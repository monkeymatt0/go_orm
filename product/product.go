package product

import (
	"gorm.io/gorm"
)

// gorm.Model definition
type Product struct {
	gorm.Model
	Code  string
	Price uint
}