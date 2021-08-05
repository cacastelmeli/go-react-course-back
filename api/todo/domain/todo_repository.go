package domain_todo

type TodoRepository interface {
	Add(todo *Todo)
	GetAll() []*Todo
	Remove(id float64)
	Update(todo *Todo)
}
