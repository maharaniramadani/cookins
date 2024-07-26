package models

import "gorm.io/gorm"

type SaveResep struct {
	gorm.Model
	ID      int
	IdUser  uint `json:"id_user"`
	IdResep uint `json:"id_resep"`
}