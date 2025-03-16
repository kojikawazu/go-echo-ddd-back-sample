package test_todo_usecase

import (
	"errors"
	"testing"
	"time"

	domain_todo "backend/internal/domain/todo"

	"github.com/stretchr/testify/assert"
)

// GetAllTodosのテスト
func TestGetAllTodos(t *testing.T) {
	// モックの挙動をリセット
	mockRepo.ExpectedCalls = nil
	// テストデータ
	todos := []domain_todo.Todos{
		{ID: "1", Description: "Todo 1", Completed: false, UserId: "1", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: "2", Description: "Todo 2", Completed: false, UserId: "2", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	// モックの挙動を設定
	mockRepo.On("GetAllTodos").Return(todos, nil)

	// ユースケースのメソッドを呼び出し
	result, err := useCase.GetAllTodos()

	// 検証
	assert.NoError(t, err)
	assert.Equal(t, len(todos), len(result))
	assert.Equal(t, todos[0].Description, result[0].Description)
	assert.Equal(t, todos[1].Description, result[1].Description)

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockRepo.AssertExpectations(t)
}

// GetAllTodosのテスト(空のデータ)
func TestGetAllTodosEmpty(t *testing.T) {
	// モックの挙動をリセット
	mockRepo.ExpectedCalls = nil
	// テストデータ
	todos := []domain_todo.Todos{}

	// モックの挙動を設定
	mockRepo.On("GetAllTodos").Return(todos, nil)

	// ユースケースのメソッドを呼び出し
	result, err := useCase.GetAllTodos()

	// 検証
	assert.NoError(t, err)
	assert.Equal(t, len(todos), len(result))

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockRepo.AssertExpectations(t)
}

// GetAllTodosのテスト(異常系)
func TestGetAllTodosError(t *testing.T) {
	// モックの挙動をリセット
	mockRepo.ExpectedCalls = nil

	// モックの挙動を設定
	mockRepo.On("GetAllTodos").Return(([]domain_todo.Todos)(nil), errors.New("error"))

	// ユースケースのメソッドを呼び出し
	result, err := useCase.GetAllTodos()

	// 検証
	assert.Error(t, err)
	assert.Nil(t, result)

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockRepo.AssertExpectations(t)
}
