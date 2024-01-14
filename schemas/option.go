package schemas

type Option struct {
	Text string `json:"text" binding:"required"`
}
