package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID int
	Name     string `json:"name" form:"name" validate:"require"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}