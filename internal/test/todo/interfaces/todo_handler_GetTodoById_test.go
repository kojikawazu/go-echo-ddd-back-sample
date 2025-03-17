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

// GetTodoByIdのテスト(正常系)
func TestGetTodoById(t *testing.T) {
	// モックの挙動をリセット
	mockUsecase.ExpectedCalls = nil

	// テストデータ
	id := "1"
	todo := domain_todo.Todo{
		ID:          "1",
		Description: "Todo 1",
		Completed:   false,
		UserId:      "1",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// モックの挙動を設定
	mockUsecase.On("GetTodoById", id).Return(todo, nil)

	// ハンドラのメソッドを呼び出し
	e := echo.New()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/todo/"+id, nil)
	c := e.NewContext(req, res)

	// パスパラメータを設定
	c.SetParamNames("id")
	c.SetParamValues(id)

	// ハンドラのメソッドを呼び出し
	handler.GetTodoById(c)

	// JSONレスポンスのデコード
	var resTodo domain_todo.Todo
	err := json.Unmarshal(res.Body.Bytes(), &resTodo)
	if err != nil {
		t.FailNow()
	}

	// JSONのデコード後に比較
	assert.Equal(t, todo.ID, resTodo.ID)
	assert.Equal(t, todo.Description, resTodo.Description)
	assert.Equal(t, todo.Completed, resTodo.Completed)
	assert.Equal(t, todo.UserId, resTodo.UserId)
	assert.WithinDuration(t, todo.CreatedAt, resTodo.CreatedAt, time.Second)
	assert.WithinDuration(t, todo.UpdatedAt, resTodo.UpdatedAt, time.Second)

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockUsecase.AssertExpectations(t)
}

// GetTodoByIdのテスト(異常系 - idが空)
func TestGetTodoByIdErrorIdEmpty(t *testing.T) {
	// モックの挙動をリセット
	mockUsecase.ExpectedCalls = nil

	// テストデータ
	id := ""

	// モックの挙動を設定
	mockUsecase.On("GetTodoById", id).Return(domain_todo.Todo{}, errors.New("id is empty"))

	// ハンドラのメソッドを呼び出し
	e := echo.New()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/todo/", nil)
	c := e.NewContext(req, res)

	// ハンドラのメソッドを呼び出し
	handler.GetTodoById(c)

	// JSONレスポンスのデコード
	var resTodo domain_todo.Todo
	err := json.Unmarshal(res.Body.Bytes(), &resTodo)
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

// GetTodoByIdのテスト(異常系 - Usecase異常)
func TestGetTodoByIdError(t *testing.T) {
	// モックの挙動をリセット
	mockUsecase.ExpectedCalls = nil

	// テストデータ
	id := "1"

	// モックの挙動を設定
	mockUsecase.On("GetTodoById", id).Return(domain_todo.Todo{}, errors.New("error"))

	// ハンドラのメソッドを呼び出し
	e := echo.New()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/todo/"+id, nil)
	c := e.NewContext(req, res)

	// パスパラメータを設定
	c.SetParamNames("id")
	c.SetParamValues(id)

	// ハンドラのメソッドを呼び出し
	handler.GetTodoById(c)

	// JSONレスポンスのデコード
	var resTodo domain_todo.Todo
	err := json.Unmarshal(res.Body.Bytes(), &resTodo)
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
