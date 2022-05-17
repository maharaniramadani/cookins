package models

import "gorm.io/gorm"

type DetailResep struct {
	gorm.Model
	ID          int
	IdResep     int
	Title       string
	Deskription string
}