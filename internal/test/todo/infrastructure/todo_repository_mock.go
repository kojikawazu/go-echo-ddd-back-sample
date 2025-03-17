package test_todo_repository

import (
	domain_todo "backend/internal/domain/todo"

	"github.com/stretchr/testify/mock"
)

// モックのリポジトリ作成
type MockTodoRepository struct {
	mock.Mock
}

// GetAllTodosのモック
func (m *MockTodoRepository) GetAllTodos() ([]domain_todo.Todo, error) {
	args := m.Called()

	// `nil` チェックを追加
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]domain_todo.Todo), args.Error(1)
}

// GetTodoByIdのモック
func (m *MockTodoRepository) GetTodoById(id string) (domain_todo.Todo, error) {
	args := m.Called(id)

	// `nil` チェックを追加
	if args.Get(0) == nil {
		return domain_todo.Todo{}, args.Error(1)
	}

	return args.Get(0).(domain_todo.Todo), args.Error(1)
}

// CreateTodoのモック
func (m *MockTodoRepository) CreateTodo(todo domain_todo.Todo) (domain_todo.Todo, error) {
	args := m.Called(todo)

	// `nil` チェックを追加
	if args.Get(0) == nil {
		return domain_todo.Todo{}, args.Error(1)
	}

	return args.Get(0).(domain_todo.Todo), args.Error(1)
}

// UpdateTodoのモック
func (m *MockTodoRepository) UpdateTodo(todo domain_todo.Todo) (domain_todo.Todo, error) {
	args := m.Called(todo)

	// `nil` チェックを追加
	if args.Get(0) == nil {
		return domain_todo.Todo{}, args.Error(1)
	}

	return args.Get(0).(domain_todo.Todo), args.Error(1)
}

// DeleteTodoのモック
func (m *MockTodoRepository) DeleteTodo(id string) error {
	args := m.Called(id)

	// `nil` チェックを追加
	if args.Get(0) == nil {
		return args.Error(0)
	}

	return args.Error(0)
}
