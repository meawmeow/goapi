package entity

import "gorm.io/gorm"

type Orders struct {
	gorm.Model
	Count     int
	NetPrices uint64
	ProductID uint64
	Product   Product
	UserID    uint64
	User      User
}
