package logic

import (
	"felixwie.com/producer/models"
	"gorm.io/gorm"
)

var db *gorm.DB

type QueryOptions struct {
	Take   int
	Skip   int
	Select []string
}

type DbModel interface {
	Joins() string
}

func init() {
	db = models.GetDB()
}

func Create[T DbModel](data T) (*T, error) {
	result := db.Create(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil
}

func GetAll[T DbModel](opts *QueryOptions) (*[]T, error) {
	var data []T
	var x T
	result := db.Joins(x.Joins()).Limit(opts.Take).Offset(opts.Skip).Select(opts.Select).Find(&data)

	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil
}

func GetOne[T DbModel](id string) (*T, error) {
	var data T
	result := db.Where("id = ?", id).Joins(data.Joins()).Find(&data)

	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil
}

func Remove[T DbModel](id string) error {
	var x T
	result := db.Where("id = ?", id).Delete(&x)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
