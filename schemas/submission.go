package schemas

type Submission struct {
	QuizID     uint     `json:"quizId"`
	AnsweredBy string   `json:"answeredBy"`
	Answers    []Answer `json:"answers"`
}
