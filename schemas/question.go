package schemas

type Question struct {
	Options          []Option `json:"options"`
	Text             string   `json:"text"`
	CorrectOptionIdx uint     `json:"correctOptionIdx"`
}
