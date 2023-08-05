package trade

import (
	"errors"
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
func (t *Trade) ReadTradeAll(db *gorm.DB, trade Trade) (*Trade, error) {
	var emptyCondition string
	readTradeAll := db.Find(&trade, emptyCondition)
	fmt.Println(readTradeAll.RowsAffected)
	return nil, nil
}

// Update
func (t *Trade) UpdateTrade(db *gorm.DB, trade Trade) (bool, error) {
	if trade.ID > 0{
		if trade.Buy.Symbol == trade.Sell.Symbol {
			if trade.Buy.Price > 0.0 && trade.Sell.Price > 0.0 {
				fmt.Println("##### trade fields are valid")
			} else {
				return false, errors.New("##### buy or sell price is negative check tha value")
			}
		} else {
			return false, errors.New("##### buy symbol and sell symbol not equal, the buy is not connected with sell")
		}
	} else {
		return false, errors.New("##### ID is missing or with a not valid value, inserted value could be id < 0 or id = 0")
	}
	db.Save(&Trade{ID: trade.ID, Buy: trade.Buy, Sell: trade.Sell})
	return true, nil
}

// Delete
func (t *Trade) Delete(db *gorm.DB, id string) (*Trade, error){
	var deleteTrade Trade
	if err := db.Delete(&deleteTrade, id).Error; err != nil{
		return nil, err
	}
	return &deleteTrade, nil
}

