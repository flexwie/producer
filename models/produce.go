package models

import "gorm.io/gorm"

type Produce struct {
	gorm.Model
	Name        string      `json:"name"`
	Ingredients []Inventory `json:"ingredients" gorm:"many2many;inventory_produce"`
}

func (p Produce) Joins() string {
	return ""
}
