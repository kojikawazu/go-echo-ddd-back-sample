package test_todo_handler

import (
	domain_todo "backend/internal/domain/todo"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// GetTodoByUserIdのテスト(正常系)
func TestGetTodoByUserId(t *testing.T) {
	// モックの挙動をリセット
	mockUsecase.ExpectedCalls = nil

	// テストデータ
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
	mockUsecase.On("GetTodoByUserId", "1").Return(todos, nil)

	// ハンドラのメソッドを呼び出し
	e := echo.New()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/todo/user", nil)
	c := e.NewContext(req, res)

	// コンテキストの設定
	c.Set("userId", "1")

	// ハンドラのメソッドを呼び出し
	handler.GetTodoByUserId(c)

	// JSONレスポンスのデコード
	var resTodos []domain_todo.Todo
	err := json.Unmarshal(res.Body.Bytes(), &resTodos)
	if err != nil {
		t.FailNow()
	}

	// JSONのデコード後に比較
	assert.Equal(t, todos[0].ID, resTodos[0].ID)
	assert.Equal(t, todos[0].Description, resTodos[0].Description)
	assert.Equal(t, todos[0].Completed, resTodos[0].Completed)
	assert.Equal(t, todos[0].UserId, resTodos[0].UserId)
	assert.WithinDuration(t, todos[0].CreatedAt, resTodos[0].CreatedAt, time.Second)
	assert.WithinDuration(t, todos[0].UpdatedAt, resTodos[0].UpdatedAt, time.Second)
	assert.Equal(t, todos[1].ID, resTodos[1].ID)
	assert.Equal(t, todos[1].Description, resTodos[1].Description)
	assert.Equal(t, todos[1].Completed, resTodos[1].Completed)
	assert.Equal(t, todos[1].UserId, resTodos[1].UserId)
	assert.WithinDuration(t, todos[1].CreatedAt, resTodos[1].CreatedAt, time.Second)
	assert.WithinDuration(t, todos[1].UpdatedAt, resTodos[1].UpdatedAt, time.Second)

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockUsecase.AssertExpectations(t)
}

// GetTodoByUserIdのテスト(異常系 - userIdが空)
func TestGetTodoByUserIdErrorUserIdEmpty(t *testing.T) {
	// モックの挙動をリセット
	mockUsecase.ExpectedCalls = nil

	// モックの挙動を設定
	mockUsecase.On("GetTodoByUserId", "").Return(nil, errors.New("userId is empty"))

	// ハンドラのメソッドを呼び出し
	e := echo.New()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/todo/user", nil)
	c := e.NewContext(req, res)

	// コンテキストの設定
	c.Set("userId", "")

	// ハンドラのメソッドを呼び出し
	handler.GetTodoByUserId(c)

	// JSONレスポンスのデコード
	var resTodos []domain_todo.Todo
	err := json.Unmarshal(res.Body.Bytes(), &resTodos)
	if err != nil {
		t.FailNow()
	}

	// JSONのデコード後に比較
	assert.Equal(t, http.StatusNotFound, res.Code)
	assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
	assert.Contains(t, res.Body.String(), `"message":`)

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockUsecase.AssertExpectations(t)
}

// GetTodoByUserIdのテスト(異常系 - Usecase異常)
func TestGetTodoByUserIdError(t *testing.T) {
	// モックの挙動をリセット
	mockUsecase.ExpectedCalls = nil

	// テストデータ
	userId := "1"

	// モックの挙動を設定
	mockUsecase.On("GetTodoByUserId", userId).Return(nil, errors.New("error"))

	// ハンドラのメソッドを呼び出し
	e := echo.New()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/todo/user", nil)
	c := e.NewContext(req, res)

	// コンテキストの設定
	c.Set("userId", userId)

	// ハンドラのメソッドを呼び出し
	handler.GetTodoByUserId(c)

	// JSONレスポンスのデコード
	var resTodos []domain_todo.Todo
	err := json.Unmarshal(res.Body.Bytes(), &resTodos)
	if err != nil {
		t.FailNow()
	}

	// JSONのデコード後に比較
	assert.Equal(t, http.StatusInternalServerError, res.Code)
	assert.Equal(t, "application/json", res.Header().Get("Content-Type"))
	assert.Contains(t, res.Body.String(), `"message":`)

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockUsecase.AssertExpectations(t)
}
