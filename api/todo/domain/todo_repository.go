package domain_todo

type TodoRepository interface {
	Add(todo *Todo)
	GetAll() []*Todo
	Remove(id TodoId)
	Update(todo *Todo)
	Find(id TodoId) *Todo
}
