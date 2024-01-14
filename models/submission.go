package models

import "time"

type Submission struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	QuizID     uint      `json:"quizId"`
	Score      uint      `json:"score"`
	AnsweredBy string    `json:"answeredBy"`
	AnsweredAt time.Time `json:"answeredAt"`
}
