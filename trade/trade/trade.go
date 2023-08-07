package trade

import (
	"fmt"

	"ORM.testORM/trade/buyOrder"
	"ORM.testORM/trade/sellOrder"
	"gorm.io/gorm"
)

type Trade struct {
	ID 	uint16	`gorm:"primaryKey,autoIncrement"`
	buyOrder.Buy
	sellOrder.Sell	
}

func (t Trade) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Trade{})
}

func (t *Trade) CreateTrade(db *gorm.DB, buy buyOrder.Buy, sell sellOrder.Sell ) (*Trade, error){
	var newTrade  = Trade{
		0,
		buy,
		sell,
	} 
	
	fmt.Println(newTrade)
	// @todo debug here
	if err := db.Create(&newTrade).Error; err != nil {
		return nil, err
	}

	return &newTrade, nil
}