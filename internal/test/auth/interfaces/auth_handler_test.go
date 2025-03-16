package test_auth_handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// Loginのテスト(正常系)
func TestLogin(t *testing.T) {
	// モックの挙動をリセット
	mockUsecase.ExpectedCalls = nil

	// テストデータ
	userId := "1234567890"
	email := "test@example.com"
	password := "password"

	// モックの挙動を設定
	mockUsecase.On("Login", email, password).Return(userId, nil)

	// リクエストボディの作成
	requestBody := `{"email": "test@example.com", "password": "password"}`
	request := httptest.NewRequest("POST", "/api/users", strings.NewReader(requestBody))
	request.Header.Set("Content-Type", "application/json")

	// Echo のコンテキスト作成
	e := echo.New()
	response := httptest.NewRecorder()
	ctx := e.NewContext(request, response)

	// ハンドラのメソッドを呼び出し
	handler.Login(ctx)

	// 検証
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
	assert.Contains(t, response.Body.String(), `"token":`)

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockUsecase.AssertExpectations(t)
}

// Loginのテスト(異常系 - メールアドレスが空)
func TestLoginErrorEmailEmpty(t *testing.T) {
	// モックの挙動をリセット
	mockUsecase.ExpectedCalls = nil

	// テストデータ
	email := ""
	password := "password"

	// モックの挙動を設定
	mockUsecase.On("Login", email, password).Return("", errors.New("error"))

	// リクエストボディの作成
	requestBody := `{"email": "", "password": "password"}`
	request := httptest.NewRequest("POST", "/api/users", strings.NewReader(requestBody))
	request.Header.Set("Content-Type", "application/json")

	// Echo のコンテキスト作成
	e := echo.New()
	response := httptest.NewRecorder()
	ctx := e.NewContext(request, response)

	// ハンドラのメソッドを呼び出し
	handler.Login(ctx)

	// 検証
	assert.Equal(t, http.StatusInternalServerError, response.Code)
	assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
	assert.Contains(t, response.Body.String(), `"message":`)

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockUsecase.AssertExpectations(t)
}

// Loginのテスト(異常系 - パスワードが空)
func TestLoginErrorPasswordEmpty(t *testing.T) {
	// モックの挙動をリセット
	mockUsecase.ExpectedCalls = nil

	// テストデータ
	email := "test@example.com"
	password := ""

	// モックの挙動を設定
	mockUsecase.On("Login", email, password).Return("", errors.New("error"))

	// リクエストボディの作成
	requestBody := `{"email": "test@example.com", "password": ""}`
	request := httptest.NewRequest("POST", "/api/users", strings.NewReader(requestBody))
	request.Header.Set("Content-Type", "application/json")

	// Echo のコンテキスト作成
	e := echo.New()
	response := httptest.NewRecorder()
	ctx := e.NewContext(request, response)

	// ハンドラのメソッドを呼び出し
	handler.Login(ctx)

	// 検証
	assert.Equal(t, http.StatusInternalServerError, response.Code)
	assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
	assert.Contains(t, response.Body.String(), `"message":`)

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockUsecase.AssertExpectations(t)
}

// Loginのテスト(異常系 - 失敗)
func TestLoginError(t *testing.T) {
	// モックの挙動をリセット
	mockUsecase.ExpectedCalls = nil

	// テストデータ
	email := "test@example.com"
	password := "password"

	// モックの挙動を設定
	mockUsecase.On("Login", email, password).Return("", errors.New("error"))

	// リクエストボディの作成
	requestBody := `{"email": "test@example.com", "password": "password"}`
	request := httptest.NewRequest("POST", "/api/users", strings.NewReader(requestBody))
	request.Header.Set("Content-Type", "application/json")

	// Echo のコンテキスト作成
	e := echo.New()
	response := httptest.NewRecorder()
	ctx := e.NewContext(request, response)

	// ハンドラのメソッドを呼び出し
	handler.Login(ctx)

	// 検証
	assert.Equal(t, http.StatusInternalServerError, response.Code)
	assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
	assert.Contains(t, response.Body.String(), `"message":`)

	// モックのメソッドが期待通りに呼ばれたことを確認
	mockUsecase.AssertExpectations(t)
}
