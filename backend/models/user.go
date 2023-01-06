package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `gorm:"not null"`
	Name  string `gorm:"not null"`
}
