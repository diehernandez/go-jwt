package database

import (
	"github.com/diehernandez/go-auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(conUrl *string) {
	connection, err := gorm.Open(mysql.Open(*conUrl), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	DB = connection
	connection.AutoMigrate(&models.User{})
}
