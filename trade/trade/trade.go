package trade

import (
	"fmt"

	"ORM.testORM/trade/buyOrder"
	"ORM.testORM/trade/sellOrder"
	"gorm.io/gorm"
)

type Trade struct {
	buyOrder.Buy
	sellOrder.Sell	
}

func (t Trade) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Trade{})
}

func (t *Trade) CreateTrade(db *gorm.DB, buy buyOrder.Buy, sell sellOrder.Sell ) (*Trade, error){
	var newTrade  = Trade{
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