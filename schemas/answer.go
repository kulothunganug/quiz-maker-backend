package schemas

type Answer struct {
	QuestionID uint `json:"questionId" binding:"required"`
	OptionID   uint `json:"optionId" binding:"required"`
}
