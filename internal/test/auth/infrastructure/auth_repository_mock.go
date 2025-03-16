package test_auth_repository

import (
	"github.com/stretchr/testify/mock"
)

// モックのリポジトリ作成
type MockAuthRepository struct {
	mock.Mock
}

// Loginのモック
func (m *MockAuthRepository) Login(email string, password string) (string, error) {
	args := m.Called(email, password)

	// `nil` チェックを追加
	if args.Get(0) == nil {
		return "", args.Error(1)
	}

	return args.Get(0).(string), args.Error(1)
}
