package domain_todo

type Todo struct {
	Id   float64 `json:"id"`
	Text string  `json:"text"`
	Done bool    `json:"done"`
}
