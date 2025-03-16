package usecase_todo

import (
	domain_todo "backend/internal/domain/todo"
	pkg_logger "backend/internal/pkg/logger"
)

// Todoユースケース(IF)
type ITodoUsecase interface {
	// 全てのTodoを取得
	GetAllTodos() ([]domain_todo.Todos, error)
}

// Todoユースケース(Impl)
type TodoUsecase struct {
	Logger         *pkg_logger.AppLogger
	todoRepository domain_todo.ITodoRepository
}

// Todoユースケースのインスタンス化
func NewTodoUsecase(l *pkg_logger.AppLogger, tr domain_todo.ITodoRepository) ITodoUsecase {
	return &TodoUsecase{
		Logger:         l,
		todoRepository: tr,
	}
}

// 全てのTodoを取得
func (u *TodoUsecase) GetAllTodos() ([]domain_todo.Todos, error) {
	u.Logger.InfoLog.Println("GetAllTodos called")

	// Todoリポジトリから全てのTodoを取得(repository層)
	todos, err := u.todoRepository.GetAllTodos()
	if err != nil {
		u.Logger.ErrorLog.Printf("Failed to get all todos: %v", err)
		return nil, err
	}

	u.Logger.InfoLog.Printf("Fetched %d todos", len(todos))
	return todos, nil
}
