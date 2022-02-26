package config

import (
	"log"

	"github.com/anandawira/anandapay/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	// Hardcore, later change to env variable
	dsn := "root:example@tcp(127.0.0.1:3306)/anandapay?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&domain.User{})

	DB = db
	return DB
}
