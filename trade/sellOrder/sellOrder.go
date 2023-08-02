package sellOrder

import (
	"ORM.testORM/trade/order"
)

type Sell struct {

	order.Order			`gorm:"embedded;embeddedPrefix:sell_"`
	SellSide			string	`gorm:"default:SELL;embeddedPrefix:s_"`

}