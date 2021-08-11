package domain_todo

// Value objects
type TodoId float64
type TodoText string
type TodoDone bool

type Todo struct {
	Id   TodoId   `json:"id"`
	Text TodoText `json:"text"`
	Done TodoDone `json:"done"`
}
