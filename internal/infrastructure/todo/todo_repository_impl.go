package infrastructure_todo

import (
	domain_todo "backend/internal/domain/todo"
	pkg_logger "backend/internal/pkg/logger"
	pkg_supabase "backend/internal/pkg/supabase"
)

// Todoリポジトリ(Impl)
type TodoRepositoryImpl struct {
	Logger         *pkg_logger.AppLogger
	SupabaseClient *pkg_supabase.SupabaseClient
}

// Todoリポジトリのインスタンス化
func NewTodoRepository(l *pkg_logger.AppLogger, sc *pkg_supabase.SupabaseClient) domain_todo.ITodoRepository {
	return &TodoRepositoryImpl{
		Logger:         l,
		SupabaseClient: sc,
	}
}

// 全てのTodoを取得
func (r *TodoRepositoryImpl) GetAllTodos() ([]domain_todo.Todos, error) {
	r.Logger.InfoLog.Println("GetAllTodos called")

	query := `
		SELECT id, description, completed, user_id, created_at, updated_at
		FROM todos
	`

	// Supabaseからクエリを実行し、条件に一致するユーザーを取得
	rows, err := r.SupabaseClient.Pool.Query(r.SupabaseClient.Ctx, query)
	if err != nil {
		r.Logger.ErrorLog.Printf("Failed to fetch todos: %v", err)
		return nil, err
	}

	// Todosのリストを作成
	todos := []domain_todo.Todos{}
	for rows.Next() {
		var todo domain_todo.Todos
		err = rows.Scan(
			&todo.ID,
			&todo.Description,
			&todo.Completed,
			&todo.UserId,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			r.Logger.ErrorLog.Printf("Failed to scan todo: %v", err)
			return nil, err
		}
		todos = append(todos, todo)
	}

	r.Logger.InfoLog.Printf("Fetched %d todos", len(todos))
	return todos, nil
}
