package test_todo_usecase

import (
	domain_todo "backend/internal/domain/todo"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// GetTodoByIdのテスト
func TestGetTodoById(t *testing.T) {
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
	mockRepo.On("GetTodoById", "1").Return(todo, nil)

	// ユースケースのメソッドを呼び出し
	result, err := useCase.GetTodoById("1")

	// 検証
	assert.NoError(t, err)
	assert.Equal(t, todo, result)

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockRepo.AssertExpectations(t)
}

// GetTodoByIdのテスト(異常系 - 引数が空)
func TestGetTodoByIdEmpty(t *testing.T) {
	// モックの挙動をリセット
	mockRepo.ExpectedCalls = nil

	// モックの挙動を設定
	mockRepo.On("GetTodoById", "").Return(domain_todo.Todo{}, errors.New("error"))

	// ユースケースのメソッドを呼び出し
	result, err := useCase.GetTodoById("")

	// 検証
	assert.Error(t, err)
	assert.Equal(t, domain_todo.Todo{}, result)

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockRepo.AssertNotCalled(t, "GetTodoById", "")
}

// GetTodoByIdのテスト(異常系)
func TestGetTodoByIdError(t *testing.T) {
	// モックの挙動をリセット
	mockRepo.ExpectedCalls = nil

	// モックの挙動を設定
	mockRepo.On("GetTodoById", "1").Return(domain_todo.Todo{}, errors.New("error"))

	// ユースケースのメソッドを呼び出し
	result, err := useCase.GetTodoById("1")

	// 検証
	assert.Error(t, err)
	assert.Equal(t, domain_todo.Todo{}, result)

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockRepo.AssertExpectations(t)
}
