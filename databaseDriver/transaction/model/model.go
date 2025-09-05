package model

import "gorm.io/gorm"

type Accounts struct {
	gorm.Model
	Balance float64
}

type Transactions struct {
	gorm.Model
	FromAccountId uint
	ToAccountId   uint
	Amount        float64
}
