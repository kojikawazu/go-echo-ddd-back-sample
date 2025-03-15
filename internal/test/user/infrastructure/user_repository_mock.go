package test_user_repository

import (
	domain_user "backend/internal/domain/user"

	"github.com/stretchr/testify/mock"
)

// モックのリポジトリ作成
type MockUserRepository struct {
	mock.Mock
}

// GetAllUsersのモック
func (m *MockUserRepository) GetAllUsers() ([]domain_user.User, error) {
	args := m.Called()

	// `nil` チェックを追加
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]domain_user.User), args.Error(1)
}
