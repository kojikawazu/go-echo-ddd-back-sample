package test_todo_handler

import (
	domain_todo "backend/internal/domain/todo"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// DeleteTodoのテスト(正常系)
func TestDeleteTodo(t *testing.T) {
	// モックの挙動をリセット
	mockUsecase.ExpectedCalls = nil

	// テストデータ
	id := "1"

	// モックの挙動を設定 (時間の影響を受けないように比較)
	mockUsecase.On("DeleteTodo", id).Return(nil)

	// リクエストの作成
	req := httptest.NewRequest("DELETE", "/api/todo/"+id, nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)

	// パスパラメータを設定
	c.SetParamNames("id")
	c.SetParamValues(id)

	// ハンドラの実行
	handler.DeleteTodo(c)

	// JSONレスポンスのデコード
	var resTodo domain_todo.Todo
	err := json.Unmarshal(rec.Body.Bytes(), &resTodo)
	assert.NoError(t, err)

	// レスポンスの検証
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockUsecase.AssertExpectations(t)
}

// DeleteTodoのテスト(異常系 - idが空)
func TestDeleteTodoErrorIdEmpty(t *testing.T) {
	// モックの挙動をリセット
	mockUsecase.ExpectedCalls = nil

	// テストデータ
	id := ""

	// モックの挙動を設定 (エラーを返す)
	mockUsecase.On("DeleteTodo", id).Return(errors.New("id is empty"))

	// リクエストの作成
	req := httptest.NewRequest("DELETE", "/api/todo/"+id, nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)

	// パスパラメータを設定
	c.SetParamNames("id")
	c.SetParamValues(id)

	// ハンドラの実行
	handler.DeleteTodo(c)

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

// DeleteTodoのテスト(異常系 - usecase異常)
func TestDeleteTodoErrorUsecase(t *testing.T) {
	// モックの挙動をリセット
	mockUsecase.ExpectedCalls = nil

	// テストデータ
	id := "1"

	// モックの挙動を設定 (エラーを返す)
	mockUsecase.On("DeleteTodo", id).Return(errors.New("error"))

	// リクエストの作成
	req := httptest.NewRequest("DELETE", "/api/todo/"+id, nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)

	// パスパラメータを設定
	c.SetParamNames("id")
	c.SetParamValues(id)

	// ハンドラの実行
	handler.DeleteTodo(c)

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
