package database

import (
	"log"
	"quiz-maker-backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	Instance *gorm.DB
}

func New(databaseName string) *Database {
	log.Println("Initiating Database!")
	db, err := gorm.Open(sqlite.Open(databaseName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
	return &Database{Instance: db}
}

func (db *Database) Migrate() {
	db.Instance.AutoMigrate(
		&models.Quiz{},
		&models.Question{},
		&models.Option{},
		&models.Submission{})
	log.Println("Database Migration Completed!")
}
