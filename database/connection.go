package database

import (
	"github.com/diehernandez/go-auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("diehernandez:pepito.p0@tcp(host.docker.internal:3306)/auth"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	DB = connection
	connection.AutoMigrate(&models.User{})
}
