package config

import (
	"fmt"
	"log"
	authModel "meetspace_backend/auth/models"
	"meetspace_backend/chat/models"
	userModel "meetspace_backend/user/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() *gorm.DB {

  dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
	os.Getenv("DB_HOST"),
	os.Getenv("DB_USERNAME"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_NAME"),
	os.Getenv("DB_PORT"),
	os.Getenv("DB_SSLMODE"))

  	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default,
	})
	if err != nil{
		log.Fatal("failed to connect to db")
	}
	// db.Logger = logger.Default.LogMode(logger.Info)

	if err != nil { 
		panic(err)
	}

	db.AutoMigrate(&authModel.Verification{})
	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&models.ChatMessage{})
	db.AutoMigrate(&models.ChatRoom{})

	DB = db
	return db
}

func GetDB() *gorm.DB {
    return DB
}