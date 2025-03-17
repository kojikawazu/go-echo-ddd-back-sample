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
func (m *MockTodoUsecase) GetAllTodos() ([]domain_todo.Todo, error) {
	args := m.Called()

	// `nil` チェックを追加
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]domain_todo.Todo), args.Error(1)
}

// GetTodoByIdのモック
func (m *MockTodoUsecase) GetTodoById(id string) (domain_todo.Todo, error) {
	args := m.Called(id)

	// `nil` チェックを追加
	if args.Get(0) == nil {
		return domain_todo.Todo{}, args.Error(1)
	}

	return args.Get(0).(domain_todo.Todo), args.Error(1)
}

// GetTodoByUserIdのモック
func (m *MockTodoUsecase) GetTodoByUserId(userId string) ([]domain_todo.Todo, error) {
	args := m.Called(userId)

	// `nil` チェックを追加
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]domain_todo.Todo), args.Error(1)
}

// CreateTodoのモック
func (m *MockTodoUsecase) CreateTodo(todo domain_todo.Todo) (domain_todo.Todo, error) {
	args := m.Called(todo)

	// `nil` チェックを追加
	if args.Get(0) == nil {
		return domain_todo.Todo{}, args.Error(1)
	}

	return args.Get(0).(domain_todo.Todo), args.Error(1)
}

// UpdateTodoのモック
func (m *MockTodoUsecase) UpdateTodo(todo domain_todo.Todo) (domain_todo.Todo, error) {
	args := m.Called(todo)

	// `nil` チェックを追加
	if args.Get(0) == nil {
		return domain_todo.Todo{}, args.Error(1)
	}

	return args.Get(0).(domain_todo.Todo), args.Error(1)
}

// DeleteTodoのモック
func (m *MockTodoUsecase) DeleteTodo(id string) error {
	args := m.Called(id)

	// `nil` チェックを追加
	if args.Get(0) == nil {
		return args.Error(0)
	}

	return args.Error(0)
}
