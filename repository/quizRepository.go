package repository

import (
	"errors"
	"math/rand"
	"quiz-maker-backend/database"
	"quiz-maker-backend/models"
	"quiz-maker-backend/schemas"
)

// should check for if the id already exists or not
// if exists generate a new id
func RandomUint() uint {
	return uint(rand.Uint32())
}

func CreateQuiz(quizReq schemas.Quiz) (uint, error) {
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

	record := database.Instance.Create(&quiz)
	if record.Error != nil {
		return 0, errors.New(record.Error.Error())
	}

	return quiz.ID, nil
}

func GetQuiz(quizId string, quiz *models.Quiz) error {
	record := database.Instance.Where("ID = ?", quizId).Preload("Questions.Options").First(&quiz)
	if record.Error != nil {
		return errors.New(record.Error.Error())
	}
	return nil
}

func GetQuestionById(questionId string, question *models.Question) error {
	record := database.Instance.Where("ID = ?", questionId).First(&question)
	if record.Error != nil {
		return errors.New(record.Error.Error())
	}
	return nil
}
