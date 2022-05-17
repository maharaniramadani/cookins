package models

import (
	"time"

	"gorm.io/gorm"
)

type Bahan struct {
	gorm.Model
	ID int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IdUser      int 	`json:"id_user" form:"id_user"`
	Ingredients string	`json:"ingredients" form:"ingredients"`
}