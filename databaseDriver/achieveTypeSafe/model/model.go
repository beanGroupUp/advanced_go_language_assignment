package model

import "gorm.io/gorm"

type Books struct {
	gorm.Model
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

type Book struct {
	gorm.Model
	Book Books `json:"book"`
}
