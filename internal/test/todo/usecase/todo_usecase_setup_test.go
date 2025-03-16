package test_todo_usecase

import (
	pkg_config "backend/config"
	pkg_logger "backend/internal/pkg/logger"
	test_todo_repository "backend/internal/test/todo/infrastructure"
	usecase_todo "backend/internal/usecase/todo"
	"os"
	"testing"
)

// テストの変数(グローバル用)
var (
	logger   *pkg_logger.AppLogger
	useCase  usecase_todo.ITodoUsecase
	mockRepo *test_todo_repository.MockTodoRepository
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
	mockRepo = new(test_todo_repository.MockTodoRepository)
	useCase = usecase_todo.NewTodoUsecase(logger, mockRepo)

	// テスト実行
	code := m.Run()

	// 終了コードを返す
	os.Exit(code)
}
