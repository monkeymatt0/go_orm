package sellOrder

import (
	"ORM.testORM/trade/order"
)

type Sell struct {

	order.Order			`gorm:"embedded;embeddedPrefix:sell_"`
	Side			string	`gorm:"default:'SELL'"`

}