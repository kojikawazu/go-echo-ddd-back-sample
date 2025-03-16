package test_auth_usecase

import (
	pkg_config "backend/config"
	pkg_logger "backend/internal/pkg/logger"
	test_auth_repository "backend/internal/test/auth/infrastructure"
	usecase_auth "backend/internal/usecase/auth"
	"os"
	"testing"
)

// テストの変数(グローバル用)
var (
	logger   *pkg_logger.AppLogger
	useCase  usecase_auth.IAuthUsecase
	mockRepo *test_auth_repository.MockAuthRepository
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
	mockRepo = new(test_auth_repository.MockAuthRepository)
	useCase = usecase_auth.NewAuthUsecase(logger, mockRepo)

	// テスト実行
	code := m.Run()

	// 終了コードを返す
	os.Exit(code)
}
