package db

import (
	"log"
	"slack-payment-request-bot/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(connectionString string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Request{})

	return db
}
