package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
