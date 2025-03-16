package interfaces_user

import (
	pkg_logger "backend/internal/pkg/logger"
	usecase_user "backend/internal/usecase/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ユーザーハンドラ
type UserHandler struct {
	Logger      *pkg_logger.AppLogger
	userUsecase usecase_user.IUserUsecase
}

// ユーザーハンドラのインスタンス化
func NewUserHandler(l *pkg_logger.AppLogger, u usecase_user.IUserUsecase) *UserHandler {
	return &UserHandler{
		Logger:      l,
		userUsecase: u,
	}
}

// 全てのユーザーを取得
func (h *UserHandler) GetAllUsers(c echo.Context) error {
	h.Logger.InfoLog.Println("GetAllUsers called")

	// 全てのユーザーを取得(usecase層)
	users, err := h.userUsecase.GetAllUsers()
	// エラーハンドリング
	if err != nil {
		h.Logger.ErrorLog.Printf("Failed to get all users: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	h.Logger.InfoLog.Printf("Fetched %d users", len(users))
	return c.JSON(http.StatusOK, users)
}
