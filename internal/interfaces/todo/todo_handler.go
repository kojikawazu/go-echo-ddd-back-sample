package interfaces_todo

import (
	domain_todo "backend/internal/domain/todo"
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

// idを指定してTodoを取得
func (h *TodoHandler) GetTodoById(c echo.Context) error {
	h.Logger.InfoLog.Println("GetTodoById called")

	// パスパラメータからidを取得
	id := c.Param("id")

	// Todoユースケースからidを指定してTodoを取得
	todo, err := h.todoUsecase.GetTodoById(id)
	// エラーハンドリング
	if err != nil {
		switch err.Error() {
		case "id is empty":
			h.Logger.ErrorLog.Printf("Failed to get todo by id: %v", err)
			return c.JSON(http.StatusNotFound, map[string]string{
				"message": err.Error(),
			})
		default:
			h.Logger.ErrorLog.Printf("Failed to get todo by id: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
			})
		}
	}

	// TodoをJSON形式で返す
	h.Logger.InfoLog.Printf("Todo: %v", todo)
	return c.JSON(http.StatusOK, todo)
}

// 特定のユーザーのTodoを取得
func (h *TodoHandler) GetTodoByUserId(c echo.Context) error {
	h.Logger.InfoLog.Println("GetTodoByUserId called")

	// Contextからuser_idを取得
	userID := c.Get("userId")

	// Todoユースケースから特定のユーザーのTodoを取得
	todos, err := h.todoUsecase.GetTodoByUserId(userID.(string))
	// エラーハンドリング
	if err != nil {
		h.Logger.ErrorLog.Printf("Failed to get todo by user_id: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	// TodoのリストをJSON形式で返す
	h.Logger.InfoLog.Printf("Todos: %v", len(todos))
	return c.JSON(http.StatusOK, todos)
}

// 新しいTodoを作成
func (h *TodoHandler) CreateTodo(c echo.Context) error {
	h.Logger.InfoLog.Println("CreateTodo called")

	// リクエストボディからTodoを取得
	todo := domain_todo.Todo{}
	if err := c.Bind(&todo); err != nil {
		h.Logger.ErrorLog.Printf("Failed to bind todo: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	// Todoユースケースから新しいTodoを作成
	createdTodo, err := h.todoUsecase.CreateTodo(todo)
	// エラーがあればエラーレスポンスを返す
	if err != nil {
		switch err.Error() {
		case "description is empty":
			h.Logger.ErrorLog.Printf("Failed to create todo: %v", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
		case "user_id is empty":
			h.Logger.ErrorLog.Printf("Failed to create todo: %v", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
		default:
			h.Logger.ErrorLog.Printf("Failed to create todo: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
			})
		}
	}

	// 作成したTodoをJSON形式で返す
	h.Logger.InfoLog.Printf("Created todo: %v", createdTodo)
	return c.JSON(http.StatusCreated, createdTodo)
}

// Todoを更新
func (h *TodoHandler) UpdateTodo(c echo.Context) error {
	h.Logger.InfoLog.Println("UpdateTodo called")

	// パスパラメータからidを取得
	id := c.Param("id")

	// リクエストボディからTodoを取得
	todo := domain_todo.Todo{}
	if err := c.Bind(&todo); err != nil {
		h.Logger.ErrorLog.Printf("Failed to bind todo: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	todo.ID = id

	// TodoユースケースからTodoを更新
	updatedTodo, err := h.todoUsecase.UpdateTodo(todo)
	// エラーがあればエラーレスポンスを返す
	if err != nil {
		switch err.Error() {
		case "id is empty":
			h.Logger.ErrorLog.Printf("Failed to update todo: %v", err)
			return c.JSON(http.StatusNotFound, map[string]string{
				"message": err.Error(),
			})
		case "description is empty":
			h.Logger.ErrorLog.Printf("Failed to update todo: %v", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
		case "user_id is empty":
			h.Logger.ErrorLog.Printf("Failed to update todo: %v", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
		default:
			h.Logger.ErrorLog.Printf("Failed to update todo: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
			})
		}
	}

	// 更新したTodoをJSON形式で返す
	h.Logger.InfoLog.Printf("Updated todo: %v", updatedTodo)
	return c.JSON(http.StatusOK, updatedTodo)
}

// Todoを削除
func (h *TodoHandler) DeleteTodo(c echo.Context) error {
	h.Logger.InfoLog.Println("DeleteTodo called")

	// パスパラメータからidを取得
	id := c.Param("id")

	// Todoユースケースからidを指定してTodoを削除
	err := h.todoUsecase.DeleteTodo(id)
	// エラーがあればエラーレスポンスを返す
	if err != nil {
		switch err.Error() {
		case "id is empty":
			h.Logger.ErrorLog.Printf("Failed to delete todo: %v", err)
			return c.JSON(http.StatusNotFound, map[string]string{
				"message": err.Error(),
			})
		default:
			h.Logger.ErrorLog.Printf("Failed to delete todo: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
			})
		}
	}

	// 削除したTodoをJSON形式で返す
	h.Logger.InfoLog.Println("Deleted todo")
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Todo deleted successfully",
	})
}
