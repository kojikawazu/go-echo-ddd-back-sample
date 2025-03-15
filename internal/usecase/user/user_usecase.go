package usecase_user

import (
	domain_user "backend/internal/domain/user"
	pkg_logger "backend/internal/pkg/logger"
)

// ユーザーユースケース(IF)
type IUserUsecase interface {
	GetAllUsers() ([]domain_user.User, error)
}

// ユーザーユースケース(Impl)
type UserUsecase struct {
	Logger         *pkg_logger.AppLogger
	userRepository domain_user.IUserRepository
}

// ユーザーユースケースのインスタンス化
func NewUserUsecase(l *pkg_logger.AppLogger, u domain_user.IUserRepository) IUserUsecase {
	return &UserUsecase{
		Logger:         l,
		userRepository: u,
	}
}

// 全てのユーザーを取得
func (u *UserUsecase) GetAllUsers() ([]domain_user.User, error) {
	u.Logger.InfoLog.Println("GetAllUsers called")

	// ユーザーリポジトリから全てのユーザーを取得(repository層)
	users, err := u.userRepository.GetAllUsers()
	if err != nil {
		u.Logger.ErrorLog.Printf("Failed to get all users: %v", err)
		return nil, err
	}

	u.Logger.InfoLog.Printf("Fetched %d users", len(users))
	return users, nil
}
