package schemas

type Question struct {
	Options          []Option `json:"options" binding:"required"`
	Text             string   `json:"text" binding:"required"`
	CorrectOptionIdx uint     `json:"correctOptionIdx" binding:"required"`
}
