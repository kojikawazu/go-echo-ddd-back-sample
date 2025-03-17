package router

import (
	interfaces_auth "backend/internal/interfaces/auth"
	interfaces_paralell "backend/internal/interfaces/paralell"
	interfaces_sample "backend/internal/interfaces/sample"
	interfaces_todo "backend/internal/interfaces/todo"
	interfaces_user "backend/internal/interfaces/user"

	"github.com/labstack/echo/v4"
)

// ルーティングの設定
func SetUpRouter(
	e *echo.Echo,
	sampleHandler *interfaces_sample.SampleHandler,
	paralellHandler *interfaces_paralell.ParalellHandler,
	userHandler *interfaces_user.UserHandler,
	authHandler *interfaces_auth.AuthHandler,
	todoHandler *interfaces_todo.TodoHandler,
) {
	api := e.Group("/api")
	{
		sample := api.Group("/sample")
		{
			sample.GET("", sampleHandler.ExecSample)
		}

		fetch := api.Group("/fetch")
		{
			fetch.GET("/parallel", paralellHandler.ExecParallel)
			fetch.GET("/series", paralellHandler.ExecSeries)
		}
		user := api.Group("/user")
		{
			user.GET("", authHandler.AuthorizationMiddleware(userHandler.GetAllUsers, "user"))
		}
		todo := api.Group("/todo")
		{
			todo.GET("", authHandler.AuthorizationMiddleware(todoHandler.GetAllTodos, "user"))
			todo.GET("/:id", authHandler.AuthorizationMiddleware(todoHandler.GetTodoById, "user"))
			todo.POST("", authHandler.AuthorizationMiddleware(todoHandler.CreateTodo, "user"))
			todo.PUT("/:id", authHandler.AuthorizationMiddleware(todoHandler.UpdateTodo, "user"))
			todo.DELETE("/:id", authHandler.AuthorizationMiddleware(todoHandler.DeleteTodo, "user"))
		}
		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
		}
	}
}
