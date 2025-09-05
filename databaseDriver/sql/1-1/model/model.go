package model

import "gorm.io/gorm"

type Students struct {
	gorm.Model
	Name  string
	Age   int
	Grade string
}
