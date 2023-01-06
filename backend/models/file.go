package model

import "gorm.io/gorm"

type File struct {
	gorm.Model
	Name    string `json:"name" gorm:"not null"`
	Url     string `json:"url" gorm:"not null"`
	Size    int64  `json:"size" gorm:"not null"`
	Caption string `json:"caption"`
}
