package test_auth_handler

import (
	pkg_config "backend/config"
	interfaces_auth "backend/internal/interfaces/auth"
	pkg_logger "backend/internal/pkg/logger"
	test_auth_usecase "backend/internal/test/auth/usecase"
	"os"
	"testing"
)

// テストの変数(グローバル用)
var (
	logger      *pkg_logger.AppLogger
	handler     *interfaces_auth.AuthHandler
	mockUsecase *test_auth_usecase.MockAuthUsecase
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
	mockUsecase = new(test_auth_usecase.MockAuthUsecase)
	handler = interfaces_auth.NewAuthHandler(logger, mockUsecase)

	// テスト実行
	code := m.Run()

	// 終了コードを返す
	os.Exit(code)
}
