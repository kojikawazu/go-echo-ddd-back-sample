package test_todo_usecase

import (
	domain_todo "backend/internal/domain/todo"

	"github.com/stretchr/testify/mock"
)

// モックのリポジトリ作成
type MockTodoUsecase struct {
	mock.Mock
}

// GetAllTodosのモック
func (m *MockTodoUsecase) GetAllTodos() ([]domain_todo.Todos, error) {
	args := m.Called()

	// `nil` チェックを追加
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]domain_todo.Todos), args.Error(1)
}
