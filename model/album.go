package model

import "github.com/jinzhu/gorm"

type Album struct {
	gorm.Model

	ID        string  `json:"id"`
	Title     string  `json:"title"`
	Artist    string  `json:"artist"`
	Price     float64 `json:"price"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
