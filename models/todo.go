package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Name        string `gorm:"not null" json:"name"`
	Description string `json:"description"`
	Status      bool   `gorm:"default:false" json:"status"`
}
