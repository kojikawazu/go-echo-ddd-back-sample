package test_auth_usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Loginのテスト
func TestLogin(t *testing.T) {
	// モックの挙動をリセット
	mockRepo.ExpectedCalls = nil
	// テストデータ
	auth := "1234567890"
	email := "test@example.com"
	password := "password"

	// モックの挙動を設定
	mockRepo.On("Login", email, password).Return(auth, nil)

	// ユースケースのメソッドを呼び出し
	result, err := useCase.Login(email, password)

	// 検証
	assert.NoError(t, err)
	assert.Equal(t, auth, result)

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockRepo.AssertExpectations(t)
}

// Loginのテスト(空のデータ)
func TestLoginEmpty(t *testing.T) {
	// モックの挙動をリセット
	mockRepo.ExpectedCalls = nil
	// テストデータ
	auth := ""
	email := "test@example.com"
	password := "password"

	// モックの挙動を設定
	mockRepo.On("Login", email, password).Return(auth, nil)

	// ユースケースのメソッドを呼び出し
	result, err := useCase.Login(email, password)

	// 検証
	assert.NoError(t, err)
	assert.Equal(t, auth, result)

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockRepo.AssertExpectations(t)
}

// Loginのテスト(異常系 - メールアドレスが空)
func TestLoginErrorEmailEmpty(t *testing.T) {
	// モックの挙動をリセット
	mockRepo.ExpectedCalls = nil
	// テストデータ
	email := ""
	password := "password"

	// モックの挙動を設定
	mockRepo.On("Login", email, password).Return("", errors.New("error"))

	// ユースケースのメソッドを呼び出し
	result, err := useCase.Login(email, password)

	// 検証
	assert.Error(t, err)
	assert.Empty(t, result)

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockRepo.AssertNotCalled(t, "Login")
}

// Loginのテスト(異常系 - メールアドレスが形式が不正)
func TestLoginErrorEmailInvalid(t *testing.T) {
	// モックの挙動をリセット
	mockRepo.ExpectedCalls = nil
	// テストデータ
	email := "test"
	password := "password"

	// モックの挙動を設定
	mockRepo.On("Login", email, password).Return("", errors.New("error"))

	// ユースケースのメソッドを呼び出し
	result, err := useCase.Login(email, password)

	// 検証
	assert.Error(t, err)
	assert.Empty(t, result)

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockRepo.AssertNotCalled(t, "Login")
}

// Loginのテスト(異常系 - パスワードが空)
func TestLoginErrorPasswordEmpty(t *testing.T) {
	// モックの挙動をリセット
	mockRepo.ExpectedCalls = nil
	// テストデータ
	email := "test@example.com"
	password := ""

	// モックの挙動を設定
	mockRepo.On("Login", email, password).Return("", errors.New("error"))

	// ユースケースのメソッドを呼び出し
	result, err := useCase.Login(email, password)

	// 検証
	assert.Error(t, err)
	assert.Empty(t, result)

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockRepo.AssertNotCalled(t, "Login")
}

// Loginのテスト(異常系 - 失敗)
func TestLoginError(t *testing.T) {
	// モックの挙動をリセット
	mockRepo.ExpectedCalls = nil
	// テストデータ
	email := "test@example.com"
	password := "password"

	// モックの挙動を設定
	mockRepo.On("Login", email, password).Return("", errors.New("error"))

	// ユースケースのメソッドを呼び出し
	result, err := useCase.Login(email, password)

	// 検証
	assert.Error(t, err)
	assert.Empty(t, result)

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockRepo.AssertExpectations(t)
}
