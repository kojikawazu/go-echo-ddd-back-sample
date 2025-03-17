package test_todo_usecase

import (
	domain_todo "backend/internal/domain/todo"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// CreateTodoのテスト
func TestCreateTodo(t *testing.T) {
	// モックの挙動をリセット
	mockRepo.ExpectedCalls = nil

	// テストデータ
	todo := domain_todo.Todo{
		ID:          "1",
		Description: "Todo 1",
		Completed:   false,
		UserId:      "1",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// モックの挙動を設定
	mockRepo.On("CreateTodo", todo).Return(todo, nil)

	// ユースケースのメソッドを呼び出し
	result, err := useCase.CreateTodo(todo)

	// 検証
	assert.NoError(t, err)
	assert.Equal(t, todo, result)

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockRepo.AssertExpectations(t)
}

// CreateTodoのテスト(異常系 - Descriptionが空)
func TestCreateTodoEmpty(t *testing.T) {
	// モックの挙動をリセット
	mockRepo.ExpectedCalls = nil

	// テストデータ
	todo := domain_todo.Todo{
		ID:          "1",
		Description: "",
		Completed:   false,
		UserId:      "1",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	// モックの挙動を設定
	mockRepo.On("CreateTodo", todo).Return(domain_todo.Todo{}, errors.New("error"))

	// ユースケースのメソッドを呼び出し
	result, err := useCase.CreateTodo(todo)

	// 検証
	assert.Error(t, err)
	assert.Equal(t, domain_todo.Todo{}, result)

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockRepo.AssertNotCalled(t, "CreateTodo", todo)
}

// CreateTodoのテスト(異常系 - UserIdが空)
func TestCreateTodoUserIdEmpty(t *testing.T) {
	// モックの挙動をリセット
	mockRepo.ExpectedCalls = nil

	// テストデータ
	todo := domain_todo.Todo{
		ID:          "1",
		Description: "Todo 1",
		Completed:   false,
		UserId:      "",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	// モックの挙動を設定
	mockRepo.On("CreateTodo", todo).Return(domain_todo.Todo{}, errors.New("error"))

	// ユースケースのメソッドを呼び出し
	result, err := useCase.CreateTodo(todo)

	// 検証
	assert.Error(t, err)
	assert.Equal(t, domain_todo.Todo{}, result)

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockRepo.AssertNotCalled(t, "CreateTodo", todo)
}

// CreateTodoのテスト(異常系 - リポジトリでエラーが発生)
func TestCreateTodoError(t *testing.T) {
	// モックの挙動をリセット
	mockRepo.ExpectedCalls = nil

	// テストデータ
	todo := domain_todo.Todo{
		ID:          "1",
		Description: "Todo 1",
		Completed:   false,
		UserId:      "1",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	// モックの挙動を設定
	mockRepo.On("CreateTodo", todo).Return(domain_todo.Todo{}, errors.New("error"))

	// ユースケースのメソッドを呼び出し
	result, err := useCase.CreateTodo(todo)

	// 検証
	assert.Error(t, err)
	assert.Equal(t, domain_todo.Todo{}, result)

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockRepo.AssertExpectations(t)
}
