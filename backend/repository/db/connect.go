package repository

import (
	model "ybigta/bard-backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type Database struct {
	dsn string
}

func NewDatabase(dsn string) *Database {
	return &Database{
		dsn: dsn,
	}
}

func (d *Database) Connect() {

	db, err := gorm.Open(mysql.Open(d.dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.User{}, &model.Social{}, &model.File{}, &model.Story{})
	if err != nil {
		panic(err)
	}

	mysqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	mysqlDB.SetMaxIdleConns(10)
	mysqlDB.SetMaxOpenConns(100)

	DB = db
}
