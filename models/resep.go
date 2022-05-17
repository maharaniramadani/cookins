package models

import "gorm.io/gorm"

type Resep struct {
	gorm.Model
	ID          int
	Title       string   `json:"title"`
	Ingredients []string `json:"ingredients"`
	Link        string   `json:"link"`
	Gif         string   `json:"gif"`
}
