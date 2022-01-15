package models

import (
	"time"

	"gorm.io/gorm"
)

type Metric string

const (
	ML Metric = "ml"
	L         = "l"
	G         = "g"
	KG        = "kg"
)

type Inventory struct {
	gorm.Model
	ID                int             `json:"id" gorm:"primaryKey"`
	Name              string          `json:"name"`
	QuantityAvailable int             `json:"available"`
	QuantityDesired   int             `json:"desired"`
	Refills           []RefillHistory `json:"refills"`
	Metric            Metric          `json:"metric"`

	Active bool `gorm:"-"`
}

type RefillHistory struct {
	gorm.Model
	InventoryID int
	Quantity    int       `json:"quantity"`
	Time        time.Time `json:"time"`
}

func (i Inventory) Joins() string {
	return "left join refill_history on refill_history.inventory_id = inventory.id"
}

func (i Inventory) FilterValue() string {
	return ""
}
