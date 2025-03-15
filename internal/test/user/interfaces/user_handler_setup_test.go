package test_user_handler

import (
	pkg_config "backend/config"
	interfaces_user "backend/internal/interfaces/user"
	pkg_logger "backend/internal/pkg/logger"
	test_user_usecase "backend/internal/test/user/usecase"
	"os"
	"testing"
)

// テストの変数(グローバル用)
var (
	logger      *pkg_logger.AppLogger
	handler     *interfaces_user.UserHandler
	mockUsecase *test_user_usecase.MockUserUsecase
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
	mockUsecase = new(test_user_usecase.MockUserUsecase)
	handler = interfaces_user.NewUserHandler(logger, mockUsecase)

	// テスト実行
	code := m.Run()

	// 終了コードを返す
	os.Exit(code)
}
