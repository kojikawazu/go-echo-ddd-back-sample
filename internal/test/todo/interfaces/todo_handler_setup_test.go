package test_todo_handler

import (
	pkg_config "backend/config"
	interfaces_todo "backend/internal/interfaces/todo"
	pkg_logger "backend/internal/pkg/logger"
	test_todo_usecase "backend/internal/test/todo/usecase"
	"os"
	"testing"
)

// テストの変数(グローバル用)
var (
	logger      *pkg_logger.AppLogger
	handler     *interfaces_todo.TodoHandler
	mockUsecase *test_todo_usecase.MockTodoUsecase
)

// テストのメイン関数
func TestMain(m *testing.M) {
	// 設定
	appConfig := pkg_config.NewAppConfig()
	appConfig.SetUpEnv()

	// ログ
	logger = pkg_logger.NewAppLogger()
	logger.SetUpLogger()

	// モック
	mockUsecase = new(test_todo_usecase.MockTodoUsecase)
	handler = interfaces_todo.NewTodoHandler(logger, mockUsecase)

	// テスト実行
	code := m.Run()

	// 終了コードを返す
	os.Exit(code)
}
