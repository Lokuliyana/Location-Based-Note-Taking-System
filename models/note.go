package models

import "github.com/jinzhu/gorm"

type Note struct {
	gorm.Model
	UserID    uint    `json:"userId"`
	Title     string  `json:"title"`
	Content   string  `json:"content"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
