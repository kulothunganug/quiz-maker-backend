package models

type Question struct {
	ID              uint     `json:"id" gorm:"primaryKey"`
	QuizID          uint     `json:"quizId"`
	Options         []Option `json:"options" gorm:"foreignKey:QuestionID"`
	CorrectOptionID uint     `json:"correctOptionId"`
	Text            string   `json:"text"`
}
