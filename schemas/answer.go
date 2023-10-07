package schemas

type Answer struct {
	QuestionID uint `json:"questionId"`
	OptionID   uint `json:"optionId"`
}
