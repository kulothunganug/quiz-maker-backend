package schemas

type Quiz struct {
	Title     string     `json:"title"`
	CreatedBy string     `json:"createdBy"`
	Questions []Question `json:"questions"`
}
