package interfaces_todo

import (
	pkg_logger "backend/internal/pkg/logger"
	usecase_todo "backend/internal/usecase/todo"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Todoハンドラ(Impl)
type TodoHandler struct {
	Logger      *pkg_logger.AppLogger
	todoUsecase usecase_todo.ITodoUsecase
}

// Todoハンドラのインスタンス化
func NewTodoHandler(l *pkg_logger.AppLogger, tu usecase_todo.ITodoUsecase) *TodoHandler {
	return &TodoHandler{
		Logger:      l,
		todoUsecase: tu,
	}
}

// 全てのTodoを取得
func (h *TodoHandler) GetAllTodos(c echo.Context) error {
	h.Logger.InfoLog.Println("GetAllTodos called")

	// Todoユースケースから全てのTodoを取得
	todos, err := h.todoUsecase.GetAllTodos()
	// エラーがあればエラーレスポンスを返す
	if err != nil {
		h.Logger.ErrorLog.Printf("Failed to get all todos: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	// TodoのリストをJSON形式で返す
	h.Logger.InfoLog.Printf("Todos: %v", len(todos))
	return c.JSON(http.StatusOK, todos)
}
