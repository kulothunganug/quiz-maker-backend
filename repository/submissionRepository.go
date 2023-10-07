package repository

import (
	"errors"
	"fmt"
	"quiz-maker-backend/database"
	"quiz-maker-backend/models"
	"quiz-maker-backend/schemas"
)

func AddSubmission(submissionReq *schemas.Submission) (*models.Submission, error) {
	var score uint

	for _, answer := range submissionReq.Answers {
		var question models.Question
		err := GetQuestionById(fmt.Sprint(answer.QuestionID), &question)
		if err != nil {
			return nil, err
		}

		if answer.OptionID == question.CorrectOptionID {
			score++
		}
	}

	var submission models.Submission
	submission.QuizID = submissionReq.QuizID
	submission.Score = score
	submission.AnsweredBy = submissionReq.AnsweredBy

	record := database.Instance.Create(&submission)
	if record.Error != nil {
		return nil, errors.New(record.Error.Error())
	}

	return &submission, nil
}

func GetSubmissions(quizId uint, submissions *[]models.Submission) error {
	record := database.Instance.Where("quiz_id = ?", quizId).Find(&submissions).Order("answered_at ASC")
	if record.Error != nil {
		return errors.New(record.Error.Error())
	}
	return nil
}
