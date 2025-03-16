package test_auth_usecase

import (
	"github.com/stretchr/testify/mock"
)

// モックのリポジトリ作成
type MockAuthUsecase struct {
	mock.Mock
}

// Loginのモック
func (m *MockAuthUsecase) Login(email string, password string) (string, error) {
	args := m.Called(email, password)

	// `nil` チェックを追加
	if args.Get(0) == nil {
		return "", args.Error(1)
	}

	return args.Get(0).(string), args.Error(1)
}
