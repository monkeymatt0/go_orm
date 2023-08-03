package trade

import (
	"fmt"
	"math/rand"

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

// Create
func (t *Trade) CreateTrade(db *gorm.DB, buy buyOrder.Buy, sell sellOrder.Sell) (*Trade, error){
	var newTrade  = Trade{
		0,
		buy,
		sell,
	} 
	
	if err := db.Create(&newTrade).Error; err != nil {
		return nil, err
	}

	return &newTrade, nil
}

func (t *Trade) CreateTradeRandomPrice(db *gorm.DB) (*Trade, error){

	var newTrade Trade
	newTrade.Buy.Symbol = "BTC/BUSD"
	newTrade.Buy.Price = (rand.Float64() * 25000) + 15000 // [25_000, 40_000]

	newTrade.Sell.Symbol = "BTC/BUSD"
	newTrade.Sell.Price = (rand.Float64() * 25000) + 15000 // [25_000, 40_000]

	return &newTrade, nil
}

// Read
func (t *Trade) ReadTrade(db *gorm.DB, id string) (*Trade, error){ 
	var readTrade Trade
	if err := db.First(&readTrade, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &readTrade, nil
}
func (t *Trade) ReadTradeAll(db *gorm.DB, trade Trade) { 
	readTradeAll := db.Find(&trade)
	fmt.Println(readTradeAll.RowsAffected)
}

// Delete
func (t *Trade) Delete(db *gorm.DB, id string) (*Trade, error){
	var deleteTrade Trade
	if err := db.Delete(&deleteTrade, id).Error; err != nil{
		return nil, err
	}
	return &deleteTrade, nil
}

