package test_user_handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	domain_user "backend/internal/domain/user"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// GetAllUsersのテスト(正常系)
func TestGetAllUsers(t *testing.T) {
	// モックの挙動をリセット
	mockUsecase.ExpectedCalls = nil

	// テストデータ
	fixedTime := "2021-01-01T00:00:00Z"
	users := []domain_user.User{
		{ID: "1", Username: "Alice", Email: "alice@example.com", Password: "", CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
		{ID: "2", Username: "Bob", Email: "bob@example.com", Password: "", CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
	}

	// モックの挙動を設定
	mockUsecase.On("GetAllUsers").Return(users, nil)

	// ハンドラのメソッドを呼び出し
	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/api/users", nil)
	handler.GetAllUsers(echo.New().NewContext(request, response))

	// 検証
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
	assert.JSONEq(t, `[
		{"id": "1", "username": "Alice", "email": "alice@example.com", "password": "", "created_at": "`+fixedTime+`", "updated_at": "`+fixedTime+`"},
		{"id": "2", "username": "Bob", "email": "bob@example.com", "password": "", "created_at": "`+fixedTime+`", "updated_at": "`+fixedTime+`"}
	]`, response.Body.String())
	// モックのメソッドが期待通りに呼ばれたことを確認
	mockUsecase.AssertExpectations(t)
}

// GetAllUsersのテスト(空のデータ)
func TestGetAllUsersEmpty(t *testing.T) {
	// モックの挙動をリセット
	mockUsecase.ExpectedCalls = nil

	// テストデータ
	users := []domain_user.User{}

	// モックの挙動を設定
	mockUsecase.On("GetAllUsers").Return(users, nil)

	// ハンドラのメソッドを呼び出し
	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/api/users", nil)
	handler.GetAllUsers(echo.New().NewContext(request, response))

	// 検証
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
	assert.JSONEq(t, `[]`, response.Body.String())
	// モックのメソッドが期待通りに呼ばれたことを確認
	mockUsecase.AssertExpectations(t)
}

// GetAllUsersのテスト(異常系)
func TestGetAllUsersError(t *testing.T) {
	// モックの挙動をリセット
	mockUsecase.ExpectedCalls = nil

	// モックの挙動を設定
	mockUsecase.On("GetAllUsers").Return(nil, errors.New("error"))

	// ハンドラのメソッドを呼び出し
	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/api/users", nil)
	handler.GetAllUsers(echo.New().NewContext(request, response))

	// 検証
	assert.Equal(t, http.StatusInternalServerError, response.Code)
	assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"message": "error"}`, response.Body.String())
	// モックのメソッドが期待通りに呼ばれたことを確認
	mockUsecase.AssertExpectations(t)
}
