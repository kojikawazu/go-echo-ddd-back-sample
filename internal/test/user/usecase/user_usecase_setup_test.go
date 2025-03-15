package test_user_usecase

import (
	pkg_config "backend/config"
	pkg_logger "backend/internal/pkg/logger"
	test_user_repository "backend/internal/test/user/infrastructure"
	usecase_user "backend/internal/usecase/user"
	"os"
	"testing"
)

// テストの変数(グローバル用)
var (
	logger   *pkg_logger.AppLogger
	useCase  usecase_user.IUserUsecase
	mockRepo *test_user_repository.MockUserRepository
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
	mockRepo = new(test_user_repository.MockUserRepository)
	useCase = usecase_user.NewUserUsecase(logger, mockRepo)

	// テスト実行
	code := m.Run()

	// 終了コードを返す
	os.Exit(code)
}
