package main

import (
	"fmt"
	"log"
	"math/rand"

	"ORM.testORM/trade/trade"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}

	trad := &trade.Trade{}
	trad.Migrate(db)

	randomPriceTrade, err := trad.CreateTradeRandomPrice(db)
	if err != nil {
		log.Fatal("Error during creating random")
		return
	}
	fmt.Println("\n\nRandom Price Trade: ")
	fmt.Println(randomPriceTrade)

	trad.Buy.Symbol = "BTC/BUSD"
	trad.Buy.Price = (rand.Float64() * 25000) + 15000 // [25_000, 40_000]

	trad.Sell.Symbol = "BTC/BUSD"
	trad.Sell.Price = (rand.Float64() * 25000) + 15000 // [25_000, 40_000]

	newTrad, err := trad.CreateTrade(db, trad.Buy, trad.Sell)
	if err != nil {
		return
	}
	fmt.Println("\n\nNew Trade: ")
	fmt.Println(newTrad)

	tradeRead, err := trad.ReadTrade(db, "2")
	if err != nil{
		log.Fatal("Error during the reading by id")
		return
	}

	fmt.Println("\n\nRead Trade: ")
	fmt.Println(tradeRead)


	deletedTrade, err := trad.Delete(db, "8")
	if err != nil{
		log.Fatal("Error during the delete")
		return
	}

	fmt.Println("\n\nDeleted Trade: ")
	fmt.Println(deletedTrade)
   
	// // Read
	// var _product product.Product
	// db.First(&_product, 1) // find product with integer primary key
	// db.First(&_product, "code = ?", "D42") // find product with code D42
   
	// // Update - update product's price to 200
	// db.Model(&_product).Update("Price", 200)
	// // Update - update multiple fields
	// db.Model(&_product).Updates(product.Product{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&_product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
   
	// // Delete - delete product
	// db.Delete(&_product, 1)
   }