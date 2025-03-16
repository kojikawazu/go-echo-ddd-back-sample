package test_user_usecase

import (
	domain_user "backend/internal/domain/user"

	"github.com/stretchr/testify/mock"
)

// モックのリポジトリ作成
type MockUserUsecase struct {
	mock.Mock
}

// GetAllUsersのモック
func (m *MockUserUsecase) GetAllUsers() ([]domain_user.Users, error) {
	args := m.Called()

	// `nil` チェックを追加
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]domain_user.Users), args.Error(1)
}
