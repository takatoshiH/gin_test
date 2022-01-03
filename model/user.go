package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model

	ID        string `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}