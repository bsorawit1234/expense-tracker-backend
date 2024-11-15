package config

import (
	"log"

	"github.com/bsorawit1234/expense-tracker-backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:admin1234@tcp(127.0.0.1:3306)/expense_tracker?charset=utf8mb4&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	database.AutoMigrate(&models.User{}, &models.Expense{})

	DB = database
}
