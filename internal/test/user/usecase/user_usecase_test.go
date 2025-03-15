package test_user_usecase

import (
	"errors"
	"testing"

	domain_user "backend/internal/domain/user"

	"github.com/stretchr/testify/assert"
)

// GetAllUsersのテスト
func TestGetAllUsers(t *testing.T) {
	// モックの挙動をリセット
	mockRepo.ExpectedCalls = nil
	// テストデータ
	users := []domain_user.User{
		{ID: "1", Username: "Alice", Email: "alice@example.com"},
		{ID: "2", Username: "Bob", Email: "bob@example.com"},
	}

	// モックの挙動を設定
	mockRepo.On("GetAllUsers").Return(users, nil)

	// ユースケースのメソッドを呼び出し
	result, err := useCase.GetAllUsers()

	// 検証
	assert.NoError(t, err)
	assert.Equal(t, len(users), len(result))
	assert.Equal(t, users[0].Username, result[0].Username)
	assert.Equal(t, users[1].Email, result[1].Email)

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockRepo.AssertExpectations(t)
}

// GetAllUsersのテスト(空のデータ)
func TestGetAllUsersEmpty(t *testing.T) {
	// モックの挙動をリセット
	mockRepo.ExpectedCalls = nil
	// テストデータ
	users := []domain_user.User{}

	// モックの挙動を設定
	mockRepo.On("GetAllUsers").Return(users, nil)

	// ユースケースのメソッドを呼び出し
	result, err := useCase.GetAllUsers()

	// 検証
	assert.NoError(t, err)
	assert.Equal(t, len(users), len(result))

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockRepo.AssertExpectations(t)
}

// GetAllUsersのテスト(異常系)
func TestGetAllUsersError(t *testing.T) {
	// モックの挙動をリセット
	mockRepo.ExpectedCalls = nil

	// モックの挙動を設定
	mockRepo.On("GetAllUsers").Return(([]domain_user.User)(nil), errors.New("error"))

	// ユースケースのメソッドを呼び出し
	result, err := useCase.GetAllUsers()

	// 検証
	assert.Error(t, err)
	assert.Nil(t, result)

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockRepo.AssertExpectations(t)
}
