package models

import "gorm.io/gorm"

type UserLogin struct {
	gorm.Model
	ID int
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}