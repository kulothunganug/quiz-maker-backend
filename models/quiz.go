package models

import (
	"time"
)

type Quiz struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time  `json:"createdAt" gorm:"not null"`
	Title     string     `json:"title" gorm:"not null"`
	CreatedBy string     `json:"createdBy" gorm:"not null"`
	Questions []Question `json:"questions" gorm:"foreignKey:QuizID"`
}
