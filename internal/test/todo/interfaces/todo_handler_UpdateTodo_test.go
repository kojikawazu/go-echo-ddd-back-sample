package test_todo_handler

import (
	domain_todo "backend/internal/domain/todo"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// UpdateTodoのテスト(正常系)
func TestUpdateTodo(t *testing.T) {
	// モックの挙動をリセット
	mockUsecase.ExpectedCalls = nil

	// テストデータ
	todo := domain_todo.Todo{
		ID:          "1",
		Description: "Todo 1",
		Completed:   false,
		UserId:      "1",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// モックの挙動を設定 (時間の影響を受けないように比較)
	mockUsecase.On("UpdateTodo", mock.MatchedBy(func(t domain_todo.Todo) bool {
		return t.ID == todo.ID && t.Description == todo.Description && t.UserId == todo.UserId
	})).Return(todo, nil)

	// リクエストの作成
	body, _ := json.Marshal(todo)
	req := httptest.NewRequest("PUT", "/api/todo/"+todo.ID, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)

	// パスパラメータを設定
	c.SetParamNames("id")
	c.SetParamValues(todo.ID)

	// ハンドラの実行
	handler.UpdateTodo(c)

	// JSONレスポンスのデコード
	var resTodo domain_todo.Todo
	err := json.Unmarshal(rec.Body.Bytes(), &resTodo)
	assert.NoError(t, err)

	// レスポンスの検証
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))
	assert.Equal(t, todo.ID, resTodo.ID)
	assert.Equal(t, todo.Description, resTodo.Description)
	assert.Equal(t, todo.Completed, resTodo.Completed)
	assert.Equal(t, todo.UserId, resTodo.UserId)
	assert.WithinDuration(t, todo.CreatedAt, resTodo.CreatedAt, time.Second)
	assert.WithinDuration(t, todo.UpdatedAt, resTodo.UpdatedAt, time.Second)

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockUsecase.AssertExpectations(t)
}

// UpdateTodoのテスト(異常系 - idが空)
func TestUpdateTodoErrorIdEmpty(t *testing.T) {
	// モックの挙動をリセット
	mockUsecase.ExpectedCalls = nil

	// テストデータ
	todo := domain_todo.Todo{
		ID:          "",
		Description: "Todo 1",
		Completed:   false,
		UserId:      "1",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// モックの挙動を設定 (エラーを返す)
	mockUsecase.On("UpdateTodo", mock.Anything).Return(domain_todo.Todo{}, errors.New("id is empty"))

	// リクエストの作成
	body, _ := json.Marshal(todo)
	req := httptest.NewRequest("PUT", "/api/todo/"+todo.ID, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)

	// パスパラメータを設定
	c.SetParamNames("id")
	c.SetParamValues(todo.ID)

	// ハンドラの実行
	handler.UpdateTodo(c)

	// レスポンスのデコード
	var resBody map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &resBody)
	assert.NoError(t, err)

	// レスポンスの検証
	assert.Equal(t, http.StatusNotFound, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))
	assert.Equal(t, "id is empty", resBody["message"])

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockUsecase.AssertExpectations(t)
}

// UpdateTodoのテスト(異常系 - descriptionが空)
func TestUpdateTodoErrorDescriptionEmpty(t *testing.T) {
	// モックの挙動をリセット
	mockUsecase.ExpectedCalls = nil

	// テストデータ
	todo := domain_todo.Todo{
		ID:          "1",
		Description: "",
		Completed:   false,
		UserId:      "1",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// モックの挙動を設定 (エラーを返す)
	mockUsecase.On("UpdateTodo", mock.Anything).Return(domain_todo.Todo{}, errors.New("description is empty"))

	// リクエストの作成
	body, _ := json.Marshal(todo)
	req := httptest.NewRequest("PUT", "/api/todo/"+todo.ID, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)

	// パスパラメータを設定
	c.SetParamNames("id")
	c.SetParamValues(todo.ID)

	// ハンドラの実行
	handler.UpdateTodo(c)

	// レスポンスのデコード
	var resBody map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &resBody)
	assert.NoError(t, err)

	// レスポンスの検証
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))
	assert.Equal(t, "description is empty", resBody["message"])

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockUsecase.AssertExpectations(t)
}

// UpdateTodoのテスト(異常系 - user_idが空)
func TestUpdateTodoErrorUserIdEmpty(t *testing.T) {
	// モックの挙動をリセット
	mockUsecase.ExpectedCalls = nil

	// テストデータ
	todo := domain_todo.Todo{
		ID:          "1",
		Description: "Todo 1",
		Completed:   false,
		UserId:      "",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// モックの挙動を設定 (エラーを返す)
	mockUsecase.On("UpdateTodo", mock.Anything).Return(domain_todo.Todo{}, errors.New("user_id is empty"))

	// リクエストの作成
	body, _ := json.Marshal(todo)
	req := httptest.NewRequest("PUT", "/api/todo/"+todo.ID, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)

	// パスパラメータを設定
	c.SetParamNames("id")
	c.SetParamValues(todo.ID)

	// ハンドラの実行
	handler.UpdateTodo(c)

	// レスポンスのデコード
	var resBody map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &resBody)
	assert.NoError(t, err)

	// レスポンスの検証
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))
	assert.Equal(t, "user_id is empty", resBody["message"])

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockUsecase.AssertExpectations(t)
}

// UpdateTodoのテスト(異常系 - usecase異常)
func TestUpdateTodoErrorUsecase(t *testing.T) {
	// モックの挙動をリセット
	mockUsecase.ExpectedCalls = nil

	// テストデータ
	todo := domain_todo.Todo{
		ID:          "1",
		Description: "Todo 1",
		Completed:   false,
		UserId:      "1",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// モックの挙動を設定 (エラーを返す)
	mockUsecase.On("UpdateTodo", mock.Anything).Return(domain_todo.Todo{}, errors.New("error"))

	// リクエストの作成
	body, _ := json.Marshal(todo)
	req := httptest.NewRequest("PUT", "/api/todo/"+todo.ID, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)

	// パスパラメータを設定
	c.SetParamNames("id")
	c.SetParamValues(todo.ID)

	// ハンドラの実行
	handler.UpdateTodo(c)

	// レスポンスのデコード
	var resBody map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &resBody)
	assert.NoError(t, err)

	// レスポンスの検証
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))
	assert.Equal(t, "error", resBody["message"])

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockUsecase.AssertExpectations(t)
}
