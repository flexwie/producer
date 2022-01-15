package models

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	if db != nil {
		return db
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Panicf("could not open db: %v", err)
	}

	db.AutoMigrate(&RefillHistory{})
	db.AutoMigrate(&Inventory{})
	db.AutoMigrate(&Produce{})

	return db
}
