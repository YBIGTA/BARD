package model

import "gorm.io/gorm"

type Story struct {
	gorm.Model

	Title          string `json:"title" gorm:"column:title"`
	Body           string `json:"body" gorm:"column:body"`
	SummarizedBody string `json:"summarized_body" gorm:"column:summarized_body"`
	ImageUrl       string `json:"image_url" gorm:"column:image_url"`
	UserId         uint
	User           User
}
