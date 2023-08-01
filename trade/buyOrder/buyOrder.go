package buyOrder

import (
	"ORM.testORM/trade/order"
)

type Buy struct {

	order.Order			`gorm:"embedded;embeddedPrefix:buy_"`
	Side			string	`gorm:"default:'BUY'"`

}