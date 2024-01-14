package schemas

type Submission struct {
	QuizID     uint     `json:"quizId" binding:"required"`
	AnsweredBy string   `json:"answeredBy" binding:"required"`
	Answers    []Answer `json:"answers" binding:"required"`
}
