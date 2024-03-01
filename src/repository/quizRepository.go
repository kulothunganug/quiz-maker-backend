package repository

import (
	"errors"
	"math/rand"
	"quiz-maker-backend/models"
	"quiz-maker-backend/schemas"

	"gorm.io/gorm"
)

// should check for if the id already exists or not
// if exists generate a new id
func RandomUint() uint {
	return uint(rand.Uint32())
}

func CreateQuiz(dbInstance *gorm.DB, quizReq schemas.Quiz) (uint, error) {
	var quiz models.Quiz
	quiz.ID = RandomUint()
	quiz.Title = quizReq.Title
	quiz.CreatedBy = quizReq.CreatedBy

	for _, questionReq := range quizReq.Questions {
		var question models.Question
		question.ID = RandomUint()
		question.Text = questionReq.Text
		question.QuizID = quiz.ID

		for idx, optionText := range questionReq.Options {
			var option models.Option
			option.ID = RandomUint()
			option.Text = optionText
			option.QuestionID = question.ID
			question.Options = append(question.Options, option)

			if uint(idx) == questionReq.CorrectOptionIdx {
				question.CorrectOptionID = option.ID
			}
		}

		quiz.Questions = append(quiz.Questions, question)
	}

	record := dbInstance.Create(&quiz)
	if record.Error != nil {
		return 0, errors.New(record.Error.Error())
	}

	return quiz.ID, nil
}

func GetQuiz(dbInstance *gorm.DB, quizId string, quiz *models.Quiz) error {
	record := dbInstance.Where("ID = ?", quizId).Preload("Questions.Options").First(&quiz)
	if record.Error != nil {
		return errors.New(record.Error.Error())
	}
	return nil
}

func GetQuestionById(dbInstance *gorm.DB, questionId string, question *models.Question) error {
	record := dbInstance.Where("ID = ?", questionId).First(&question)
	if record.Error != nil {
		return errors.New(record.Error.Error())
	}
	return nil
}
