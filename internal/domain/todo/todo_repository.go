package domain_todo

// Todoリポジトリ(IF)
type ITodoRepository interface {
	// 全てのTodoを取得
	GetAllTodos() ([]Todos, error)
}
