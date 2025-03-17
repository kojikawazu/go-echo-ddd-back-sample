package test_todo_usecase

import (
	domain_todo "backend/internal/domain/todo"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// GetTodoByUserIdのテスト
func TestGetTodoByUserId(t *testing.T) {
	// モックの挙動をリセット
	mockRepo.ExpectedCalls = nil

	// テストデータ(複数)
	todos := []domain_todo.Todo{
		{
			ID:          "1",
			Description: "Todo 1",
			Completed:   false,
			UserId:      "1",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          "2",
			Description: "Todo 2",
			Completed:   false,
			UserId:      "1",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	// モックの挙動を設定
	mockRepo.On("GetTodoByUserId", "1").Return(todos, nil)

	// ユースケースのメソッドを呼び出し
	result, err := useCase.GetTodoByUserId("1")

	// 検証
	assert.NoError(t, err)
	assert.Len(t, result, len(todos))

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockRepo.AssertExpectations(t)
}

// GetTodoByUserIdのテスト(異常系 - 引数が空)
func TestGetTodoByUserIdEmpty(t *testing.T) {
	// モックの挙動をリセット
	mockRepo.ExpectedCalls = nil

	// モックの挙動を設定
	mockRepo.On("GetTodoByUserId", "").Return(nil, errors.New("error"))

	// ユースケースのメソッドを呼び出し
	result, err := useCase.GetTodoByUserId("")

	// 検証
	assert.Error(t, err)
	assert.Len(t, result, 0)

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockRepo.AssertNotCalled(t, "GetTodoByUserId", "")
}

// GetTodoByUserIdのテスト(異常系)
func TestGetTodoByUserIdError(t *testing.T) {
	// モックの挙動をリセット
	mockRepo.ExpectedCalls = nil

	// モックの挙動を設定
	mockRepo.On("GetTodoByUserId", "1").Return(nil, errors.New("error"))

	// ユースケースのメソッドを呼び出し
	result, err := useCase.GetTodoByUserId("1")

	// 検証
	assert.Error(t, err)
	assert.Len(t, result, 0)

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockRepo.AssertExpectations(t)
}
