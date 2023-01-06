package model

type Social struct {
	Id       string `json:"social_id" gorm:"column:social_id;primaryKey;autoIncrement:false"`
	Provider string `json:"provider" gorm:"column:provider;primaryKey;autoIncrement:false"`
	Email    string `json:"email" gorm:"column:email"`
	UserId   uint
	User     User
}
