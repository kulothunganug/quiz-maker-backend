package models

type Option struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	QuestionID uint   `json:"questionId"`
	Text       string `json:"text"`
}
