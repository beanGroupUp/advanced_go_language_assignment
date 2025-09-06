package model

import "gorm.io/gorm"

type Employees struct {
	gorm.Model
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}
