package database

import (
	"log"
	"quiz-maker-backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Instance *gorm.DB
	dbError  error
)

func Connect(databaseName string) {
	Instance, dbError = gorm.Open(sqlite.Open(databaseName), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}

func Migrate() {
	Instance.AutoMigrate(
		&models.Quiz{},
		&models.Question{},
		&models.Option{},
		&models.Submission{})
	log.Println("Database Migration Completed!")
}
