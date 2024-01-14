package models

type Question struct {
	ID      uint     `json:"id" gorm:"primaryKey"`
	QuizID  uint     `json:"quizId"`
	Options []Option `json:"options" gorm:"foreignKey:QuestionID"`
	// don't show correct answers on client-side
	CorrectOptionID uint   `json:"-"`
	Text            string `json:"text"`
}
