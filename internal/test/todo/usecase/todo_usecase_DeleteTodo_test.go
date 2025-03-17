package test_todo_usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// DeleteTodoのテスト
func TestDeleteTodo(t *testing.T) {
	// モックの挙動をリセット
	mockRepo.ExpectedCalls = nil

	// テストデータ
	id := "1"

	// モックの挙動を設定
	mockRepo.On("DeleteTodo", id).Return(nil)

	// ユースケースのメソッドを呼び出し
	err := useCase.DeleteTodo(id)

	// 検証
	assert.NoError(t, err)

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockRepo.AssertExpectations(t)
}

// DeleteTodoのテスト(異常系 - idが空)
func TestDeleteTodoIdEmpty(t *testing.T) {
	// モックの挙動をリセット
	mockRepo.ExpectedCalls = nil

	// テストデータ
	id := ""

	// モックの挙動を設定
	mockRepo.On("DeleteTodo", id).Return(errors.New("error"))

	// ユースケースのメソッドを呼び出し
	err := useCase.DeleteTodo(id)

	// 検証
	assert.Error(t, err)

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockRepo.AssertNotCalled(t, "DeleteTodo", id)
}

// DeleteTodoのテスト(異常系 - リポジトリでエラーが発生)
func TestDeleteTodoError(t *testing.T) {
	// モックの挙動をリセット
	mockRepo.ExpectedCalls = nil

	// テストデータ
	id := "1"

	// モックの挙動を設定
	mockRepo.On("DeleteTodo", id).Return(errors.New("error"))

	// ユースケースのメソッドを呼び出し
	err := useCase.DeleteTodo(id)

	// 検証
	assert.Error(t, err)

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockRepo.AssertExpectations(t)
}
