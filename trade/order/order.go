package order

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model // Here we have the primaryKey from gorm.model and also the date
	Symbol		string 
	Price		float64 
}