package test_todo_handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	domain_todo "backend/internal/domain/todo"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// GetAllUsersのテスト(正常系)
func TestGetAllUsers(t *testing.T) {
	// モックの挙動をリセット
	mockUsecase.ExpectedCalls = nil

	// テストデータ
	fixedTime := "2021-01-01T00:00:00Z"
	todos := []domain_todo.Todos{
		{ID: "1", Description: "alice@example.com", Completed: false, UserId: "1", CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
		{ID: "2", Description: "bob@example.com", Completed: false, UserId: "2", CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
	}

	// モックの挙動を設定
	mockUsecase.On("GetAllTodos").Return(todos, nil)

	// ハンドラのメソッドを呼び出し
	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/api/todo", nil)
	handler.GetAllTodos(echo.New().NewContext(request, response))

	// 検証
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
	assert.JSONEq(t, `[
		{"id": "1", "description": "alice@example.com", "completed": false, "user_id": "1", "created_at": "`+fixedTime+`", "updated_at": "`+fixedTime+`"},
		{"id": "2", "description": "bob@example.com", "completed": false, "user_id": "2", "created_at": "`+fixedTime+`", "updated_at": "`+fixedTime+`"}
	]`, response.Body.String())
	// モックのメソッドが期待通りに呼ばれたことを確認
	mockUsecase.AssertExpectations(t)
}

// GetAllTodosのテスト(空のデータ)
func TestGetAllTodosEmpty(t *testing.T) {
	// モックの挙動をリセット
	mockUsecase.ExpectedCalls = nil

	// テストデータ
	todos := []domain_todo.Todos{}

	// モックの挙動を設定
	mockUsecase.On("GetAllTodos").Return(todos, nil)

	// ハンドラのメソッドを呼び出し
	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/api/todo", nil)
	handler.GetAllTodos(echo.New().NewContext(request, response))

	// 検証
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
	assert.JSONEq(t, `[]`, response.Body.String())
	// モックのメソッドが期待通りに呼ばれたことを確認
	mockUsecase.AssertExpectations(t)
}

// GetAllTodosのテスト(異常系)
func TestGetAllTodosError(t *testing.T) {
	// モックの挙動をリセット
	mockUsecase.ExpectedCalls = nil

	// モックの挙動を設定
	mockUsecase.On("GetAllTodos").Return(nil, errors.New("error"))

	// ハンドラのメソッドを呼び出し
	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/api/todo", nil)
	handler.GetAllTodos(echo.New().NewContext(request, response))

	// 検証
	assert.Equal(t, http.StatusInternalServerError, response.Code)
	assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"message": "error"}`, response.Body.String())
	// モックのメソッドが期待通りに呼ばれたことを確認
	mockUsecase.AssertExpectations(t)
}
