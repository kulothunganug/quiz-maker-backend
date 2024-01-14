package schemas

type Quiz struct {
	Title     string     `json:"title" binding:"required"`
	CreatedBy string     `json:"createdBy" binding:"required"`
	Questions []Question `json:"questions" binding:"required"`
}
