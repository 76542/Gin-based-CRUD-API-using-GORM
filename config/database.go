package config

import (
	"example/Crud/logger"
	"example/Crud/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// declaring global var ad connecting it to a existing database
// DB holds the database connection
var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:mysql_password@tcp(127.0.0.1:3308)/MyDatabase?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.CustomLogger.Error("Failed to connect to the database: ", err)
		log.Fatal("Failed to connect to database: ", err)
	}

	db.AutoMigrate(&models.Item{})

	DB = db //db represents a database connection object that indicates a unique session with a datasource
	logger.CustomLogger.Info("Database connected successfully")
}
